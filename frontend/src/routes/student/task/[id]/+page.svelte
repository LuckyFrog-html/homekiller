<script lang="ts">
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
</script>

<div class="flex h-full w-full">
    <div class="flex w-full flex-col gap-2 p-6">
        <div>
            <a href="/student/tasks" class="text-3xl">&lt;-</a>
            <span class="text-3xl">{data.task.ID}</span>
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
    </div>
</div>
