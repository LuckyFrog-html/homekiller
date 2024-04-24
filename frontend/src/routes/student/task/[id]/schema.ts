import { z } from "zod";

export const formSchema = z.object({
    answer: z.string().min(1, "Ответ не может быть пустым"),
    files: z
        .custom<File>((value) => {
            if (!value) return true;
            return true;
        })
        .array(),
})

export type FormSchema = typeof formSchema;
