# SG Bus API
## Background
This package aims to simplify and provide a Golang interface for interacting with Singapore's Bus API. More information about the API can be located inside the documentation [here](https://www.mytransport.sg/content/dam/datamall/datasets/LTA_DataMall_API_User_Guide.pdf). 

The API is subjected to changes from time to time (depends on **MyTransport**), so it is not always possible to keep this wrapper up-to-date. That being said, I will still try my best to maintain this repository.

## Prerequisites
Before using this wrapper, the first step will be to request for an API key from **MyTransport**, which can be located [here](https://www.mytransport.sg/content/mytransport/home/dataMall/request-for-api.html). If the provided link is broken due to some unknown reasons, the link to the registration page should still be easy to locate in their [home page](https://www.mytransport.sg/content/mytransport/home/dataMall.html).

## Supported Endpoints
Under **MyTransport** services, there are many endpoints, but not all of them are currently present inside this Golang wrapper. In future, I am intending to support more endpoints and ideally, it should be all of them.

1. **Bus Arrivals**
2. **Bus Services**
3. **Bus Routes**
4. **Bus Stops**

For in-depth details on each endpoint, you should visit their [documentation](https://www.mytransport.sg/content/dam/datamall/datasets/LTA_DataMall_API_User_Guide.pdf) to find out more. Alternatively, you may refer to `model/response.go` for the `struct` of each endpoint's response.

## Usage
1. Install the package
```go
    go get github.com/LeonLyx/sg-bus-api
```

2. Import the library into your codebase
```go
    import (
        "github.com/LeonLyx/sg-bus-api/api"
    )
```

3. Begin to write code
```go
    client := api.NewBusAPIClient("<your-api-key-here>")
    busArrival, err := client.GetBusArrival("83139", "15")
```
For more examples, please refer to `main.go`. If you are unsure about the `struct` implementation for the corresponding endpoint's response, do refer to `model/response.go`.

## Extending this Repository
1. Clone this repository
```
    git clone https://github.com/LeonLyx/sg-bus-api.git
```

2. Navigate to the project directory
```
    cd sg-bus-api/
```

3. Build and run `main.go`
```
    go build
    ./sg-bus-api --account-key <your-api-key-here>
```

## Contributing
As mentioned, it is not quite possible to be updated with the services provided by **MyTransport**. I am hoping to have a collaboration with a group of interested developers. If you would like to extend a helping hand, do hit me up at leonlyxl@gmail.com. Thank you in advance!