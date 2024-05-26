import { z } from "zod";

export const groupFormSchema = z.object({
    name: z.string().min(1, "Описание обязательно"),
});

export const studentFormSchema = z.object({
    name: z.string().min(1, "Имя обязательно"),
    stage: z.number().min(1).max(11),
    login: z.string().min(1, "Логин обязателен"),
    password: z.string().min(1, "Пароль обязателен"),
});

export type GroupFormSchema = typeof groupFormSchema;
export type StudentFormSchema = typeof studentFormSchema;

