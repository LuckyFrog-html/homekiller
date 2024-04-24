import { z } from "zod";

export const formSchema = z.object({
    score: z.number(),
    comment: z.string(),
});

export type FormSchema = typeof formSchema;

