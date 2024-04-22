import { api } from "$lib/api";
import { redirect } from "@sveltejs/kit";
import type { PageServerLoad } from "./$types";

/** @type {PageServerLoad} */
export async function load({ cookies }: Parameters<PageServerLoad>[0]): Promise<ReturnType<PageServerLoad>> {
    const token = cookies.get('token');
    const req = await api.get<any>('/homeworks', { token });


    console.log(req);
    if (req.type === "success") {
        return {
            tasks: req.data || []
        }
    }


    return redirect(302, '/login');
}
