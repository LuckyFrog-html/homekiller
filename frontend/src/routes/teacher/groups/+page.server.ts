import { api } from "$lib/api";
import { redirect, type Actions, fail } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { Group } from "lucide-svelte";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { groupFormSchema, studentFormSchema } from "./schema";
import type { Student } from "$lib/types";

/** @type {PageServerLoad} */
export async function load({ cookies }: Parameters<PageServerLoad>[0]) {
    const token = cookies.get('teacher_token');
    const groupsReq = await api.get<{ groups: Group[] }>('/teacher/groups', { token });
    const studentsReq = await api.get<{ students: Student[] }>(`/teacher/students`, { token });

    if (groupsReq.type === "success" && studentsReq.type == "success") {
        const groups = groupsReq.data.groups as Group[] || [];
        const students = studentsReq.data.students as Student[] || [];
        return {
            groups,
            students,
            groupForm: await superValidate(zod(groupFormSchema)),
            studentForm: await superValidate(zod(studentFormSchema)),
        };
    }

    return redirect(302, '/login');
}

export const actions: Actions = {
    addGroup: async (event) => {
        const form = await superValidate(event, zod(groupFormSchema));

        if (!form.valid) {
            return fail(400, { form });
        }

        const token = event.cookies.get('teacher_token');
        const title = form.data.name;
        const req = await api.post<{ group: Group }>('/groups', { title }, { token });

        if (req.type === "success") {
            return { group: req.data.group, groupForm: form };
        }

        return redirect(302, '/login');
    },

    addStudent: async (event) => {
        const form = await superValidate(event, zod(studentFormSchema));

        if (!form.valid) {
            return fail(400, { form });
        }

        const token = event.cookies.get('teacher_token');
        const req = await api.post<{ student: Student }>('/students', form.data, { token });

        if (req.type === "success") {
            return { student: req.data.student, studentForm: form };
        }

        return redirect(302, '/login');
    }
};
