<script lang="ts">
    import * as Form from "$lib/components/ui/form";
    import type { Solution } from "$lib/types";
    import type { Infer, SuperValidated } from "sveltekit-superforms";
    import { zodClient } from "sveltekit-superforms/adapters";
    import { superForm } from "sveltekit-superforms/client";
    import { formSchema, type FormSchema } from "./schema";
    import { Input } from "$lib/components/ui/input";

    export let solution: Solution;
    export let data: SuperValidated<Infer<FormSchema>>;

    const form = superForm(data, {
        validators: zodClient(formSchema),
        id: solution.ID.toString(),
    });

    const { form: formData, enhance, formId } = form;

    $: $formData.score = +$formData.score;
</script>

<div class="bg-slate-200 dark:bg-slate-900 p-6">
    <div class="flex flex-row gap-2">
        <div class="w-full">
            {#each solution.Text.split("\n") as paragraph}
                <p>{paragraph}</p>
            {/each}
        </div>
        <div>
            <p>{solution.Student?.Name}</p>
            {#if solution.HomeworkAnswerFiles}
                {#each solution.HomeworkAnswerFiles as file}
                    <a
                        class="text-xl p-2 dark:bg-slate-800 bg-slate-300 rounded-lg w-fit"
                        href={`/${file.Filepath}`}
                        download
                    >
                        {file.Filepath.split("/").pop()}
                    </a>
                {/each}
            {/if}
        </div>
    </div>

    <form method="POST" use:enhance>
        <Form.Field {form} name="score">
            <Form.Control let:attrs>
                <Form.Label>Оценка</Form.Label>
                <Input {...attrs} type="number" bind:value={$formData.score} />
            </Form.Control>
            <Form.FieldErrors />
        </Form.Field>
        <Form.Field {form} name="comment">
            <Form.Control let:attrs>
                <Form.Label>Комментарий</Form.Label>
                <Input {...attrs} bind:value={$formData.comment} />
            </Form.Control>
            <Form.FieldErrors />
        </Form.Field>
        <div class="flex justify-between items-center">
            <Form.Button>Отправить</Form.Button>
        </div>
    </form>
</div>
