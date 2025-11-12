<script lang="ts">
    import FolderButton from "$lib/components/FolderButton.svelte";
    import {
        OpenFile,
        ConvertImage,
        SaveFile,
    } from "$lib/wailsjs/go/image/ImageConverter";

    type Message = {
        Error: string | null;
        Success: string | null;
    };

    let targetType: string = $state("png");
    let selectedPath: string = $state("");
    let convertFinished: boolean = $state(false);
    let converting: boolean = $state(false);

    let msg: Message = $state({
        Error: null,
        Success: null,
    });

    function resetMessages() {
        msg.Error = null;
        msg.Success = null;
    }

    async function handleSelectFile() {
        resetMessages();
        convertFinished = false;
        try {
            selectedPath = await OpenFile();
            if (selectedPath) {
                msg.Success =
                    "File selected successfully. Choose target format.";
            }
        } catch (err: any) {
            msg.Error = String(err);
            selectedPath = "";
        }
    }

    async function handleConvertFile() {
        resetMessages();

        if (!selectedPath) {
            msg.Error = "Error: Please select a file first.";
            convertFinished = false;
            return;
        }

        converting = true;
        convertFinished = false;

        try {
            await ConvertImage(targetType);
            convertFinished = true;
            msg.Success = `Conversion successful to ${targetType.toUpperCase()}! Ready to save.`;
        } catch (err: any) {
            msg.Error = `Conversion Failed: ${String(err)}`;
            convertFinished = false;
        } finally {
            converting = false;
        }
    }

    async function handleSaveFile() {
        resetMessages();

        if (!convertFinished) {
            msg.Error = "Error: Convert the file first before saving.";
            return;
        }

        try {
            await SaveFile();
            msg.Success = "File saved successfully!";
        } catch (err: any) {
            msg.Error = `Save Failed: ${String(err)}`;
        }
    }
</script>

<main
    class="min-h-screen flex flex-col items-center justify-center bg-base-200 text-base-content p-6"
>
    <div class="card bg-base-100 shadow-2xl w-full max-w-lg">
        <div class="card-body text-center">
            <h1 class="text-3xl font-bold text-primary mb-6">
                Image Converter
            </h1>

            <!-- File selection -->
            <div class="flex items-center space-x-3 mb-4">
                <input
                    type="text"
                    placeholder="No file selected"
                    bind:value={selectedPath}
                    readonly
                    class="input input-bordered w-full truncate"
                />
                <FolderButton onClick={handleSelectFile} />
            </div>

            <!-- Convert To dropdown -->
            <div class="form-control mb-6">
                <div class="flex items-center space-x-3">
                    <label for="target-format" class="label w-auto">
                        <span class="label-text font-medium">Convert To:</span>
                    </label>
                    <select
                        id="target-format"
                        bind:value={targetType}
                        class="select select-primary w-full"
                        onchange={() => (convertFinished = false)}
                    >
                        <option value="png">PNG</option>
                        <option value="jpeg">JPEG</option>
                        <option value="gif">GIF</option>
                        <option value="webp">WEBP</option>
                    </select>
                </div>
            </div>

            <!-- Buttons -->
            <div class="flex space-x-4 mb-4">
                <button
                    class="btn btn-primary flex-1"
                    onclick={handleConvertFile}
                    disabled={!selectedPath || converting}
                >
                    {#if converting}
                        <span class="loading loading-spinner"></span>
                        Converting...
                    {:else}
                        Convert
                    {/if}
                </button>

                <button
                    class="btn btn-accent flex-1"
                    onclick={handleSaveFile}
                    disabled={!convertFinished}
                >
                    Save File
                </button>
            </div>

            <!-- Alerts -->
            {#if msg.Error}
                <div class="alert alert-error mt-4 text-left">
                    <span>{msg.Error}</span>
                </div>
            {/if}

            {#if msg.Success}
                <div class="alert alert-success mt-4 text-left">
                    <span>{msg.Success}</span>
                </div>
            {/if}
        </div>
    </div>
</main>
