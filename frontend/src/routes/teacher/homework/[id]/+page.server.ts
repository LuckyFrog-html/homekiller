import { api } from '$lib/api';
import type { Actions, PageServerLoad } from './$types';
import type { Solution, Student, Task } from '$lib/types';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { formSchema } from './schema';
import { fail, redirect } from '@sveltejs/kit';

export let load: PageServerLoad = async function load({ params, cookies }) {
    const token = cookies.get('teacher_token');
    const homeworkRes = await api.get<Task>(`/teacher/homeworks/${params.id}`, { token });

    if (homeworkRes.type === "error" && homeworkRes.status === 401) {
        return redirect(303, '/login');
    }

    if (homeworkRes.type === "networkerror" || homeworkRes.type === "error") {
        return redirect(303, '/login');
    }

    const task = homeworkRes.data;

    type SolvesRes = { solutions: Solution[] };
    const solutionsRes = await api.get<SolvesRes>(`/homeworks/${params.id}/solves`, { token });

    if (solutionsRes.type === "networkerror" || solutionsRes.type === "error") {
        return redirect(303, '/login');
    }

    let solutions = solutionsRes.data.solutions;

    //TODO:: this should be returned from api by default
    const studentsReq = await api.get<{ students: Student[] }>('/teacher/students', { token });
    if (studentsReq.type === "networkerror" || studentsReq.type === "error") {
        return redirect(303, '/login');
    }
    const studentMap: Record<number, Student> = {};
    for (const student of studentsReq.data.students) {
        studentMap[student.ID] = student;
    }
    for (const solution of solutions) {
        solution.Student = studentMap[solution.StudentID];
    }

    return {
        task,
        solutions,
        id: params.id,
        form: await superValidate(zod(formSchema)),
    };
}

export const actions: Actions = {
    default: async (event) => {
        const form = await superValidate(event, zod(formSchema));
        if (!form.valid) {
            return fail(400, { form });
        }

        console.log(form);

        return {
            form,
        };
    },
};


