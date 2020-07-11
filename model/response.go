package model

type BusArrivalResponse struct {
	Services []*BusArrivalService `json:"Services"`
}

type BusServiceResponse struct {
	Services []*BusService `json:"value"`
}

type BusRouteResponse struct {
	Routes []*BusRoute `json:"value"`
}

type BusStopResponse struct {
	BusStops []*BusStop `json:"value"`
}

type BusArrivalService struct {
	ServiceNo string         `json:"ServiceNo"`
	Operator  string         `json:"Operator"`
	NextBus   BusArrivalTime `json:"NextBus"`
	NextBus2  BusArrivalTime `json:"NextBus2"`
	NextBus3  BusArrivalTime `json:"NextBus3"`
}

type BusArrivalTime struct {
	OriginCode       string `json:"OriginCode"`
	DestinationCode  string `json:"DestinationCode"`
	EstimatedArrival string `json:"EstimatedArrival"`
	Latitude         string `json:"Latitude"`
	Longitude        string `json:"Longitude"`
	VisitNumber      string `json:"VisitNumber"`
	Load             string `json:"Load"`
	Feature          string `json:"Feature"`
	Type             string `json:"Type"`
}

type BusService struct {
	ServiceNo          string `json:"ServiceNo"`
	Operator           string `json:"Operator"`
	Direction          uint   `json:"Direction"`
	Category           string `json:"Category"`
	OriginCode         string `json:"OriginCode"`
	DestinationCode    string `json:"DestinationCode"`
	AMPeakFrequency    string `json:"AM_Peak_Freq"`
	AMOffPeakFrequency string `json:"AM_Offpeak_Freq"`
	PMPeakFrequency    string `json:"PM_Peak_Freq"`
	PMOffPeakFrequency string `json:"PM_Offpeak_Freq"`
	LoopDescription    string `json:"LoopDesc"`
}

type BusRoute struct {
	ServiceNo        string  `json:"ServiceNo"`
	Operator         string  `json:"Operator"`
	Direction        uint    `json:"Direction"`
	StopSequence     int     `json:"StopSequence"`
	BusStopCode      string  `json:"BusStopCode"`
	Distance         float32 `json:"Distance"`
	WeekdayFirstBus  string  `json:"WD_FirstBus"`
	WeekdayLastBus   string  `json:"WD_LastBus"`
	SaturdayFirstBus string  `json:"SAT_FirstBus"`
	SaturdayLastBus  string  `json:"SAT_LastBus"`
	SundayFirstBus   string  `json:"SUN_FirstBus"`
	SundayLastBus    string  `json:"SUN_LastBus"`
}

type BusStop struct {
	BusStopCode string  `json:"BusStopCode"`
	RoadName    string  `json:"RoadName"`
	Description string  `json:"Description"`
	Latitude    float64 `json:"Latitude"`
	Longitude   float64 `json:"Longitude"`
}
