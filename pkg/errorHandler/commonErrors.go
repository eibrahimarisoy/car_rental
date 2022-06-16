package errorHandler

import (
	"net/http"

	_type "github.com/eibrahimarisoy/car_rental/pkg/errorHandler/type"
)

var (
	NotFoundError = _type.ErrorType{
		Code:    http.StatusNotFound,
		Message: "Record not found.",
	}
	BindError = _type.ErrorType{
		Code:    http.StatusBadRequest,
		Message: "Unable to bind the request body.",
	}
	UnmarshalError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Unmarshal Error: Unable to decode into struct",
	}
	MarshalError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Marshal Error",
	}
	GormOpenError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Failed to open db session matching the entered values",
	}
	SqlDBError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "SQL DB Error",
	}
	SqlDBPingError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "SQL DB Ping Error",
	}
	ConfigNotFoundError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Config File Not Found Error",
	}
	DBMigrateError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBMigrate error.",
	}
	DBCreateError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBCreate error.",
	}
	DBUpdateError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBUpdate error.",
	}
	DBDeleteError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "DBDelete error.",
	}
	InternalServerError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "Internal server error.",
	}
	ClientError = _type.ErrorType{
		Code:    http.StatusInternalServerError,
		Message: "HttpClient error.",
	}
)
