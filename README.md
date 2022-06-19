# Car Rental API Documentation

This is the documentation for the Car Rental API.
This repository contains Car Rental API code written with golang. It is a RESTful API.

The project has two middleware layers. The first layer is the pagination middleware. It is used to paginate the results.
The second layer is the filter middleware. It is used to filter the car list endpoint results.

Clients can;
 - list active locations
 - list all cars based on query parameters (e.g. pickup location/date/time, dropoff location/date/time, pagination data etc.)
 - create car reservations and list all reservations. Reservations also must consist of a pickup location, dropoff location, start/end date/time and driver information. 
 - More information about the API can be found in the Swagger documentation.

## Using Tools
 - Gin
 - Gorm
 - Postgres
 - net/http
 - viper
 - swagger
 - docker
 - zap (logger)
 - mockgen
 - testify
 - air

## Dependencies used in the project
 - Go 1.18
 - Postgres 14.2
 - Docker

## Clone the repository
```
git clone https://github.com/eibrahimarisoy/car_rental.git
cd car_rental
```

## Configuration
App uses viper to read configuration file.
You can edit file `config.yaml` in the project pkg/config directory.
There is a sample file `config.yaml.example` in the project pkg/config directory.

For Docker you have to set the following environment variables in .env file, details can be found in the .env.example file.

## Run the tests
```
$ go test ./...
```
## Run the project
```
$ docker compose up
```

## Run the project with swagger
Installing swagger please follow the instructions on the link below.

https://goswagger.io/install.html

```
$ swagger serve ./docs/swagger.yaml
```

## Routes
Default **Car Rental API** routes are listed below. 

| METHOD | ROUTE                     | DETAILS                                                 |
|------------------------------------|---------------------------------------------------------|
| GET    | /api/v1/locations/        | list active locations (with Paging and query parameter) |
| POST   | /api/v1/locations/        | create location                                         | 
| GET    | /api/v1/vendors/          | list vendors (with Paging)                              |
| POST   | /api/v1/vendors/          | create vendor                                           |
| GET    | /api/v1/offices/          | list offices (with Paging)                              |
| POST   | /api/v1/offices/          | create office                                           |
| GET    | /api/v1/cars/             | list cars (with Paging and filter parameters)           |
| POST   | /api/v1/cars/             | create car                                              |                   
| GET    | /api/v1/reservations/     | list reservations (with Paging)                         | 
| POST   | /api/v1/reservations/     | create reservation                                      |
| GET    | /api/v1/healthz           | health check                                            | 
| GET    | /api/v1/readyz            | readiness check                                         |

## Nice to have
 - [ ] add more test
 - [ ] jwt authentication
 - [ ] add more endpoints

## Contact
If you want to contact me you can reach me at <eibrahimarisoy@gmail.com>.
