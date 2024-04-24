<script lang="ts">
    import type { PageData } from "./$types";

    export let data: PageData;
    const groups = data.groups;
</script>

<div class="flex h-full w-full">
    <main class="flex h-full w-full flex-col p-3 gap-3">
        <h1 class="text-3xl">Ваши группы:</h1>

        <div class="flex flex-col gap-3 max-h-screen">
            {#each groups as group}
                <a
                    class="w-full rounded border-slate-200 dark:border-slate-800 border bg-slate-100 dark:bg-slate-800 h-fit p-3"
                    href="/teacher/group/{group.ID}"
                >
                    <h4 class="text-xl">{group.Title}</h4>
                </a>
            {/each}
        </div>
    </main>
</div>
