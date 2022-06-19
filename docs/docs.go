// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "eibrahimarisoy@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cars/": {
            "get": {
                "description": "List all cars with pagination and search",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "List all cars",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Pickup date",
                        "name": "pickup_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Pickup time",
                        "name": "pickup_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Dropoff date",
                        "name": "dropoff_date",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Dropoff time",
                        "name": "dropoff_time",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "UUID formatted ID",
                        "name": "location_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/car.CarListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a car with payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "car"
                ],
                "summary": "Create a car",
                "parameters": [
                    {
                        "description": "Car payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/car.CarRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/car.CarResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/locations/": {
            "get": {
                "description": "List all locations with pagination and search",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "location"
                ],
                "summary": "List all locations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/location.LocationListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a location with payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "location"
                ],
                "summary": "Create a location",
                "parameters": [
                    {
                        "description": "Location payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/location.LocationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/location.LocationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/offices/": {
            "get": {
                "description": "List all offices with pagination and search",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offices"
                ],
                "summary": "List all offices",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/office.OfficeListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a office with payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "offices"
                ],
                "summary": "Create a office",
                "parameters": [
                    {
                        "description": "Office payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/office.OfficeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/office.OfficeResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/reservations/": {
            "get": {
                "description": "List all reservations with pagination and search",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "List all reservations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/reservation.ReservationListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a reservation with payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "reservation"
                ],
                "summary": "Create a reservation",
                "parameters": [
                    {
                        "description": "Reservation payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/reservation.ReservationRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/reservation.ReservationResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            }
        },
        "/vendors/": {
            "get": {
                "description": "List all vendors with pagination and search",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "List all vendors",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search query",
                        "name": "q",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Page limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/vendors.VendorListResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a vendor with payload",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "vendors"
                ],
                "summary": "Create a vendor",
                "parameters": [
                    {
                        "description": "Vendor payload",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/vendors.VendorRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/vendors.VendorResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/_type.APIErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "_type.APIErrorResponse": {
            "type": "object",
            "properties": {
                "errCode": {
                    "type": "integer"
                },
                "errDetails": {},
                "errMessage": {
                    "type": "string"
                }
            }
        },
        "car.CarListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/car.CarResponse"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "q": {
                    "type": "string"
                },
                "total_pages": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "car.CarRequest": {
            "type": "object",
            "properties": {
                "fuel": {
                    "type": "string",
                    "enum": [
                        "gas",
                        "diesel",
                        "electric"
                    ]
                },
                "name": {
                    "type": "string"
                },
                "office_id": {
                    "type": "string"
                },
                "status": {
                    "type": "string",
                    "enum": [
                        "available",
                        "rented"
                    ]
                },
                "transmission": {
                    "type": "string",
                    "enum": [
                        "automatic",
                        "manual"
                    ]
                },
                "vendor_id": {
                    "type": "string"
                }
            }
        },
        "car.CarResponse": {
            "type": "object",
            "properties": {
                "fuel": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "office": {
                    "$ref": "#/definitions/office.OfficeSimpleResponse"
                },
                "status": {
                    "type": "string"
                },
                "transmission": {
                    "type": "string"
                },
                "vendor": {
                    "$ref": "#/definitions/vendors.VendorResponse"
                }
            }
        },
        "car.CarSimpleResponse": {
            "type": "object",
            "properties": {
                "fuel": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "office": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "transmission": {
                    "type": "string"
                },
                "vendor": {
                    "type": "string"
                }
            }
        },
        "driver.DriverRequest": {
            "type": "object",
            "required": [
                "birthday",
                "email",
                "first_name",
                "identification_number",
                "last_name",
                "phone"
            ],
            "properties": {
                "birthday": {
                    "type": "string",
                    "format": "02-01-2006"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "identification_number": {
                    "type": "string",
                    "format": "12345678901"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string",
                    "format": "+905012345678"
                }
            }
        },
        "driver.DriverResponse": {
            "type": "object",
            "properties": {
                "birthday": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "identification_number": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "location.LocationListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/location.LocationResponse"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "q": {
                    "type": "string"
                },
                "total_pages": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "location.LocationRequest": {
            "type": "object",
            "properties": {
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "location.LocationResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "is_active": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "office.OfficeListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/office.OfficeResponse"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "q": {
                    "type": "string"
                },
                "total_pages": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "office.OfficeRequest": {
            "type": "object",
            "properties": {
                "closing_hours": {
                    "type": "string"
                },
                "location_id": {
                    "type": "string"
                },
                "opening_hours": {
                    "type": "string"
                },
                "vendor_id": {
                    "type": "string"
                },
                "working_days": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        },
        "office.OfficeResponse": {
            "type": "object",
            "properties": {
                "closing_hours": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "$ref": "#/definitions/location.LocationResponse"
                },
                "opening_hours": {
                    "type": "string"
                },
                "vendor": {
                    "$ref": "#/definitions/vendors.VendorResponse"
                },
                "working_days": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "office.OfficeSimpleResponse": {
            "type": "object",
            "properties": {
                "closing_hours": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "location": {
                    "type": "string"
                },
                "opening_hours": {
                    "type": "string"
                }
            }
        },
        "reservation.ReservationListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/reservation.ReservationResponse"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "q": {
                    "type": "string"
                },
                "total_pages": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "reservation.ReservationRequest": {
            "type": "object",
            "required": [
                "car_id",
                "driver",
                "drop_off_date",
                "drop_off_time",
                "dropoff_location_id",
                "pick_up_date",
                "pick_up_time",
                "pickup_location_id"
            ],
            "properties": {
                "car_id": {
                    "type": "string",
                    "format": "UUID"
                },
                "driver": {
                    "$ref": "#/definitions/driver.DriverRequest"
                },
                "drop_off_date": {
                    "type": "string",
                    "format": "02-01-2006"
                },
                "drop_off_time": {
                    "type": "string",
                    "format": "15:04"
                },
                "dropoff_location_id": {
                    "type": "string",
                    "format": "UUID"
                },
                "pick_up_date": {
                    "type": "string",
                    "format": "02-01-2006"
                },
                "pick_up_time": {
                    "type": "string",
                    "format": "15:04"
                },
                "pickup_location_id": {
                    "type": "string",
                    "format": "UUID"
                }
            }
        },
        "reservation.ReservationResponse": {
            "type": "object",
            "properties": {
                "car": {
                    "$ref": "#/definitions/car.CarSimpleResponse"
                },
                "driver_response": {
                    "$ref": "#/definitions/driver.DriverResponse"
                },
                "drop_off_date": {
                    "type": "string"
                },
                "drop_off_time": {
                    "type": "string"
                },
                "dropoff_location": {
                    "$ref": "#/definitions/location.LocationResponse"
                },
                "id": {
                    "type": "string"
                },
                "pick_up_date": {
                    "type": "string"
                },
                "pick_up_time": {
                    "type": "string"
                },
                "pickup_location": {
                    "$ref": "#/definitions/location.LocationResponse"
                }
            }
        },
        "vendors.VendorListResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/vendors.VendorResponse"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "page": {
                    "type": "integer"
                },
                "q": {
                    "type": "string"
                },
                "total_pages": {
                    "type": "integer"
                },
                "total_rows": {
                    "type": "integer"
                }
            }
        },
        "vendors.VendorRequest": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        },
        "vendors.VendorResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Car Rental API",
	Description:      "This is a Car Rental API implementation.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
