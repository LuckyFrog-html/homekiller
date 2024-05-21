<script lang="ts">
    import { Button } from "$lib/components/ui/button";

    export let status: number;
    export let message: string;
</script>

<div class="flex h-full items-center justify-center">
    {#if status === undefined && message === undefined}
        <div class="flex flex-col items-center gap-10">
            <h1 class="text-3xl font-bold">Страница не найдена</h1>
            <Button href="/">Вернуться на главную</Button>
        </div>
    {:else}
        <p class="text-4xl font-bold">{status}</p>
        <p class="text-4xl font-bold">{message}</p>
    {/if}
</div>
