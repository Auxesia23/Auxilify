<script lang="ts">
    import "../app.css";
    import favicon from "$lib/assets/favicon.svg";

    let { children } = $props();

    function goBack() {
        history.back();
    }
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
</svelte:head>

<!-- Tombol back -->
<nav class="w-full flex items-center p-4 bg-base-200 fixed">
    <button class="btn btn-ghost btn-sm" onclick={goBack}> ← Back </button>
</nav>

<!-- Konten halaman -->
<main class="flex-1">
    {@render children()}
</main>
