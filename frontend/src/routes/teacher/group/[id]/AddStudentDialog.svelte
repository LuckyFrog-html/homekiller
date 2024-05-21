<script lang="ts">
    import * as Dialog from "$lib/components/ui/dialog";
    import { type Infer, type SuperValidated } from "sveltekit-superforms";
    import type { FormSchema } from "./schema";
    import { Button } from "$lib/components/ui/button";
    import type { Student } from "$lib/types";
    import AddStudentForm from "./AddStudentForm.svelte";

    export let data: SuperValidated<Infer<FormSchema>>;
    export let allStudents: Student[] = [];
    export let groupStudents: Student[] = [];
</script>

<Dialog.Root>
    <Dialog.Trigger>
        <Button>Добавить учеников</Button>
    </Dialog.Trigger>
    <Dialog.Content>
        <Dialog.Header>
            <Dialog.Title>Добавить учеников</Dialog.Title>
            <Dialog.Description>
                <AddStudentForm {data} {allStudents} {groupStudents} />
            </Dialog.Description>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
