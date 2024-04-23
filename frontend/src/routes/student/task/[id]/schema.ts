import { z } from "zod";

export const formSchema = z.object({
    answer: z.string().min(1, "Ответ не может быть пустым"),
    files: z
        .instanceof(File, { message: 'Please upload a file.' })
        .array()
})

export type FormSchema = typeof formSchema;
