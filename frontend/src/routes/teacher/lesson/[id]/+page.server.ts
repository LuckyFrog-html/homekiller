import { error, type Actions, fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from "./$types";
import { api } from '$lib/api';
import type { Homework, Lesson, Student } from '$lib/types';
import type { Group } from 'lucide-svelte';

/** @type {PageServerLoad} */
export async function load({ params, cookies }: Parameters<PageServerLoad>[0]) {
    const token = cookies.get('teacher_token');

    const lessonRes = await api.get<{ lessons: Lesson[] }>(`/teacher/lessons`, { token });
    console.log(lessonRes);

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
        lesson
    };
}

