import { z } from "zod";

export const formSchema = z.object({
    name: z.string().min(1, "Описание обязательно"),
});

export type FormSchema = typeof formSchema;

