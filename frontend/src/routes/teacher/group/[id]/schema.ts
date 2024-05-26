import { z } from "zod";

export const formSchema = z.object({
    studentIds: z.number().array().min(1, "Выберите студентов"),
})

export const lessonFormSchema = z.object({
    date: z.date().min(new Date(), "Нельзя добавлять прошедшие занятия").optional(),
    hour: z.number().min(0).max(23),
    minute: z.number().min(0).max(59),
});

export type FormSchema = typeof formSchema;

export type LessonFormSchema = typeof lessonFormSchema;
