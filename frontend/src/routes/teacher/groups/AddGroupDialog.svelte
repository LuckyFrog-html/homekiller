<script lang="ts">
    import * as Dialog from "$lib/components/ui/dialog";
    import {
        superForm,
        type Infer,
        type SuperValidated,
    } from "sveltekit-superforms";
    import { groupFormSchema } from "./schema";
    import type { GroupFormSchema } from "./schema";
    import { zodClient } from "sveltekit-superforms/adapters";
    import { Textarea } from "$lib/components/ui/textarea";
    import * as Form from "$lib/components/ui/form";
    import { Button } from "$lib/components/ui/button";

    export let data: SuperValidated<Infer<GroupFormSchema>>;

    const form = superForm(data, {
        validators: zodClient(groupFormSchema),
    });
    const { form: formData, enhance, errors } = form;
</script>

<Dialog.Root>
    <Dialog.Trigger>
        <Button>Добавить группу</Button>
    </Dialog.Trigger>
    <Dialog.Content>
        <Dialog.Header>
            <Dialog.Title>Добавить группу</Dialog.Title>
            <Dialog.Description>
                <form
                    method="POST"
                    action="?/addGroup"
                    use:enhance
                    enctype="multipart/form-data"
                >
                    <div class="flex flex-col">
                        <Form.Field {form} name="name">
                            <Form.Control let:attrs>
                                <Form.Label class="text-xl">Название</Form.Label
                                >
                                <Textarea
                                    {...attrs}
                                    bind:value={$formData.name}
                                />
                            </Form.Control>
                            <Form.FieldErrors />
                        </Form.Field>
                        <Form.Button class="w-44 mt-2">Добавить</Form.Button>
                    </div>
                </form>
            </Dialog.Description>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
