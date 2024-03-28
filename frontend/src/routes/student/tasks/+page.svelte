<script lang="ts">
    import * as ToggleGroup from "$lib/components/ui/toggle-group";
    type Task = {
        description: string;
        subject: string;
        completed: boolean;
    };

    const tasks: Task[] = [
        {
            description:
                "Lorem, ipsum dolor sit amet consectetur adipisicing elit. Expedita soluta sed quis porro suscipit illum cupiditate tenetur maxime, ex quasi, numquam, quidem reprehenderit! Id facilis nulla modi ipsam veniam quidem natus laboriosam dolorum delectus amet aliquam nobis assumenda recusandae harum, voluptatem atque libero maiores accusamus fugit ex, dolore dolor adipisci voluptates? Doloribus suscipit saepe pariatur veritatis molestias, vel corrupti asperiores id, mollitia, harum sequi ipsam totam est! Aperiam magnam ducimus nihil doloremque, excepturi pariatur voluptates, nam quaerat labore reiciendis, debitis consequuntur asperiores animi nisi odit autem. Consequuntur, quas, vitae nobis suscipit illum laudantium dolor repellendus deleniti natus quidem, voluptatibus quae.",
            subject: "Math",
            completed: false,
        },
        {
            description:
                "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
            subject: "Math",
            completed: true,
        },
        {
            description:
                "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
            subject: "informatics",
            completed: false,
        },
        {
            description:
                "Lorem ipsum dolor sit amet, consectetur adipiscing elit",
            subject: "Math",
            completed: true,
        },
    ];

    const subjects = ["Все", ...new Set(tasks.map((task) => task.subject))];

    let selectedSubject = subjects[0];
</script>

<div class="flex h-full w-full">
    <div class="flex flex-col gap-2 p-3 w-96">
        <ToggleGroup.Root
            variant="outline"
            class="flex flex-col gap-2 p-3 w-96"
            bind:value={selectedSubject}
        >
            {#each subjects as subject}
                <ToggleGroup.Item class="w-full" value={subject}>
                    {subject}
                </ToggleGroup.Item>
            {/each}
        </ToggleGroup.Root>
    </div>
    <main class="flex h-full w-full flex-col p-3 gap-3">
        <h1 class="text-3xl">Нужно сделать:</h1>

        {#each tasks as task}
            {#if !task.completed && (selectedSubject == "Все" || selectedSubject == task.subject)}
                <div
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                >
                    <h4 class="text-xl">{task.subject}</h4>
                    <p class="overflow-ellipsis line-clamp-1">
                        {task.description}
                    </p>
                </div>
            {/if}
        {/each}

        <h1 class="text-3xl">Сделано:</h1>

        {#each tasks as task}
            {#if task.completed && (selectedSubject == "Все" || selectedSubject == task.subject)}
                <div
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                >
                    <h4 class="text-xl">{task.subject}</h4>
                    <p>{task.description}</p>
                </div>
            {/if}
        {/each}
    </main>
</div>
