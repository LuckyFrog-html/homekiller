<script lang="ts">
    import type { PageData } from "./$types";
    import Form from "./form.svelte";
    export let data: PageData;
</script>

<div class="flex h-full w-full">
    <div class="flex w-full flex-col gap-2 p-6">
        <div>
            <a href="/student/tasks" class="text-3xl">&lt;-</a>
            <span class="text-3xl">{data.task.subject}</span>
        </div>
        <main class="flex lg:flex-row flex-col gap-3">
            <div class="w-full">
                <p>{data.task.description}</p>
            </div>
            <div class="bg-slate-300 max-w-[800px] w-full h-[400px] text-black">
                IMAGE STUB
            </div>
        </main>
        <div class="max-w-80">
            <Form data={data.form} />
        </div>
    </div>
</div>
