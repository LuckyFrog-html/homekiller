<script lang="ts">
    import * as Form from "$lib/components/ui/form";
    import { Input } from "$lib/components/ui/input";
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

    const { form: formData, enhance } = form;
</script>

<form method="POST" use:enhance>
    <Form.Field {form} name="login">
        <Form.Control let:attrs>
            <Form.Label>Логин</Form.Label>
            <Input {...attrs} bind:value={$formData.login} />
        </Form.Control>
        <Form.FieldErrors />
    </Form.Field>
    <Form.Field {form} name="password">
        <Form.Control let:attrs>
            <Form.Label>Пароль</Form.Label>
            <Input {...attrs} bind:value={$formData.password} type="password" />
        </Form.Control>
        <Form.FieldErrors />
    </Form.Field>
    <div class="flex justify-between items-center">
        <Form.Button>Войти</Form.Button>
        <p>
            Учитель? <a href="/login/teacher" class="text-blue-600">Войти</a>
        </p>
    </div>
</form>
