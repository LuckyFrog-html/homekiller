<script lang="ts">
    import * as Form from "$lib/components/ui/form";
    import { Textarea } from "$lib/components/ui/textarea";
    import { formSchema, type FormSchema } from "./schema";
    import {
        type SuperValidated,
        type Infer,
        superForm,
    } from "sveltekit-superforms";
    import { zodClient } from "sveltekit-superforms/adapters";

    export let data: SuperValidated<Infer<FormSchema>>;

    const form = superForm(data, {
        validators: zodClient(formSchema),
    });
    const { form: formData, enhance, errors } = form;
</script>

<form method="POST" use:enhance enctype="multipart/form-data">
    <div class="flex flex-col">
        <Form.Field {form} name="answer">
            <Form.Control let:attrs>
                <Form.Label class="text-xl">Ответ</Form.Label>
                <Textarea {...attrs} bind:value={$formData.answer} />
            </Form.Control>
            <Form.FieldErrors />
        </Form.Field>
        <input
            type="file"
            multiple
            name="files"
            accept="image/png, image/jpeg"
            on:input={(e) =>
                ($formData.files = Array.from(e.currentTarget.files ?? []))}
        />
        {#if $errors.files}<span>{$errors.files}</span>{/if}
        <Form.Button class="w-44 mt-2">Отправить</Form.Button>
    </div>
</form>
