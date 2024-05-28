<script lang="ts">
    import type { PageData } from "./$types";
    import AddGroupDialog from "./AddGroupDialog.svelte";
    import AddStudentDialog from "./AddStudentDialog.svelte";

    export let data: PageData;
    $: groups = data.groups;
    $: students = data.students;
</script>

<div class="flex h-full w-full">
    <main class="flex w-full flex-col lg:flex-row p-3 gap-3">
        <div class="w-full">
            <h1 class="text-3xl mb-3">Ваши группы:</h1>
            <div class="flex flex-col gap-3 max-h-screen">
                {#each groups as group}
                    <a
                        class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                        href="/teacher/group/{group.ID}"
                    >
                        <h4 class="text-xl">{group.Title}</h4>
                    </a>
                {/each}
                <AddGroupDialog data={data.groupForm} />
            </div>
        </div>
        <div class="w-full">
            <h1 class="text-3xl mb-3">Все студенты:</h1>
            <div class="flex flex-col gap-3 max-h-screen">
                {#each students as student}
                    <div
                        class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                    >
                        <h4 class="text-xl">{student.Name}, {student.Stage}</h4>
                    </div>
                {/each}
                <AddStudentDialog data={data.studentForm} />
            </div>
        </div>
    </main>
</div>
