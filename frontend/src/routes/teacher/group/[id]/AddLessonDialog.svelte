<script lang="ts">
    import * as Dialog from "$lib/components/ui/dialog";
    import { type Infer, type SuperValidated } from "sveltekit-superforms";
    import type { FormSchema } from "./schema";
    import { Button } from "$lib/components/ui/button";
    import AddLessonForm from "./AddLessonForm.svelte";

    export let data: SuperValidated<Infer<FormSchema>>;
</script>

<Dialog.Root>
    <Dialog.Trigger>
        <Button>Добавить занятие</Button>
    </Dialog.Trigger>
    <Dialog.Content>
        <Dialog.Header>
            <Dialog.Title>Добавить занятие</Dialog.Title>
            <Dialog.Description>
                <AddLessonForm {data} />
            </Dialog.Description>
        </Dialog.Header>
    </Dialog.Content>
</Dialog.Root>
