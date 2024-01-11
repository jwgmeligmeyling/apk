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
	"github.com/wso2/apk/common-go-libs/apis/dp/v1alpha2"
	"k8s.io/apimachinery/pkg/types"
	gwapiv1b1 "sigs.k8s.io/gateway-api/apis/v1beta1"
)

// deployGQLAPIInGateway deploys the related API in CREATE and UPDATE events.
func deployGQLAPIInGateway(apiState APIState) error {
	var err error
	if len(apiState.OldOrganizationID) != 0 {
		xds.RemoveAPIFromOrgAPIMap(string((*apiState.APIDefinition).ObjectMeta.UID), apiState.OldOrganizationID)
	}
	if apiState.ProdGQLRoute == nil {
		var adapterInternalAPI model.AdapterInternalAPI
		adapterInternalAPI.SetInfoAPICR(*apiState.APIDefinition)
		xds.RemoveAPICacheForEnv(adapterInternalAPI, constants.Production)
	}
	if apiState.SandGQLRoute == nil {
		var adapterInternalAPI model.AdapterInternalAPI
		adapterInternalAPI.SetInfoAPICR(*apiState.APIDefinition)
		xds.RemoveAPICacheForEnv(adapterInternalAPI, constants.Sandbox)
	}
	if apiState.ProdGQLRoute != nil {
		_, err = generateGQLAdapterInternalAPI(apiState, apiState.ProdGQLRoute, constants.Production)
	}
	if err != nil {
		return err
	}
	if apiState.SandGQLRoute != nil {
		_, err = generateGQLAdapterInternalAPI(apiState, apiState.SandGQLRoute, constants.Sandbox)
	}
	return err
}

// generateGQLAdapterInternalAPI this will populate a AdapterInternalAPI representation for an GQLRoute
func generateGQLAdapterInternalAPI(apiState APIState, gqlRoute *GQLRouteState, envType string) (*model.AdapterInternalAPI, error) {
	var adapterInternalAPI model.AdapterInternalAPI
	adapterInternalAPI.SetIsDefaultVersion(apiState.APIDefinition.Spec.IsDefaultVersion)
	adapterInternalAPI.SetInfoAPICR(*apiState.APIDefinition)
	adapterInternalAPI.SetAPIDefinitionFile(apiState.APIDefinitionFile)
	adapterInternalAPI.SetAPIDefinitionEndpoint(apiState.APIDefinition.Spec.DefinitionPath)
	adapterInternalAPI.SetSubscriptionValidation(apiState.SubscriptionValidation)
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
		BackendMapping:            gqlRoute.BackendMapping,
		APIPolicies:               apiState.APIPolicies,
		ResourceAPIPolicies:       apiState.ResourceAPIPolicies,
		ResourceScopes:            gqlRoute.Scopes,
		InterceptorServiceMapping: apiState.InterceptorServiceMapping,
		BackendJWTMapping:         apiState.BackendJWTMapping,
		RateLimitPolicies:         apiState.RateLimitPolicies,
		ResourceRateLimitPolicies: apiState.ResourceRateLimitPolicies,
	}
	if err := adapterInternalAPI.SetInfoGQLRouteCR(gqlRoute.GQLRouteCombined, resourceParams); err != nil {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2631, logging.MAJOR,
			"Error setting GQLRoute CR info to adapterInternalAPI. %v", err))
		return nil, err
	}
	vHosts := getVhostsForGQLAPI(gqlRoute.GQLRouteCombined)
	labels := getLabelsForGQLAPI(gqlRoute.GQLRouteCombined)
	listeners, relativeSectionNames := getListenersForGQLAPI(gqlRoute.GQLRouteCombined, adapterInternalAPI.UUID)
	// We dont have a use case where a perticular API's two different gql routes refer to two different gateway. Hence get the first listener name for the list for processing.
	if len(listeners) == 0 || len(relativeSectionNames) == 0 {
		loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2633, logging.MINOR, "Failed to find a matching listener for gql route: %v. ",
			gqlRoute.GQLRouteCombined.Name))
		return nil, errors.New("failed to find matching listener name for the provided gql route")
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
func getVhostsForGQLAPI(gqlRoute *v1alpha2.GQLRoute) []string {
	var vHosts []string
	for _, hostName := range gqlRoute.Spec.Hostnames {
		vHosts = append(vHosts, string(hostName))
	}
	fmt.Println("vhosts size: ", len(vHosts))
	return vHosts
}

// getLabelsForAPI returns the labels related to an API.
func getLabelsForGQLAPI(gqlRoute *v1alpha2.GQLRoute) []string {
	var labels []string
	var err error
	for _, parentRef := range gqlRoute.Spec.ParentRefs {
		err = xds.SanitizeGateway(string(parentRef.Name), false)
		if err != nil {
			loggers.LoggerAPKOperator.ErrorC(logging.PrintError(logging.Error2653, logging.CRITICAL, "Gateway Label is invalid: %s", string(parentRef.Name)))
		} else {
			labels = append(labels, string(parentRef.Name))
		}
	}
	return labels
}

// getListenersForGQLAPI returns the listeners related to an API.
func getListenersForGQLAPI(gqlRoute *v1alpha2.GQLRoute, apiUUID string) ([]string, []string) {
	var listeners []string
	var sectionNames []string
	for _, parentRef := range gqlRoute.Spec.ParentRefs {
		namespace := gqlRoute.GetNamespace()
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
		loggers.LoggerAPKOperator.Errorf("Failed to find matching listeners for the gqlroute: %+v", gqlRoute.Name)
	}
	return listeners, sectionNames
}

func deleteGQLAPIFromEnv(gqlRoute *v1alpha2.GQLRoute, apiState APIState) error {
	labels := getLabelsForGQLAPI(gqlRoute)
	org := apiState.APIDefinition.Spec.Organization
	uuid := string(apiState.APIDefinition.ObjectMeta.UID)
	return xds.DeleteAPICREvent(labels, uuid, org)
}

// undeployGQLAPIInGateway undeploys the related API in CREATE and UPDATE events.
func undeployGQLAPIInGateway(apiState APIState) error {
	var err error
	if apiState.ProdGQLRoute != nil {
		err = deleteGQLAPIFromEnv(apiState.ProdGQLRoute.GQLRouteCombined, apiState)
	}
	if err != nil {
		loggers.LoggerXds.ErrorC(logging.PrintError(logging.Error2630, logging.MAJOR, "Error undeploying prod gqlRoute of API : %v in Organization %v from environments."+
			" Hence not checking on deleting the sand gqlRoute of the API", string(apiState.APIDefinition.ObjectMeta.UID), apiState.APIDefinition.Spec.Organization))
		return err
	}
	if apiState.SandGQLRoute != nil {
		err = deleteGQLAPIFromEnv(apiState.SandGQLRoute.GQLRouteCombined, apiState)
	}
	return err
}
