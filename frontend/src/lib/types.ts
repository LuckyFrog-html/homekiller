export type Task = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string,
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
    UpdatedAt: string,
    DeletedAt: string,
    Date: string,
    GroupID: number,
    Homeworks: null,
    Students: null,
    Group: null
}
