<script lang="ts">
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    import type { Group, Homework, Lesson, Student } from "$lib/types";
    import { parseDateFromString } from "$lib/utils";

    /* @type {import('./$types').PageData} */
    // export let data;
    // const groups = data.tasks;

    const lesson: Lesson = {
        ID: 1,
        CreatedAt: "2024-03-26 16:13:41.557+07",
        DeletedAt: null,
        UpdatedAt: "2024-03-26 16:13:41.557+07",
        Date: "2024-03-26 16:13:41.557+07",
        Group: {
            ID: 1,
            CreatedAt: "2024-03-26T16:33:43.335+07:00",
            UpdatedAt: "2024-03-26T16:33:43.335+07:00",
            DeletedAt: null,
            Title: "├Ёєяяр яєяёшъют",
            IsActive: true,
        },
    };

    const students: Student[] = [
        {
            ID: 1,
            CreatedAt: "2024-03-26 16:13:41.557+07",
            UpdatedAt: "2024-03-26 16:13:41.557+07",
            DeletedAt: null,
            Login: "student0",
            Name: "Вася",
            Stage: 1,
        },
    ];

    const homeworks: Homework[] = [
        {
            ID: 1,
            CreatedAt: "2024-03-26 16:13:41.557+07",
            DeletedAt: null,
            UpdatedAt: "2024-03-26 16:13:41.557+07",
            Deadline: "2024-03-26 16:13:41.557+07",
            Description:
                "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Deleniti itaque voluptate recusandae placeat commodi hic sapiente cumque quia ipsa ipsam?",
            MaxScore: 10,
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
    <h2 class="text-3xl mt-3">
        Урок {lesson.ID} у группы {lesson.Group?.Title}
    </h2>

    <main class="grid grid-cols-2 w-full p-3 gap-3">
        <div class="flex flex-col gap-3">
            <h1 class="text-3xl">Студенты:</h1>
            {#each students as student}
                <h4
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                >
                    {student.Name}
                </h4>
            {/each}
        </div>

        <!-- <h1 class="text-3xl">Сделано:</h1> -->

        <div class="flex flex-col gap-3">
            <h1 class="text-3xl">Домашки:</h1>

            {#each homeworks as homework}
                <!-- {#if !task.IsDone && (selectedGroup == "-1" || selectedGroup == task.GroupId.toString())} -->
                <a
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                    href="/teacher/homework/{homework.ID}"
                >
                    <p class="text-xl">
                        {parseDateFromString(homework.Deadline)}
                    </p>
                    <p class="text-xl">{homework.Description}</p>
                    <!-- <p class="overflow-ellipsis line-clamp-1"> -->
                    <!-- {group.Description} -->
                    <!-- </p> -->
                </a>
                <!-- {/if} -->
            {/each}
        </div>
    </main>
</div>
