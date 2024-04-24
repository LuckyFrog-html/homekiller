<script lang="ts">
    import { parseDateFromString } from "$lib/utils";

    /* @type {import('./$types').PageData} */
    export let data;
    const homeworks = data.homeworks;
    const lesson = data.lesson;
</script>

<div class="flex flex-col items-center h-full w-full">
    <a href="/teacher/group/{lesson.GroupID}" class="text-3xl">&lt;-</a>
    <h2 class="text-3xl mt-3">
        Урок {lesson.ID} у группы {lesson.Group?.Title}
    </h2>

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
        </div>
    </main>
</div>
