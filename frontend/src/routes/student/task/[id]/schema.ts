import { z } from "zod";

export const formSchema = z.object({
    answer: z.string().min(1, "Ответ не может быть пустым"),
})

export type FormSchema = typeof formSchema;
