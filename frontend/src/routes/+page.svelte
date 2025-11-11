<script lang="ts">
    import { goto } from "$app/navigation";
</script>

<main
    class="min-h-screen flex flex-col items-center justify-center bg-base-200 p-6"
>
    <h1 class="text-4xl font-bold text-primary mb-10">Media Converter</h1>

    <div class="grid grid-cols-1 sm:grid-cols-2 gap-8 w-full max-w-3xl">
        <!-- Image Converter Card -->
        <div class="card bg-base-100 shadow-xl hover:shadow-2xl transition">
            <div class="card-body items-center text-center">
                <h2 class="card-title text-secondary">Image Converter</h2>
                <p class="text-sm opacity-80">
                    Convert and save your images in different formats.
                </p>
                <div class="card-actions mt-4">
                    <button
                        class="btn btn-primary w-32"
                        onclick={() => goto("/image")}
                    >
                        Open
                    </button>
                </div>
            </div>
        </div>

        <!-- Video Converter Card -->
        <div class="card bg-base-100 shadow-xl hover:shadow-2xl transition">
            <div class="card-body items-center text-center">
                <h2 class="card-title text-accent">Video Converter</h2>
                <p class="text-sm opacity-80">
                    Convert videos to other formats easily.
                </p>
                <div class="card-actions mt-4">
                    <button
                        class="btn btn-secondary w-32"
                        onclick={() => goto("/video")}
                    >
                        Open
                    </button>
                </div>
            </div>
        </div>
    </div>
</main>
