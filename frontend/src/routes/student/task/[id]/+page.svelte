<script lang="ts">
    import Solution from "../../../teacher/homework/[id]/solution.svelte";
    import type { PageData } from "./$types";
    import Form from "./form.svelte";
    export let data: PageData;

    function isImage(file: string) {
        return (
            file.endsWith(".png") ||
            file.endsWith(".jpg") ||
            file.endsWith(".jpeg") ||
            file.endsWith(".gif")
        );
    }

    const files = data.task.HomeworkFiles.filter(
        (file) => !isImage(file.Filepath),
    );

    const images = data.task.HomeworkFiles.filter((file) =>
        isImage(file.Filepath),
    );

    $: solutions = data.solutions || [];
</script>

<div class="flex h-full w-full p-5">
    <div class="flex w-full flex-col gap-2">
        <div class="flex gap-6 justify-center text-3xl">
            <a class="rounded border border-current border-solid px-2 py-1" href="/student/tasks">&lt;- Назад</a>
            <span> Домашнее задание номер {data.task.ID}</span>
        </div>
        <main class="flex lg:flex-row flex-col gap-3">
            <div class="flex flex-col w-full grow justify-between">
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
            <div class="max-w-[800px] w-full">
                {#each images as image}
                    <img
                        class="w-full max-h-[500px] object-contain"
                        src={`/${image.Filepath}`}
                        alt={image.Filepath}
                    />
                {/each}
            </div>
        </main>
        <div class="max-w-80">
            <Form data={data.form} />
        </div>
        <div>
            {#each solutions as solution}
                <div class="p-2 rounded bg-slate-200 dark:bg-slate-700">
                    <div
                        class="dark:bg-slate-800 bg-slate-300 rounded-lg p-3 my-3"
                    >
                        Ответ: {solution.Text}
                    </div>
                    {#each solution.Reviews || [] as review}
                        <div
                            class="dark:bg-slate-800 bg-slate-300 rounded-lg p-3 my-3"
                        >
                            Оценка:
                            <p>{review.Score}</p>
                            Комментарий:
                            <p>{review.Comment}</p>
                        </div>
                    {/each}
                </div>
            {/each}
        </div>
    </div>
</div>
