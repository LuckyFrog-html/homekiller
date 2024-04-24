import { z } from "zod";

export const formSchema = z.object({
    login: z.string()
        .min(2, "Логин должен быть не короче 2 символов")
        .max(50, "Логин должен быть не длиннее 50 символов")
        // TODO: Сделать более жесткую проверку
        .refine((s) => s.indexOf(" ") === -1, "Логин не должен содержать пробелов"),
    password: z.string()
        .min(2, "Пароль должен быть не короче 2 символов")
        .max(50, "Пароль должен быть не длиннее 50 символов")
        // TODO: Сделать более жесткую проверку
        .refine((s) => s.indexOf(" ") === -1, "Пароль не должен содержать пробелов"),
});

export type FormSchema = typeof formSchema;

