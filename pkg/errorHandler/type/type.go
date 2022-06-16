package _type

type APIErrorResponse struct {
	Code    int64       `json:"errCode"`
	Message string      `json:"errMessage"`
	Details interface{} `json:"errDetails"`
}
