package json

type AddGroupJson struct {
	Title string `json:"title"`
}

type AddStudentToGroupJson struct {
	StudentsIds []uint `json:"students_ids"`
}

type DeleteGroupJson struct {
	GroupId uint `json:"group_id"`
}

type DeleteStudentFromGroupJson struct {
	GroupId   uint `json:"group_id"`
	StudentId uint `json:"student_id"`
}
