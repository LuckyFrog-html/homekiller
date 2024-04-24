import type { PageServerLoad, Actions } from "./$types";
import { fail } from "@sveltejs/kit";
import { setError, superValidate } from "sveltekit-superforms";
import { zod } from "sveltekit-superforms/adapters";
import { formSchema } from "./schema";
import { api } from "$lib/api";

export const load: PageServerLoad = async () => {
    return {
        form: await superValidate(zod(formSchema)),
    };
};

const MONTH = 1000 * 60 * 60 * 24 * 30;

type tokenResponse = {
    token: string;
}

export const actions: Actions = {
    default: async (event) => {
        const form = await superValidate(event, zod(formSchema));
        if (!form.valid) {
            return fail(400, {
                form,
            });
        }

        const response = await api.post<tokenResponse>('/teachers/login', {
            login: form.data.login,
            password: form.data.password
        })

        if (response.type === "error" && response.status === 401) {
            return setError(form, 'login', 'Неверный логин или пароль');
        }

        if (response.type === "networkerror" || response.type === "error") {
            return setError(form, 'login', 'Неизвестная ошибка, попробуйте снова');
        }

        event.cookies.set("teacher_token", response.data.token, { path: '/', expires: new Date(Date.now() + MONTH) });

        return {
            form,
        };

    },
};


