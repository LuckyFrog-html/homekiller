<script lang="ts">
    import "../app.pcss";
    import { Button } from "$lib/components/ui/button";
    import Sun from "lucide-svelte/icons/sun";
    import Moon from "lucide-svelte/icons/moon";
    import { ModeWatcher } from "mode-watcher";
    import type { PageData } from "./$types";

    import { toggleMode } from "mode-watcher";
    import { enhance } from "$app/forms";

    export let data: PageData;
    $: isLogined = data.isLogined;
</script>

<ModeWatcher />

<div class="flex flex-col h-dvh">
    <nav
        class="flex h-30 p-4 justify-between bg-slate-200 dark:bg-slate-900 items-center"
    >
        <a class="text-2xl" href="/">Homelander</a>
        <div class="flex items-center gap-2">
            <Button on:click={toggleMode} variant="outline" size="icon">
                <Sun class="h-[1.2rem] w-[1.2rem] rotate-0 dark:hidden" />
                <Moon class="h-[1.2rem] w-[1.2rem] hidden dark:block" />
                <span class="sr-only">Переключить тему</span>
            </Button>
            {#if isLogined}
                <form use:enhance action="/logout" method="POST">
                    <Button type="submit">Выход</Button>
                </form>
            {:else}
                <Button href="/login">Войти</Button>
            {/if}
        </div>
    </nav>
    <slot />
</div>
