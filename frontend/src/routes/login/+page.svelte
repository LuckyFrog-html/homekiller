<script lang="ts">
    import type { PageData } from "./$types";
    import LoginForm from "./login-form.svelte";
    import * as Card from "$lib/components/ui/card";

    export let data: PageData;
</script>

<main class="flex h-full items-center justify-center">
    <Card.Root class="w-[30rem]">
        <Card.Header>
            <Card.Title>Вход в аккаунт ученика</Card.Title>
        </Card.Header>
        <Card.Content>
            <LoginForm data={data.form} />
        </Card.Content>
    </Card.Root>
</main>
