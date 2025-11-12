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
    <button class="btn btn-ghost btn-sm" onclick={goBack} title="Back">
        <svg
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
            stroke-width="1.5"
            stroke="currentColor"
            class="size-6"
        >
            <path
                stroke-linecap="round"
                stroke-linejoin="round"
                d="M9 15 3 9m0 0 6-6M3 9h12a6 6 0 0 1 0 12h-3"
            />
        </svg>
    </button>
</nav>

<!-- Konten halaman -->
<main class="flex-1">
    {@render children()}
</main>
