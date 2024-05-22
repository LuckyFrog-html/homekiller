import { error, type Actions, fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from "./$types";
import { api } from '$lib/api';
import type { Lesson, Student } from '$lib/types';
import type { Group } from 'lucide-svelte';
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { formSchema, lessonFormSchema } from './schema';

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

    const allStudentsRes = await api.get<{ students: Student[] }>(`/teacher/students`, { token });

    if (allStudentsRes.type === "networkerror" || allStudentsRes.type === "error") {
        return redirect(303, '/login');
    }

    const lessons = lessonsRes.data.lessons;
    const students = studentsRes.data.students;
    const group = groupRes.data.groups.filter((group) => group.ID === +params.id)[0];
    const allStudents = allStudentsRes.data.students;

    return {
        students,
        lessons,
        group,
        allStudents,
        studentsForm: await superValidate(zod(formSchema)),
        lessonForm: await superValidate(zod(lessonFormSchema)),
    };
}


export const actions: Actions = {
    addStudents: async (event) => {
        const form = await superValidate(event, zod(formSchema));
        if (!form.valid) {
            return fail(400, { form });
        }
        const students_ids = form.data.studentIds;

        const token = event.cookies.get('teacher_token');
        const addStudetnsRed = await api.post(`/groups/${event.params.id}/students`, {
            students_ids
        }, { token });

        if (addStudetnsRed.type === "error" && addStudetnsRed.status === 401) {
            return redirect(303, '/login');
        }

        if (addStudetnsRed.type === "networkerror" || addStudetnsRed.type === "error") {
            return redirect(303, '/login');
        }

        return {
            studentsForm: form,
        }
    },

    addLesson: async (event) => {
        const form = await superValidate(event, zod(lessonFormSchema));
        if (!form.valid || !form.data.date) {
            return fail(400, { form });
        }

        let dateD = form.data.date;
        dateD.setTime(dateD.getTime() + (form.data.hour * 60 * 60 * 1000) + (form.data.minute * 60 * 1000));
        const date = dateD.toISOString();
        console.log(date);
        const token = event.cookies.get('teacher_token');
        const addStudetnsRed = await api.post(`/lessons`, { date }, { token });

        console.log(addStudetnsRed);
        if (addStudetnsRed.type === "error" && addStudetnsRed.status === 401) {
            return redirect(303, '/login');
        }

        if (addStudetnsRed.type === "networkerror" || addStudetnsRed.type === "error") {
            return redirect(303, '/login');
        }

        return {
            lessonForm: form,
        }
    }
}
