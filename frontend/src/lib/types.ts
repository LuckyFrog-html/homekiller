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
    Lesson: Lesson,
    MaxScore: number,
    Description: string,
    Deadline: Date
}

export type Task = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string | null,
    DeletedAt: string | null,
    Description: string,
    LessonID: number,
    Deadline: string,
    MaxScore: number,
    HomeworkFiles: any[],
    HomeworkAnswers: any,
    Lesson: Lesson,
    IsDone: boolean,
    GroupId: number,
    GroupTitle: string,
}

export type Lesson = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string | null,
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

export type Teacher = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string | null,
    Name: string,
    Login: string,
}