package response

type SuccessResponse struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message interface{} `json:"message"`
}

type ErrorResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func Success(data interface{}) SuccessResponse {
	return SuccessResponse{
		Code:    200,
		Data:    data,
		Message: "",
	}
}

func Error(code int, err error) ErrorResponse {
	return ErrorResponse{
		Code:    code,
		Message: err.Error(),
	}
}
