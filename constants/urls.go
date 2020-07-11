package constants

import "fmt"

const (
	BaseApiUrl = "http://datamall2.mytransport.sg/ltaodataservice"
)

var (
	BusArrivalEndpoint = fmt.Sprintf("%s/BusArrivalv2", BaseApiUrl)
	BusServiceEndpoint = fmt.Sprintf("%s/BusServices", BaseApiUrl)
	BusRouteEndpoint = fmt.Sprintf("%s/BusRoutes", BaseApiUrl)
	BusStopEndpoint = fmt.Sprintf("%s/BusStops", BaseApiUrl)
)
