package helpers

type ResponseError struct {
	Error   string `json:"status"`
	Message string `json:"message"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
