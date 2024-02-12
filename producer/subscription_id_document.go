// Copyright 2019 free5GC.org
//
// SPDX-License-Identifier: Apache-2.0
//

/*
 * NSSF NSSAI Availability
 *
 * NSSF NSSAI Availability Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package producer

import (
	"net/http"

	"github.com/omec-project/nssf/logger"
	"github.com/omec-project/openapi/models"
	"github.com/omec-project/util/httpwrapper"
	"reflect"
)

// HandleNSSAIAvailabilityUnsubscribe - Deletes an already existing NSSAI availability notification subscription
func HandleNSSAIAvailabilityUnsubscribe(request *httpwrapper.Request) *httpwrapper.Response {
	logger.Nssaiavailability.Infof("Handle NSSAIAvailabilityUnsubscribe")

	subscriptionID := request.Params["subscriptionId"]

	problemDetails := NSSAIAvailabilityUnsubscribeProcedure(subscriptionID)

	if problemDetails == nil {
		return httpwrapper.NewResponse(http.StatusNoContent, nil, nil)
	} else if reflect.DeepEqual(*problemDetails, models.ProblemDetails{}) {
		problemDetails = &models.ProblemDetails{
			Status: http.StatusForbidden,
			Cause:  "UNSPECIFIED",
		}
		return httpwrapper.NewResponse(http.StatusForbidden, nil, problemDetails)
	}
	return httpwrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
}
