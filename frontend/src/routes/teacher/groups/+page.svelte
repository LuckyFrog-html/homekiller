<script lang="ts">
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    import type { Group } from "$lib/types";

    /* @type {import('./$types').PageData} */
    // export let data;
    // const groups = data.tasks;

    // const groupsHash: any = {};
    const groups: Group[] = [
        {
            ID: 1,
            CreatedAt: "2024-03-26T16:33:43.335+07:00",
            UpdatedAt: "2024-03-26T16:33:43.335+07:00",
            DeletedAt: null,
            Title: "├Ёєяяр яєяёшъют",
            IsActive: true,
        },
        {
            ID: 2,
            CreatedAt: "2024-03-26T16:33:43.335+07:00",
            UpdatedAt: "2024-03-26T16:33:43.335+07:00",
            DeletedAt: null,
            Title: "├Ёєяяр яєяёшъют",
            IsActive: true,
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

    let selectedGroup = "-1";
</script>

<div class="flex h-full w-full">
    <!-- <div class="flex flex-col gap-2 p-3 w-96">
        <ToggleGroup.Root
            variant="outline"
            class="flex flex-col gap-2 p-3 w-96"
            bind:value={selectedGroup}
        >
            {#each groups as group}
                <ToggleGroup.Item class="w-full" value={String(group.id)}>
                    {group.title}
                </ToggleGroup.Item>
            {/each}
        </ToggleGroup.Root>
    </div> -->
    <main class="flex h-full w-full flex-col p-3 gap-3">
        <h1 class="text-3xl">Ваши группы:</h1>

        <div class="flex flex-col gap-3 max-h-screen">
            {#each groups as group}
                <!-- {#if !task.IsDone && (selectedGroup == "-1" || selectedGroup == task.GroupId.toString())} -->
                <a
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                    href="/teacher/group/{group.ID}"
                >
                    <h4 class="text-xl">{group.Title}</h4>
                    <!-- <p class="overflow-ellipsis line-clamp-1"> -->
                    <!-- {group.Description} -->
                    <!-- </p> -->
                </a>
                <!-- {/if} -->
            {/each}
        </div>

        <!-- <h1 class="text-3xl">Сделано:</h1> -->

        <!-- {#each tasks as task}
            {#if task.IsDone && (selectedGroup == "-1" || selectedGroup == task.GroupId.toString())}
                <a
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                    href="/student/task/{task.ID}"
                >
                    <h4 class="text-xl">{task.GroupTitle}</h4>
                    <p class="overflow-ellipsis line-clamp-1">
                        {task.Description}
                    </p>
                </a>
            {/if}
        {/each} -->
    </main>
</div>
