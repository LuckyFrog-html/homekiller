import { error, type Actions, fail } from '@sveltejs/kit';
import type { PageServerLoad } from "./$types";
import { superValidate } from 'sveltekit-superforms';
import { zod } from 'sveltekit-superforms/adapters';
import { formSchema } from './schema';

/** @type {PageServerLoad} */
export async function load({ params }: Parameters<PageServerLoad>[0]): Promise<ReturnType<PageServerLoad>> {
    const task = {
        id: params.id,
        description: "lorem ipsum dolor sit amet consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
        subject: "Math",
        completed: false,
    }

    return {
        task,
        form: await superValidate(zod(formSchema)),
    };

    error(404, 'Not found');
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


