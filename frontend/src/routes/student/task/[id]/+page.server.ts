import { error, type Actions, fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from "./$types";
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { formSchema } from './schema';
import { api } from '$lib/api';
import type { Task } from '$lib/types';

/** @type {PageServerLoad} */
export async function load({ params, cookies }: Parameters<PageServerLoad>[0]) {
    const token = cookies.get('token');
    const taskReq = await api.get(`/homeworks/${params.id}`, { token });

    if (taskReq.type === 'error' && taskReq.status === 401) {
        return redirect(303, '/login');
    }

    if (taskReq.type === "networkerror" || taskReq.type === "error") {
        return redirect(303, '/login');
    }

    const task = taskReq.data as Task;

    return {
        task,
        form: await superValidate(zod(formSchema)),
    };
}


export const actions: Actions = {
    default: async (event) => {
        const form = await superValidate(event, zod(formSchema));
        if (!form.valid) {
            return fail(400, {
                form,
            });
        }

        console.log(form.data.answer)

        return {
            form,
        };
    },
};


