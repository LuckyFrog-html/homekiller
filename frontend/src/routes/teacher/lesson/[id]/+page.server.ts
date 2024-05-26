import { error, type Actions, fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from "./$types";
import { api } from '$lib/api';
import type { Homework, Lesson, Student } from '$lib/types';
import type { Group } from 'lucide-svelte';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { formSchema } from './schema';

/** @type {PageServerLoad} */
export async function load({ params, cookies }: Parameters<PageServerLoad>[0]) {
    const token = cookies.get('teacher_token');

    const lessonRes = await api.get<{ lessons: Lesson[] }>(`/teacher/lessons`, { token });

    if (lessonRes.type === "error" && lessonRes.status === 401) {
        return redirect(303, '/login');
    }

    if (lessonRes.type === "networkerror" || lessonRes.type === "error") {
        return redirect(303, '/login');
    }

    const lessons = lessonRes.data.lessons;
    const lesson = lessons.filter((lesson) => lesson.ID === +params.id)[0];

    const homeworksRes = await api.get<{ homeworks: Homework[] }>(`/lessons/${params.id}/homeworks`, { token });

    if (homeworksRes.type === "error" && homeworksRes.status === 401) {
        return redirect(303, '/login');
    }

    if (homeworksRes.type === "networkerror" || homeworksRes.type === "error") {
        return redirect(303, '/login');
    }

    const homeworks = homeworksRes.data.homeworks;

    return {
        homeworks,
        lesson,
        form: await superValidate(zod(formSchema)),
    };
}

export const actions: Actions = {
    default: async (event) => {
        const form = await superValidate(event, zod(formSchema));
        if (!form.valid) {
            return fail(400, { form });
        }

        const lessonId = event.params.id;
        if (lessonId === undefined) {
            return fail(400, {
                form,
            });
        }

        const token = event.cookies.get('teacher_token');
        const res = await api.post<any>(`/homeworks`, {
            lesson_id: +lessonId,
            description: form.data.description,
            deadline: '2025-04-21T23:42:51.090281+07:00',
            max_score: 10,
        }, { token });


        if (res.type === 'error' && res.status === 401) {
            return redirect(303, '/login');
        }

        if (res.type === "networkerror" || res.type === "error") {
            return redirect(303, '/login');
        }

        const files = form.data.files.filter((file) => !!file);
        if (form.data.files.length > 0) {
            for (const file of files) {
                const fileRes = await fetch(api.url + `/homeworks/${res.data.ID}/files`, {
                    method: 'POST',
                    headers: {
                        'Authorization': `Bearer ${token}`,
                        'Content-Type': file.type,
                        'Content-Disposition': `attachment; filename=${file.name}`,
                    },
                    body: file,
                });
            }
        }

        form.data.files = [];

        return {
            form,
        };
    },
};

