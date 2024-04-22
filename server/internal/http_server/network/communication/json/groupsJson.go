package json

type AddGroupJson struct {
	Title string `json:"title"`
}

type AddStudentToGroupJson struct {
	StudentsIds []uint `json:"students_ids"`
	GroupId     uint   `json:"group_id"`
}
