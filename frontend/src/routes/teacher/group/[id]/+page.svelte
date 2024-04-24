<script lang="ts">
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    import type { Group, Lesson, Student } from "$lib/types";
    import { parseDateFromString } from "$lib/utils";

    /* @type {import('./$types').PageData} */
    // export let data;
    // const groups = data.tasks;

    const group: Group = {
        ID: 1,
        CreatedAt: "2024-03-26T16:33:43.335+07:00",
        UpdatedAt: "2024-03-26T16:33:43.335+07:00",
        DeletedAt: null,
        Title: "├Ёєяяр яєяёшъют",
        IsActive: true,
    };

    const students: Student[] = [
        {
            ID: 1,
            CreatedAt: "2024-03-26 16:13:41.557+07",
            DeletedAt: null,
            UpdatedAt: "2024-03-26 16:13:41.557+07",
            Login: "student0",
            Name: "Вася",
            Stage: 1,
        },
    ];

    const lessons: Lesson[] = [
        {
            ID: 1,
            CreatedAt: "2024-03-26 16:13:41.557+07",
            DeletedAt: null,
            UpdatedAt: "2024-03-26 16:13:41.557+07",
            Date: "2024-03-26 16:13:41.557+07",
        },
    ];

    // for (const task of groups) {
    //     groupsHash[task.GroupId.toString()] = task.GroupTitle;
    // }

    // let groups = Object.entries(groupsHash).map((x) => ({
    //     id: x[0],
    //     title: x[1],
    // }));

    // groups = [{ id: "-1", title: "Все" }, ...groups];

    //     let selectedGroup = "-1";
</script>

<div class="flex flex-col items-center h-full w-full">
    <h2 class="text-3xl mt-3">Группа: {group.Title}</h2>

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
            <button class="rounded mt-10 border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3">Добавить нового</button>
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
