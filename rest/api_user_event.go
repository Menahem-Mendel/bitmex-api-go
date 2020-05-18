package rest

import (
	"context"

	"github.com/Menahem-Mendel/bitmex-api-go/models"
)

// UserEventService
type UserEventService []models.UserEvent

// UserEventConf
type UserEventConf struct {
	Count   float64 `url:"count,omitempty"`
	StartID float64 `url:"startId,omitempty"`
}

// GetUserEvent
func (c Client) GetUserEvent(ctx context.Context, f UserEventConf) (UserEventService, error) {
	var out UserEventService

	return out, nil
}

// func (a *UserEventApiService) UserEventGet(ctx context.Context, localVarOptionals *UserEventGetOpts) ([]UserEvent, *http.Response, error) {
// 	var (
// 		localVarHttpMethod = strings.ToUpper("Get")
// 		localVarPostBody   interface{}
// 		localVarFileName   string
// 		localVarFileBytes  []byte
// 		localVarReturnValue []UserEvent
// 	)

// 	// create path and map variables
// 	localVarPath := a.client.cfg.BasePath + "/userEvent"

// 	localVarHeaderParams := make(map[string]string)
// 	localVarQueryParams := url.Values{}
// 	localVarFormParams := url.Values{}

// 	if localVarOptionals != nil && localVarOptionals.Count.IsSet() {
// 		localVarQueryParams.Add("count", parameterToString(localVarOptionals.Count.Value(), ""))
// 	}
// 	if localVarOptionals != nil && localVarOptionals.StartId.IsSet() {
// 		localVarQueryParams.Add("startId", parameterToString(localVarOptionals.StartId.Value(), ""))
// 	}
// 	// to determine the Content-Type header
// 	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded"}

// 	// set Content-Type header
// 	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
// 	if localVarHttpContentType != "" {
// 		localVarHeaderParams["Content-Type"] = localVarHttpContentType
// 	}

// 	// to determine the Accept header
// 	localVarHttpHeaderAccepts := []string{"application/json", "application/xml", "text/xml", "application/javascript", "text/javascript"}

// 	// set Accept header
// 	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
// 	if localVarHttpHeaderAccept != "" {
// 		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
// 	}
// 	if ctx != nil {
// 		// API Key Authentication
// 		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
// 			var key string
// 			if auth.Prefix != "" {
// 				key = auth.Prefix + " " + auth.Key
// 			} else {
// 				key = auth.Key
// 			}
// 			localVarHeaderParams["api-expires"] = key

// 		}
// 	}
// 	if ctx != nil {
// 		// API Key Authentication
// 		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
// 			var key string
// 			if auth.Prefix != "" {
// 				key = auth.Prefix + " " + auth.Key
// 			} else {
// 				key = auth.Key
// 			}
// 			localVarHeaderParams["api-key"] = key

// 		}
// 	}
// 	if ctx != nil {
// 		// API Key Authentication
// 		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
// 			var key string
// 			if auth.Prefix != "" {
// 				key = auth.Prefix + " " + auth.Key
// 			} else {
// 				key = auth.Key
// 			}
// 			localVarHeaderParams["api-signature"] = key

// 		}
// 	}
// 	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
// 	if err != nil {
// 		return localVarReturnValue, nil, err
// 	}

// 	localVarHttpResponse, err := a.client.callAPI(r)
// 	if err != nil || localVarHttpResponse == nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
// 	localVarHttpResponse.Body.Close()
// 	if err != nil {
// 		return localVarReturnValue, localVarHttpResponse, err
// 	}

// 	if localVarHttpResponse.StatusCode < 300 {
// 		// If we succeed, return the data, otherwise pass on to decode error.
// 		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
// 		if err == nil {
// 			return localVarReturnValue, localVarHttpResponse, err
// 		}
// 	}

// 	if localVarHttpResponse.StatusCode >= 300 {
// 		newErr := GenericSwaggerError{
// 			body: localVarBody,
// 			error: localVarHttpResponse.Status,
// 		}

// 		if localVarHttpResponse.StatusCode == 200 {
// 			var v []UserEvent
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
// 				if err != nil {
// 					newErr.error = err.Error()
// 					return localVarReturnValue, localVarHttpResponse, newErr
// 				}
// 				newErr.model = v
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 		}

// 		if localVarHttpResponse.StatusCode == 400 {
// 			var v ModelError
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
// 				if err != nil {
// 					newErr.error = err.Error()
// 					return localVarReturnValue, localVarHttpResponse, newErr
// 				}
// 				newErr.model = v
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 		}

// 		if localVarHttpResponse.StatusCode == 401 {
// 			var v ModelError
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
// 				if err != nil {
// 					newErr.error = err.Error()
// 					return localVarReturnValue, localVarHttpResponse, newErr
// 				}
// 				newErr.model = v
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 		}

// 		if localVarHttpResponse.StatusCode == 403 {
// 			var v ModelError
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
// 				if err != nil {
// 					newErr.error = err.Error()
// 					return localVarReturnValue, localVarHttpResponse, newErr
// 				}
// 				newErr.model = v
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 		}

// 		if localVarHttpResponse.StatusCode == 404 {
// 			var v ModelError
// 			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
// 				if err != nil {
// 					newErr.error = err.Error()
// 					return localVarReturnValue, localVarHttpResponse, newErr
// 				}
// 				newErr.model = v
// 				return localVarReturnValue, localVarHttpResponse, newErr
// 		}

// 		return localVarReturnValue, localVarHttpResponse, newErr
// 	}

// 	return localVarReturnValue, localVarHttpResponse, nil
// }
