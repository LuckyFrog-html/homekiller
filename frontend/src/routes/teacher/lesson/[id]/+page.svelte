<script lang="ts">
    import { invalidateAll } from "$app/navigation";
    import Button from "$lib/components/ui/button/button.svelte";
    import { parseDateFromString } from "$lib/utils";
    import AddHomeworkDialog from "./AddHomeworkDialog.svelte";

    /* @type {import('./$types').PageData} */
    export let data;
    $: homeworks = data.homeworks;
    const lesson = data.lesson;
</script>

<div class="flex mt-5 flex-col items-center h-full w-full">
    <div class="flex gap-10 relative text-3xl">
        <a class="rounded border border-current border-solid px-2 py-1" href="/teacher/group/{lesson.GroupID}">&lt;- Назад</a>
        <h2>
            Урок {lesson.ID} у группы {lesson.Group?.Title}
        </h2>
    </div>

    <main class="grid grid-cols-2 w-full p-3 gap-3">
        <div class="flex flex-col gap-3">
            <h1 class="text-3xl">Домашки:</h1>

            {#each homeworks as homework}
                <a
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                    href="/teacher/homework/{homework.ID}"
                >
                    <p class="text-xl">
                        {parseDateFromString(homework.Deadline)}
                    </p>
                    <p class="text-xl">{homework.Description}</p>
                </a>
            {/each}
            <AddHomeworkDialog data={data.form} />
        </div>
    </main>
</div>
