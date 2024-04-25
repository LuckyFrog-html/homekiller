<script lang="ts">
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    import type { Group, Lesson, Student } from "$lib/types";
    import { parseDateFromString } from "$lib/utils";
    import type { PageData } from "./$types";

    export let data: PageData;
    const students = data.students;
    const lessons = data.lessons;
    const group = data.group;
</script>

<div class="flex flex-col items-center h-full w-full">
    <a href="/teacher/groups" class="text-3xl">&lt;-</a>
    <h2 class="text-3xl mt-3">{group.Title}</h2>

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
            <button
                class="rounded mt-10 border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                >Добавить нового</button
            >
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
    </main>
</div>
