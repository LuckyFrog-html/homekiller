import { api } from "$lib/api";
import { redirect, type Actions, fail } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { Group } from "lucide-svelte";
import { superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { formSchema } from "./schema";

/** @type {PageServerLoad} */
export async function load({ cookies }: Parameters<PageServerLoad>[0]) {
    const token = cookies.get('teacher_token');
    const req = await api.get<{ groups: Group[] }>('/teacher/groups', { token });

    if (req.type === "success") {
        const groups = req.data.groups as Group[] || [];
        return {
            groups,
            form: await superValidate(zod(formSchema)),
        };
    }

    return redirect(302, '/login');
}

export const actions: Actions = {
    default: async (event) => {
        const form = await superValidate(event, zod(formSchema));

        if (!form.valid) {
            return fail(400, { form });
        }

        const token = event.cookies.get('teacher_token');
        const title = form.data.name;
        const req = await api.post<{ group: Group }>('/groups', { title }, { token });

        if (req.type === "success") {
            return { success: true, group: req.data.group, form };
        }

        return redirect(302, '/login');
    }
};
