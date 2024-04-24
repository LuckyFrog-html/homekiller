<script lang="ts">
    import type { PageData } from "./$types";
    import Solution from "./solution.svelte";

    export let data: PageData;
    const task = data.task;
    const solutions = data.solutions;
    const files = task.HomeworkFiles || [];
</script>

<div class="flex h-full w-full">
    <div class="flex w-full flex-col gap-2 p-6">
        <div>
            <a href="/teacher/lesson/{task.LessonID}" class="text-3xl">&lt;-</a>
            <span class="text-3xl">{data.task.ID}</span>
        </div>
        <main class="flex lg:flex-row flex-col gap-3 mb-3">
            <div class="flex flex-col w-full grow justify-between">
                <h1 class="text-3xl mb-3">Задание</h1>
                <div>
                    {#each data.task.Description.split("\n") as paragraph}
                        <p>{paragraph}</p>
                    {/each}
                </div>
                <div class="flex flex-row gap-2 justify-self-end">
                    {#each files as file}
                        <a
                            class="text-xl p-2 dark:bg-slate-800 bg-slate-300 rounded-lg w-fit"
                            href={`/${file.Filepath}`}
                            download
                        >
                            {file.Filepath.split("/").pop()}
                        </a>
                    {/each}
                </div>
            </div>
        </main>
        {#if solutions}
            <h2 class="text-2xl">Решения</h2>
            {#each solutions as solution}
                <Solution {solution} data={data.form} />
            {/each}
        {:else}
            <h2>Нет решений</h2>
        {/if}
    </div>
</div>
