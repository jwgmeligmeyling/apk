/*
 *  Copyright (c) 2023, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package synchronizer

import (
	"errors"
	"fmt"

	"github.com/wso2/apk/adapter/config"
	"github.com/wso2/apk/adapter/internal/dataholder"
	"github.com/wso2/apk/adapter/internal/discovery/xds"
	"github.com/wso2/apk/adapter/internal/discovery/xds/common"
	"github.com/wso2/apk/adapter/internal/loggers"
	"github.com/wso2/apk/adapter/internal/oasparser/model"
	"github.com/wso2/apk/adapter/internal/operator/constants"
	"github.com/wso2/apk/adapter/pkg/logging"
	"k8s.io/apimachinery/pkg/types"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// deployRestAPIInGateway deploys the related API in CREATE and UPDATE events.
func deployRestAPIInGateway(apiState APIState) error {
	var err error
	if len(apiState.OldOrganizationID) != 0 {
		xds.RemoveAPIFromOrgAPIMap(string((*apiState.APIDefinition).ObjectMeta.UID), apiState.OldOrganizationID)
	}
	if apiState.ProdHTTPRoute == nil {
		var adapterInternalAPI model.AdapterInternalAPI
		adapterInternalAPI.SetInfoAPICR(*apiState.APIDefinition)
		xds.RemoveAPICacheForEnv(adapterInternalAPI, constants.Production)
	}
	if apiState.SandHTTPRoute == nil {
		var adapterInternalAPI model.AdapterInternalAPI
		adapterInternalAPI.SetInfoAPICR(*apiState.APIDefinition)
		xds.RemoveAPICacheForEnv(adapterInternalAPI, constants.Sandbox)
	}
	if apiState.ProdHTTPRoute != nil {
		_, err = GenerateAdapterInternalAPI(apiState, apiState.ProdHTTPRoute, constants.Production)
	}
	if err != nil {
		return err
	}
	if apiState.SandHTTPRoute != nil {
		_, err = GenerateAdapterInternalAPI(apiState, apiState.SandHTTPRoute, constants.Sandbox)
	}
	return err
}

// undeployAPIInGateway undeploys the related API in CREATE and UPDATE events.
func undeployRestAPIInGateway(apiState APIState) error {
	var err error
	if apiState.ProdHTTPRoute != nil {
		err = deleteAPIFromEnv(apiState.ProdHTTPRoute.HTTPRouteCombined, apiState)
	}
	if err != nil {
		loggers.LoggerXds.ErrorC(logging.PrintError(logging.Error2630, logging.MAJOR, "Error undeploying prod httpRoute of API : %v in Organization %v from environments %v."+
			" Hence not checking on deleting the sand httpRoute of the API", string(apiState.APIDefinition.ObjectMeta.UID), apiState.APIDefinition.Spec.Organization,
			getLabelsForAPI(apiState.ProdHTTPRoute.HTTPRouteCombined)))
		return err
	}
	if apiState.SandHTTPRoute != nil {
		err = deleteAPIFromEnv(apiState.SandHTTPRoute.HTTPRouteCombined, apiState)
	}
	return err
}

// GenerateAdapterInternalAPI this will populate a AdapterInternalAPI representation for an HTTPRoute
func GenerateAdapterInternalAPI(apiState APIState, httpRoute *HTTPRouteState, envType string) (*model.AdapterInternalAPI, error) {
	var adapterInternalAPI model.AdapterInternalAPI
	adapterInternalAPI.SetIsDefaultVersion(apiState.APIDefinition.Spec.IsDefaultVersion)
	adapterInternalAPI.SetInfoAPICR(*apiState.APIDefinition)
	adapterInternalAPI.SetAPIDefinitionFile(apiState.APIDefinitionFile)
	adapterInternalAPI.SetAPIDefinitionEndpoint(apiState.APIDefinition.Spec.DefinitionPath)
	adapterInternalAPI.SetSubscriptionValidation(apiState.SubscriptionValidation)
	if apiState.MutualSSL != nil && apiState.MutualSSL.Required != "" && !adapterInternalAPI.GetDisableAuthentications() {
		adapterInternalAPI.SetDisableMtls(apiState.MutualSSL.Disabled)
		adapterInternalAPI.SetMutualSSL(apiState.MutualSSL.Required)
		adapterInternalAPI.SetClientCerts(apiState.APIDefinition.Name, apiState.MutualSSL.ClientCertificates)
	} else {
		adapterInternalAPI.SetDisableMtls(true)
	}
	adapterInternalAPI.EnvType = envType

	environment := apiState.APIDefinition.Spec.Environment
	if environment == "" {
		conf := config.ReadConfigs()
		environment = conf.Adapter.Environment
	}
	adapterInternalAPI.SetEnvironment(environment)

	resourceParams := model.ResourceParams{
		AuthSchemes:               apiState.Authentications,
		ResourceAuthSchemes:       apiState.ResourceAuthentications,
		BackendMapping:            httpRoute.BackendMapping,
		APIPolicies:               apiState.APIPolicies,
		ResourceAPIPolicies:       apiState.ResourceAPIPolicies,
		ResourceScopes:            httpRoute.Scopes,
		InterceptorServiceMapping: apiState.InterceptorServiceMapping,
		BackendJWTMapping:         apiState.BackendJWTMapping,
		RateLimitPolicies:         apiState.RateLimitPolicies,
		ResourceRateLimitPolicies: apiState.ResourceRateLimitPolicies,
	}
	if err := adapterInternalAPI.SetInfoHTTPRouteCR(httpRoute.HTTPRouteCombined, resourceParams); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2631, logging.MAJOR, "Error setting HttpRoute CR info to adapterInternalAPI. %v", err))
		return nil, err
	}
	if err := adapterInternalAPI.Validate(); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2632, logging.MAJOR, "Error validating adapterInternalAPI intermediate representation. %v", err))
		return nil, err
	}
	vHosts := getVhostsForAPI(httpRoute.HTTPRouteCombined)
	labels := getLabelsForAPI(httpRoute.HTTPRouteCombined)
	listeners, relativeSectionNames := getListenersForAPI(httpRoute.HTTPRouteCombined, adapterInternalAPI.UUID)
	// We dont have a use case where a perticular API's two different http routes refer to two different gateway. Hence get the first listener name for the list for processing.
	if len(listeners) == 0 || len(relativeSectionNames) == 0 {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2633, logging.MINOR, "Failed to find a matching listener for http route: %v. ",
			httpRoute.HTTPRouteCombined.Name))
		return nil, errors.New("failed to find matching listener name for the provided http route")
	}
	listenerName := listeners[0]
	sectionName := relativeSectionNames[0]
	if len(listeners) != 0 {
		err := xds.UpdateAPICache(vHosts, labels, listenerName, sectionName, adapterInternalAPI)
		if err != nil {
			loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2633, logging.MAJOR, "Error updating the API : %s:%s in vhosts: %s, API_UUID: %v. %v",
				adapterInternalAPI.GetTitle(), adapterInternalAPI.GetVersion(), vHosts, adapterInternalAPI.UUID, err))
		}
	}

	return &adapterInternalAPI, nil
}

// getVhostForAPI returns the vHosts related to an API.
func getVhostsForAPI(httpRoute *gwapiv1b1.HTTPRoute) []string {
	var vHosts []string
	for _, hostName := range httpRoute.Spec.Hostnames {
		vHosts = append(vHosts, string(hostName))
	}
	fmt.Println("vhosts size: ", len(vHosts))
	return vHosts
}

// getLabelsForAPI returns the labels related to an API.
func getLabelsForAPI(httpRoute *gwapiv1b1.HTTPRoute) []string {
	var labels []string
	var err error
	for _, parentRef := range httpRoute.Spec.ParentRefs {
		err = xds.SanitizeGateway(string(parentRef.Name), false)
		if err != nil {
			loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2653, logging.CRITICAL, "Gateway Label is invalid: %s", string(parentRef.Name)))
		} else {
			labels = append(labels, string(parentRef.Name))
		}
	}
	return labels
}

// getListenersForAPI returns the listeners related to an API.
func getListenersForAPI(httpRoute *gwapiv1b1.HTTPRoute, apiUUID string) ([]string, []string) {
	var listeners []string
	var sectionNames []string
	for _, parentRef := range httpRoute.Spec.ParentRefs {
		namespace := httpRoute.GetNamespace()
		if parentRef.Namespace != nil && *parentRef.Namespace != "" {
			namespace = string(*parentRef.Namespace)
		}
		gateway, found := dataholder.GetGatewayMap()[types.NamespacedName{
			Namespace: namespace,
			Name:      string(parentRef.Name),
		}.String()]
		if found {
			// find the matching listener
			matchedListener, listenerFound := common.FindElement(gateway.Spec.Listeners, func(listener gwapiv1b1.Listener) bool {
				if string(listener.Name) == string(*parentRef.SectionName) {
					return true
				}
				return false
			})
			if listenerFound {
				sectionNames = append(sectionNames, string(matchedListener.Name))
				listeners = append(listeners, common.GetEnvoyListenerName(string(matchedListener.Protocol), uint32(matchedListener.Port)))
				continue
			}
		}
		loggers.LoggerAPKOperator.Errorf("Failed to find matching listeners for the httproute: %+v", httpRoute.Name)
	}
	return listeners, sectionNames
}

func deleteAPIFromEnv(httpRoute *gwapiv1b1.HTTPRoute, apiState APIState) error {
	labels := getLabelsForAPI(httpRoute)
	org := apiState.APIDefinition.Spec.Organization
	uuid := string(apiState.APIDefinition.ObjectMeta.UID)
	return xds.DeleteAPICREvent(labels, uuid, org)
}
