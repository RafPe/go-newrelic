package newrelic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"reflect"

	"github.com/google/go-querystring/query"
	log "github.com/sirupsen/logrus"
)

// Client represents API client for communicating with NewRelic REST services.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client

	// API key used for calls
	apiKey string

	// Manage many accounts with one API client
	// https://learn.akamai.com/en-us/learn_akamai/getting_started_with_akamai_developers/developer_tools/accountSwitch.html
	accountSwitchKey     string
	accountSwitchEnabled bool

	// Services offered by NewRelic
	Syntethics *SyntethicService
}

// prepareQueryParameters Allows for easy preparation of query string
func (cl *Client) prepareQueryParameters(params interface{}) (queryString string, err error) {
	v, err := query.Values(params)

	// If we do have account switch key - we will add it and toggle ASK back to disabled
	if cl.accountSwitchEnabled == true {
		v.Add("accountSwitchKey", cl.accountSwitchKey)
	}

	if err != nil {
		return "", err
	}

	return v.Encode(), nil
}

// makeAPIRequest creates an HTTP request that can be sent to Akamai APIs. It will handle security headers and signinig of the request.
//
func (cl *Client) makeAPIRequest(method, path string, queryParams, structResponse, structRequest interface{}, headers map[string]string) (*ClientResponse, error) {

	log.Debug("[NewRequest]::Prepare URL for http request")
	targetURL, err := prepareURL(cl.baseURL, path)
	if err != nil {
		return nil, err
	}

	log.Debug("[NewRequest]::Create http request")
	req, err := http.NewRequest(method, targetURL.String(), nil)
	if err != nil {
		return nil, err
	}

	/*
		Modify request for POST/PUT
	*/
	if method == http.MethodPost || method == http.MethodPut {
		log.Info("[NewRequest]::Prepare request body object")
		log.Debug(fmt.Sprintf("[NewRequest]::Method is %s", method))
		log.Debug("[NewRequest]::Marshal request object")

		reqType := reflect.TypeOf(structRequest)
		log.Debug(fmt.Sprintf("[NewRequest]::Object request provided type ( %s ) ", reqType))

		bodyBytes, err := json.Marshal(structRequest)
		if err != nil {
			return nil, err
		}
		bodyReader := bytes.NewReader(bodyBytes)

		log.Debug("[NewRequest]::Body object added to request")
		req.Body = ioutil.NopCloser(bodyReader)
		req.ContentLength = int64(bodyReader.Len())

		log.Debug("[NewRequest]::Body object is:" + string(bodyBytes))
		log.Debug("[NewRequest]::Set header Content-Type to 'application/json' ")
		req.Header.Set("Content-Type", "application/json")

	}

	/*
		Add query params
	*/
	if queryParams != nil {
		log.Debug("[NewRequest]::Add query string parameters")
		rawQueryString, queryStringError := cl.prepareQueryParameters(queryParams)

		req.URL.RawQuery = rawQueryString

		if queryStringError != nil {
			log.Debug("[NewRequest]::Error adding query string parameters")
			log.Debug(fmt.Sprintf("[NewRequest]:: %s", queryStringError.Error()))
			return nil, queryStringError
		}
	}

	/*
		Add headers
	*/
	if headers != nil {
		log.Debug("[NewRequest]::Add extra headers")
		for k, v := range headers {
			log.Debug(fmt.Sprintf("[NewRequest]::Adding %s:%s", k, v))
			req.Header.Add(k, v)
		}
	}

	/*
		Add signature header
	*/
	authorizationHeader := AuthString(cl.credentials, req, []string{})
	log.Debug("[NewRequest]::Set header Authorization")
	req.Header.Add("Authorization", authorizationHeader)

	/*
		Execute request
	*/
	log.Info("[NewRequest]::Execute http request")
	log.Debug(fmt.Sprintf("[NewRequest]::Calling %s", req.URL.RequestURI()))
	resp, err := cl.client.Do(req)
	if err != nil {
		log.Debug("[NewRequest]::Error making request")
		log.Debug(fmt.Sprintf("[NewRequest]:: %s", err.Error()))
		return nil, err
	}
	defer resp.Body.Close()

	/*
		Process response
	*/

	// Save a copy of this request for debugging.
	respDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return nil, err
	}

	log.Debug(string(respDump))

	log.Debug("[NewRequest]::Processing response")
	clientResp := &ClientResponse{}

	log.Debug("[NewRequest]::Read response body")
	byt, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Debug("[NewRequest]::Error reading response body")
		log.Debug(fmt.Sprintf("[NewRequest]:: %s", err.Error()))
		return nil, err
	}

	log.Debug("[NewRequest]::Set client object response and body")
	clientResp.Response = resp
	clientResp.Body = string(byt)

	if structResponse != nil {
		respType := reflect.TypeOf(structResponse)
		log.Debug(fmt.Sprintf("[NewRequest]::Map response to provided type ( %s ) ", respType))

		if err = json.Unmarshal([]byte(byt), &structResponse); err != nil {
			log.Debug("[NewRequest]::Error while unmarshaling response body")
			log.Debug(fmt.Sprintf("[NewRequest]:: %s", err.Error()))
			return nil, err
		}
	}

	log.Debug("[NewRequest]::Return response")

	return clientResp, nil
}
