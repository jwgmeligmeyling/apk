/*
 *  Copyright (c) 2024, WSO2 LLC. (http://www.wso2.org) All Rights Reserved.
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
 */

package semanticversion

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	logger "github.com/wso2/apk/adapter/pkg/loggers"
)

// SemVersion is the struct to store the version components of an API
type SemVersion struct {
	Version string
	Major   int
	Minor   int
	Patch   *int
}

// ValidateAndGetVersionComponents validates version string and extracts version components
func ValidateAndGetVersionComponents(version string, apiName string) (*SemVersion, error) {
	versionComponents := strings.Split(version, ".")

	// If the versionComponents length is less than 2, return error
	if len(versionComponents) < 2 || !strings.HasPrefix(versionComponents[0], "v") {
		logger.LoggerSemanticVersion.Errorf("API version validation failed for API: %v. API Version: %v", apiName, version)
		errMessage := "Invalid version: " + version + " for API: " + apiName +
			". API version should be in the format x.y.z, x.y, vx.y.z or vx.y where x,y,z are non-negative integers" +
			" and v is version prefix"
		return nil, errors.New(errMessage)
	}

	majorVersionStr := strings.TrimPrefix(versionComponents[0], "v")

	majorVersion, majorVersionConvErr := strconv.Atoi(majorVersionStr)
	minorVersion, minorVersionConvErr := strconv.Atoi(versionComponents[1])
	if majorVersionConvErr != nil || majorVersion < 0 {
		logger.LoggerSemanticVersion.Errorf(fmt.Sprintf("API major version should be a non-negative integer in API: %v. API Version: %v", apiName, version), majorVersionConvErr)
		return nil, errors.New("invalid version format")
	}

	if minorVersionConvErr != nil || minorVersion < 0 {
		logger.LoggerSemanticVersion.Errorf(fmt.Sprintf("API minor version should be a non-negative integer in API: %v. API Version: %v", apiName, version), minorVersionConvErr)
		return nil, errors.New("invalid version format")
	}

	if len(versionComponents) == 2 {
		return &SemVersion{
			Version: version,
			Major:   majorVersion,
			Minor:   minorVersion,
			Patch:   nil,
		}, nil
	}

	patchVersion, patchVersionConvErr := strconv.Atoi(versionComponents[2])
	if patchVersionConvErr != nil {
		logger.LoggerSemanticVersion.Errorf(fmt.Sprintf("API patch version should be an integer in API: %v. API Version: %v", apiName, version), patchVersionConvErr)
		return nil, errors.New("invalid version format")
	}
	return &SemVersion{
		Version: version,
		Major:   majorVersion,
		Minor:   minorVersion,
		Patch:   &patchVersion,
	}, nil
}

// Compare - compares two semantic versions and returns true
// if `version` is greater or equal than `baseVersion`
func (baseVersion SemVersion) Compare(version SemVersion) bool {
	// Compare major version
	if baseVersion.Major != version.Major {
		return baseVersion.Major < version.Major
	}

	// Compare minor version
	if baseVersion.Minor != version.Minor {
		return baseVersion.Minor < version.Minor
	}

	// Compare patch version
	if baseVersion.Patch != nil && version.Patch != nil {
		return *baseVersion.Patch < *version.Patch
	} else if baseVersion.Patch != nil {
		return false
	} else if version.Patch != nil {
		return true
	}

	// Versions are equal
	return true
}
