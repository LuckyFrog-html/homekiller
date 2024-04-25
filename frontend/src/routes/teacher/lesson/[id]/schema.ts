import { z } from "zod";

export const formSchema = z.object({
    description: z.string().min(1, "Описание обязательно"),
    files: z
        .custom<File>((value) => {
            if (!value) return true;
            return true;
        })
        .array(),
});

export type FormSchema = typeof formSchema;

