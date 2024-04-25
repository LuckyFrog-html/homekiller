import { api } from "$lib/api";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";
import type { Group } from "lucide-svelte";

/** @type {PageServerLoad} */
export async function load({ cookies }: Parameters<PageServerLoad>[0]) {
    const token = cookies.get('teacher_token');
    const req = await api.get<{ groups: Group[] }>('/teacher/groups', { token });

    console.log(req);

    if (req.type === "success") {
        const groups = req.data.groups as Group[] || [];
        return { groups };
    }

    return redirect(302, '/login');
}
