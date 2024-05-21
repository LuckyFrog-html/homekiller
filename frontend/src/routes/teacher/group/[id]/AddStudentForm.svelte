<script lang="ts">
    import SuperDebug, {
        superForm,
        type Infer,
        type SuperValidated,
    } from "sveltekit-superforms";
    import { formSchema } from "./schema";
    import type { FormSchema } from "./schema";
    import { zodClient } from "sveltekit-superforms/adapters";
    import * as Form from "$lib/components/ui/form";
    import type { Student } from "$lib/types";
    import { Checkbox } from "$lib/components/ui/checkbox";

    export let data: SuperValidated<Infer<FormSchema>>;
    export let allStudents: Student[] = [];
    export let groupStudents: Student[] = [];
    $: groupStudentsIds = groupStudents.map((s) => s.ID);
    $: leftStudents = allStudents.filter(
        (s) => !groupStudentsIds.includes(s.ID),
    );

    const form = superForm(data, {
        validators: zodClient(formSchema),
    });
    const { form: formData, enhance, errors } = form;
    $: console.log($errors);
</script>

<form method="POST" use:enhance enctype="multipart/form-data">
    <Form.Fieldset {form} name="studentIds" class="space-y-0">
        <div class="space-y-2">
            {#each leftStudents as student}
                {@const checked = $formData.studentIds.includes(student.ID)}
                <div class="flex flex-row items-start space-x-3">
                    <Form.Control let:attrs>
                        <Checkbox
                            {...attrs}
                            {checked}
                            onCheckedChange={(v) => {
                                if (v) {
                                    $formData.studentIds = [
                                        ...$formData.studentIds,
                                        student.ID,
                                    ];
                                } else {
                                    $formData.studentIds =
                                        $formData.studentIds.filter(
                                            (id) => id !== student.ID,
                                        );
                                }
                            }}
                        />
                        <Form.Label class="font-normal">
                            {student.Name}
                        </Form.Label>
                        <input
                            hidden
                            type="checkbox"
                            name={attrs.name}
                            value={student.ID}
                            {checked}
                        />
                    </Form.Control>
                </div>
            {/each}
            <Form.FieldErrors />
        </div>
    </Form.Fieldset>
    <Form.Button>Добавить</Form.Button>
</form>
