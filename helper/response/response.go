package response

type Response map[string]any

func Succes(data, code any, message string) Response {
	msg := "Your request has been successfully processed"
	if message != "" {
		msg = message
	}
	return Response{
		"success": true,
		"code":    code,
		"message": msg,
		"data":    data,
	}
}

// func Error() {

// }

// func ErrorsWithMessage() {

// }
