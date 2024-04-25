<script lang="ts">
    import * as Dialog from "$lib/components/ui/dialog";
    import {
        superForm,
        type Infer,
        type SuperValidated,
    } from "sveltekit-superforms";
    import { formSchema } from "./schema";
    import type { FormSchema } from "./schema";
    import { zodClient } from "sveltekit-superforms/adapters";
    import { Textarea } from "$lib/components/ui/textarea";
    import * as Form from "$lib/components/ui/form";
    import { invalidateAll } from "$app/navigation";
    import { Button } from "$lib/components/ui/button";

    export let data: SuperValidated<Infer<FormSchema>>;

    const form = superForm(data, {
        validators: zodClient(formSchema),
    });
    const { form: formData, enhance, errors } = form;
</script>

<Dialog.Root>
    <Dialog.Trigger>
        <Button>Добавить задание</Button>
    </Dialog.Trigger>
    <Dialog.Content>
        <Dialog.Header>
            <Dialog.Title>Добавить задание</Dialog.Title>
            <Dialog.Description>
                <form method="POST" use:enhance enctype="multipart/form-data">
                    <div class="flex flex-col">
                        <Form.Field {form} name="description">
                            <Form.Control let:attrs>
                                <Form.Label class="text-xl">Описание</Form.Label
                                >
                                <Textarea
                                    {...attrs}
                                    bind:value={$formData.description}
                                />
                            </Form.Control>
                            <Form.FieldErrors />
                        </Form.Field>
                        <input
                            type="file"
                            multiple
                            name="files"
                            on:input={(e) =>
                                ($formData.files = Array.from(
                                    e.currentTarget.files ?? [],
                                ))}
                        />
                        {#if $errors.files}<span>{$errors.files[0]}</span>{/if}

                        <Form.Button class="w-44 mt-2">Добавить</Form.Button>
                    </div>
                </form>
            </Dialog.Description>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
