import { api } from "$lib/api";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

type Task = {
    ID: number,
    CreatedAt: string,
    UpdatedAt: string,
    DeletedAt: string,
    Description: string,
    LessonID: number,
    Deadline: string,
    MaxScore: number,
    HomeworkFiles: any,
    HomeworkAnswers: any,
    Lesson: any,
    IsDone: boolean,
    GroupId: number,
    GroupTitle: string,
}

/** @type {PageServerLoad} */
export async function load({ cookies }: Parameters<PageServerLoad>[0]): Promise<{ tasks: Task[] }> {
    const token = cookies.get('token');
    const req = await api.get<any>('/homeworks', { token });

    if (req.type === "success") {
        const tasks = req.data.homeworks as Task[] || [];
        return { tasks };
    }


    return redirect(302, '/login');
}
