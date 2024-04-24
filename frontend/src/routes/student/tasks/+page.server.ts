import { api } from "$lib/api";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { Task } from "$lib/types";

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
