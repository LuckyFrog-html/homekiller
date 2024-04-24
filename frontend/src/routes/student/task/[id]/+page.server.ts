import { error, type Actions, fail, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from "./$types";
import { setError, superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { formSchema } from './schema';
import { api } from '$lib/api';
import type { Solution, Task } from '$lib/types';

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
            return fail(400, { form });
        }

        const homework_id = event.params.id;
        const text = form.data.answer;
        if (text == undefined || homework_id === undefined) {
            return fail(400, {
                form,
            });
        }

        const token = event.cookies.get('token');
        const res = await api.post<Solution>(`/solutions`, { homework_id: +homework_id, text }, { token });

        console.log(res);

        if (res.type === 'error' && res.status === 401) {
            return redirect(303, '/login');
        }

        if (res.type === "networkerror" || res.type === "error") {
            return redirect(303, '/login');
        }

        const files = form.data.files.filter((file) => !!file);
        if (form.data.files.length > 0) {
            for (const file of files) {
                const fileRes = await fetch(api.url + `/solutions/${res.data.ID}/files`, {
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


