package json

type AddStudentJson struct {
	Name     string `json:"name"`
	Stage    int64  `json:"stage"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type GetStudentJson struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
