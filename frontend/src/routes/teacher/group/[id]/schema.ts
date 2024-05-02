import { z } from "zod";

export const formSchema = z.object({
    studentIds: z.number().array().min(1, "Выберите студентов"),
})

export type FormSchema = typeof formSchema;
