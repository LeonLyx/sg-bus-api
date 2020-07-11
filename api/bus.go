package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/LeonLyx/sg-bus-api/constants"
	"github.com/LeonLyx/sg-bus-api/model"
)

type InvalidAccountKeyError struct{}

type BusAPIGenericError struct {
	StatusCode int
}

type BusAPIClient struct {
	accountKey string
}

func (e *InvalidAccountKeyError) Error() string {
	return fmt.Sprintf("Invalid account key has been provided")
}

func (e *BusAPIGenericError) Error() string {
	return fmt.Sprintf("HTTP response status code : %d", e.StatusCode)
}

func NewBusAPIClient(accountKey string) *BusAPIClient {
	return &BusAPIClient{
		accountKey: accountKey,
	}
}

func (c *BusAPIClient) checkResponseStatusCode(statusCode int) error {
	switch statusCode {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return &InvalidAccountKeyError{}
	default:
		return &BusAPIGenericError{StatusCode: statusCode}
	}
}

func (c *BusAPIClient) request(uri string, queryParameters url.Values, responsePayload interface{}) error {
	requestURI := uri
	if queryParameters != nil {
		requestURI = fmt.Sprintf("%s?%s", requestURI, queryParameters.Encode())
	}

	request, err := http.NewRequest(http.MethodGet, requestURI, bytes.NewBuffer([]byte("")))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("AccountKey", c.accountKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return err
	} else if err = c.checkResponseStatusCode(response.StatusCode); err != nil {
		return err
	}

	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	} else if err = json.Unmarshal(responseBody, responsePayload); err != nil {
		return err
	}
	return nil
}

func (c *BusAPIClient) GetBusArrival(busStopCode string, serviceNo string) (*model.BusArrivalResponse, error) {
	queryParameters := url.Values{}
	queryParameters.Set("BusStopCode", busStopCode)
	queryParameters.Set("ServiceNo", serviceNo)

	responseStruct := &model.BusArrivalResponse{}
	err := c.request(constants.BusArrivalEndpoint, queryParameters, responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct, nil
}

func (c *BusAPIClient) GetBusServices() (*model.BusServiceResponse, error) {
	responseStruct := &model.BusServiceResponse{}
	err := c.request(constants.BusServiceEndpoint, nil, responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct, nil
}

func (c *BusAPIClient) GetBusRoutes() (*model.BusRouteResponse, error) {
	responseStruct := &model.BusRouteResponse{}
	err := c.request(constants.BusRouteEndpoint, nil, responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct, nil
}

func (c *BusAPIClient) GetBusStops() (*model.BusStopResponse, error) {
	responseStruct := &model.BusStopResponse{}
	err := c.request(constants.BusStopEndpoint, nil, responseStruct)
	if err != nil {
		return nil, err
	}
	return responseStruct, nil
}
