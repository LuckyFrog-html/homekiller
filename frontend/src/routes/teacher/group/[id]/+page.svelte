<script lang="ts">
    import { parseDateFromString } from "$lib/utils";
    import type { PageData } from "./$types";
    import AddStudentDialog from "./AddStudentDialog.svelte";

    export let data: PageData;
    $: students = data.students;
    $: lessons = data.lessons;
    $: group = data.group;
    $: allStudents = data.allStudents;
</script>

<div class="flex flex-col items-center h-full w-full">
    <div class="flex flex-row gap-10">
        <a href="/teacher/groups" class="text-3xl">&lt;-</a>
        <h2 class="text-3xl">{group.Title}</h2>
    </div>

    <main class="grid grid-cols-2 w-full p-3 gap-3">
        <div class="flex flex-col gap-3">
            <h1 class="text-2xl">Студенты:</h1>
            {#each students as student}
                <h4
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                >
                    {student.Name}
                </h4>
            {/each}
        </div>

        <div class="flex flex-col gap-3">
            <h1 class="text-2xl">Уроки:</h1>

            {#each lessons as lesson}
                <a
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                    href="/teacher/lesson/{lesson.ID}"
                >
                    <h4 class="text-xl">{parseDateFromString(lesson.Date)}</h4>
                </a>
            {/each}
        </div>

        <AddStudentDialog
            groupStudents={students}
            {allStudents}
            data={data.form}
        />
    </main>
</div>
