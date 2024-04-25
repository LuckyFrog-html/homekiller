<script lang="ts">
    import { browser } from "$app/environment";
    import { goto } from "$app/navigation";

    if (browser) {
        goto("/login");
    }
</script>

<main class="flex h-full items-center justify-center">
    <p>Main</p>
</main>
