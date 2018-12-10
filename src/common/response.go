package common

type ResponseMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Success  bool              `json:"success"`
	Result   interface{}       `json:"result"`
	Errors   []ResponseMessage `json:"errors"`
	Messages []ResponseMessage `json:"messages"`
}

func NewResponse(result interface{}) Response {
	return Response{
		Success:  true,
		Result:   result,
		Errors:   []ResponseMessage{},
		Messages: []ResponseMessage{},
	}
}
