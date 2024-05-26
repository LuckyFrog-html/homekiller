<script lang="ts">
    import SuperDebug, {
        dateProxy,
        intProxy,
        superForm,
        type Infer,
        type SuperValidated,
    } from "sveltekit-superforms";
    import { lessonFormSchema } from "./schema";
    import type { LessonFormSchema } from "./schema";
    import { zodClient } from "sveltekit-superforms/adapters";
    import * as Form from "$lib/components/ui/form";
    import Input from "$lib/components/ui/input/input.svelte";

    export let data: SuperValidated<Infer<LessonFormSchema>>;
    const form = superForm(data, {
        validators: zodClient(lessonFormSchema),
    });
    const { form: formData, enhance, errors } = form;
    let hour = intProxy(form, "hour");
    let minute = intProxy(form, "minute");
    let date = dateProxy(form, "date", { format: "date" });
    $: console.log(date, $formData);
</script>

<SuperDebug data={$formData} />

<form
    method="POST"
    action="?/addLesson"
    use:enhance
    enctype="multipart/form-data"
>
    <Form.Field {form} name="date" class="space-y-3">
        <Form.Control let:attrs>
            <Input type="date" {...attrs} bind:value={date} />
        </Form.Control>
        <Form.FieldErrors />
    </Form.Field>

    <Form.Field {form} name="hour" class="space-y-3">
        <Form.Control let:attrs>
            <Input type="number" {...attrs} bind:value={hour} />
        </Form.Control>
        <Form.FieldErrors />
    </Form.Field>

    <Form.Field {form} name="minute" class="space-y-3">
        <Form.Control let:attrs>
            <Input type="number" {...attrs} bind:value={minute} />
        </Form.Control>
        <Form.FieldErrors />
    </Form.Field>

    <Form.Button>Добавить</Form.Button>
</form>
