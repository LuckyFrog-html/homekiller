export type Student = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string | null,
    DeletedAt: string | null,
    Name: string,
    Stage: number,
    Login: string,
}

export type Homework = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string | null,
    DeletedAt: string | null,
    Deadline: string
    MaxScore: number,
    Description: string,
    Lesson?: Lesson,
}

export type Task = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Description: string,
    LessonID: number,
    Deadline: string,
    MaxScore: number,
    HomeworkFiles: any[] | null,
    HomeworkAnswers: any,
    Lesson: Lesson | null,
    IsDone?: boolean,
    GroupId?: number,
    GroupTitle?: string,
}

export type Lesson = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Date: string,
    GroupID?: number,
    Homeworks?: null,
    Students?: Student[] | null,
    Group?: Group
}

export type Group = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Title: string,
    IsActive: boolean,
    Teacher?: Teacher,
}

export type Solution = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Text: string,
    HomeworkID: number,
    StudentID: number,
    HomeworkAnswerFiles: any[] | null,
    TeacherResumes: null,
    Student: Student | null,
    Homework: Task | null,
    Reviews: Review[] | null,
}

export type Teacher = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Name: string,
    Login: string,
}

export type Review = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Score: number,
    Comment: string,
}
