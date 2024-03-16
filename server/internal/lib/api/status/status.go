package status

type Response struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error,omitempty"`
}

func HttpError(status int, msg string) Response {
	return Response{StatusCode: status, Error: msg}
}

type StudentResponse struct {
	Response
	StudentId int `json:"student_id"`
}
