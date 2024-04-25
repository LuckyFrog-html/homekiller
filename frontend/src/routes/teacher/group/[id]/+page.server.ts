import { error, type Actions, fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from "./$types";
import { api } from '$lib/api';
import type { Lesson, Student } from '$lib/types';
import type { Group } from 'lucide-svelte';

/** @type {PageServerLoad} */
export async function load({ params, cookies }: Parameters<PageServerLoad>[0]) {
    const token = cookies.get('teacher_token');

    const groupRes = await api.get<{ groups: Group[] }>(`/teacher/groups`, { token });

    if (groupRes.type === "error" && groupRes.status === 401) {
        return redirect(303, '/login');
    }

    if (groupRes.type === "networkerror" || groupRes.type === "error") {
        return redirect(303, '/login');
    }

    const studentsRes = await api.get<{ students: Student[] }>(`/groups/${params.id}/students`, { token });

    if (studentsRes.type === "networkerror" || studentsRes.type === "error") {
        return redirect(303, '/login');
    }

    const lessonsRes = await api.get<{ lessons: Lesson[] }>(`/groups/${params.id}/lessons`, { token });

    if (lessonsRes.type === "networkerror" || lessonsRes.type === "error") {
        return redirect(303, '/login');
    }

    const lessons = lessonsRes.data.lessons;
    const students = studentsRes.data.students;
    const group = groupRes.data.groups.filter((group) => group.ID === +params.id)[0];


    return {
        students,
        lessons,
        group,
    };
}


export const actions: Actions = {
    delete: async ({ request, params, cookies }) => {
        const token = cookies.get('teacher_token');
        const res = api.post(`/groups/${params.id}/students`, {
            students_ids: null
        }, { token });
    }
}
