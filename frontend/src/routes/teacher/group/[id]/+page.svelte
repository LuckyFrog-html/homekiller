<script lang="ts">
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    import type { Group, Lesson, Student } from "$lib/types";
    import { parseDateFromString } from "$lib/utils";

    /* @type {import('./$types').PageData} */
    // export let data;
    // const groups = data.tasks;

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

<div class="flex h-full w-full">
    <main class="flex h-full w-full p-3 gap-3">
        <h1 class="text-3xl">Ваши группы:</h1>

        <div class="flex flex-col gap-3">
            {#each students as student}
                <!-- {#if !task.IsDone && (selectedGroup == "-1" || selectedGroup == task.GroupId.toString())} -->
                <a
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                    href="/teacher/group/{student.ID}"
                >
                    <h4 class="text-xl">{student.Name}</h4>
                    <!-- <p class="overflow-ellipsis line-clamp-1"> -->
                    <!-- {group.Description} -->
                    <!-- </p> -->
                </a>
                <!-- {/if} -->
            {/each}
        </div>

        <!-- <h1 class="text-3xl">Сделано:</h1> -->

        <div class="flex flex-col gap-3">
            {#each lessons as lesson}
                <!-- {#if !task.IsDone && (selectedGroup == "-1" || selectedGroup == task.GroupId.toString())} -->
                <a
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                    href="/teacher/group/{lesson.ID}"
                >
                    <h4 class="text-xl">{parseDateFromString(lesson.Date)}</h4>
                    <!-- <p class="overflow-ellipsis line-clamp-1"> -->
                    <!-- {group.Description} -->
                    <!-- </p> -->
                </a>
                <!-- {/if} -->
            {/each}
        </div>
    </main>
</div>
