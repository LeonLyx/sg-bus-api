package main

import (
	"flag"
	"fmt"
)

func main() {
	var accountKey string
	flag.StringVar(&accountKey, "account-key", "your-account-key", "required for endpoint access")
	flag.Parse()

	client := NewBusAPIClient(accountKey)
	busArrival, err := client.GetBusArrival("83139", "15")
	if err != nil {
		fmt.Printf("error getting bus arrival, %s\n", err.Error())
	} else if len(busArrival.Services) > 0 {
		fmt.Printf("%+v\n", busArrival.Services[0])
	}

	busServices, err := client.GetBusServices()
	if err != nil {
		fmt.Printf("error getting bus services, %s\n", err.Error())
	} else if len(busServices.Services) > 0 {
		fmt.Printf("%+v\n", busServices.Services[0])
	}

	busStops, err := client.GetBusStops()
	if err != nil {
		fmt.Printf("error getting bus stops, %s\n", err.Error())
	} else if len(busStops.BusStops) > 0 {
		fmt.Printf("%+v\n", busStops.BusStops[0])
	}

	busRoutes, err := client.GetBusRoutes()
	if err != nil {
		fmt.Printf("error getting bus routes, %s\n", err.Error())
	} else if len(busRoutes.Routes) > 0 {
		fmt.Printf("%+v\n", busRoutes.Routes[0])
	}
}
