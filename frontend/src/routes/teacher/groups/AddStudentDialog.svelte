<script lang="ts">
    import * as Dialog from "$lib/components/ui/dialog";
    import SuperDebug, {
        intProxy,
        superForm,
        type Infer,
        type SuperValidated,
    } from "sveltekit-superforms";
    import { studentFormSchema } from "./schema";
    import type { StudentFormSchema } from "./schema";
    import { zodClient } from "sveltekit-superforms/adapters";
    import * as Form from "$lib/components/ui/form";
    import { Button } from "$lib/components/ui/button";
    import { Input } from "$lib/components/ui/input";

    export let data: SuperValidated<Infer<StudentFormSchema>>;

    const form = superForm(data, {
        validators: zodClient(studentFormSchema),
    });
    const { form: formData, enhance, errors } = form;
    let stage = intProxy(form, "stage");
</script>

<Dialog.Root>
    <Dialog.Trigger>
        <Button>Добавить студента</Button>
    </Dialog.Trigger>
    <Dialog.Content>
        <Dialog.Header>
            <Dialog.Title>Добавить студента</Dialog.Title>
            <Dialog.Description>
                <form
                    method="POST"
                    action="?/addStudent"
                    use:enhance
                    enctype="multipart/form-data"
                >
                    <div class="flex flex-col">
                        <Form.Field {form} name="name">
                            <Form.Control let:attrs>
                                <Form.Label class="text-xl"
                                    >Имя Фамилия</Form.Label
                                >
                                <Input
                                    type="text"
                                    {...attrs}
                                    bind:value={$formData.name}
                                />
                            </Form.Control>
                            <Form.FieldErrors />
                        </Form.Field>
                        <Form.Field {form} name="stage">
                            <Form.Control let:attrs>
                                <Form.Label class="text-xl">Класс</Form.Label>
                                <Input
                                    type="number"
                                    {...attrs}
                                    bind:value={$stage}
                                />
                            </Form.Control>
                            <Form.FieldErrors />
                        </Form.Field>
                        <Form.Field {form} name="login">
                            <Form.Control let:attrs>
                                <Form.Label class="text-xl">Логин</Form.Label>
                                <Input
                                    type="text"
                                    {...attrs}
                                    bind:value={$formData.login}
                                />
                            </Form.Control>
                            <Form.FieldErrors />
                        </Form.Field>
                        <Form.Field {form} name="password">
                            <Form.Control let:attrs>
                                <Form.Label class="text-xl">Пароль</Form.Label>
                                <Input
                                    type="password"
                                    {...attrs}
                                    bind:value={$formData.password}
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
