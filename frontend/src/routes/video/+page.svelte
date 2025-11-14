<script lang="ts">
    import PathInput from "$lib/components/PathInput.svelte";
    import {
        OpenFile,
        ConvertVideo,
        SelectDestDir,
    } from "$lib/wailsjs/go/video/VideoConverter";

    type Message = {
        Error: string | null;
        Success: string | null;
    };

    let targetType: string = $state("mp4");
    let selectedPath: string = $state("");
    let destinationPath: string = $state("");
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
        try {
            selectedPath = await OpenFile();
        } catch (err: any) {
            msg.Error = String(err);
            selectedPath = "";
        }
    }

    async function handlerSelectDestination() {
        resetMessages();
        try {
            destinationPath = await SelectDestDir();
        } catch (err: any) {
            msg.Error = String(err);
            selectedPath = "";
        }
    }

    async function handleConvertFile() {
        resetMessages();

        if (!selectedPath) {
            msg.Error = "Error: Please select a file first.";
            return;
        }
        if (!destinationPath) {
            msg.Error = "Error: Please select a destination directory.";
            return;
        }

        converting = true;

        try {
            await ConvertVideo(targetType);
            msg.Success = `Conversion successful to ${targetType.toUpperCase()}!\nSaved to ${destinationPath}/converted.${targetType}`;
        } catch (err: any) {
            msg.Error = `Conversion Failed: ${String(err)}`;
        } finally {
            converting = false;
        }
    }
</script>

<main
    class="min-h-screen flex flex-col items-center justify-center bg-base-200 text-base-content p-6"
>
    <div class="card bg-base-100 shadow-2xl w-full max-w-lg">
        <div class="card-body text-center">
            <h1 class="text-3xl font-bold text-primary mb-6">
                Video Converter
            </h1>

            <PathInput
                onClick={handleSelectFile}
                bind:value={selectedPath}
                labelText={"File Input"}
                placeholderText={"No file selected"}
            />

            <PathInput
                onClick={handlerSelectDestination}
                bind:value={destinationPath}
                labelText={"Outpur Folder"}
                placeholderText={"No folder selected"}
            />

            <!-- Convert To dropdown -->
            <div class="form-control mb-6">
                <div class="flex items-center space-x-3">
                    <label for="target-format" class="label w-auto">
                        <span class="label-text font-medium">Convert To:</span>
                    </label>
                    <select
                        bind:value={targetType}
                        class="select select-primary"
                    >
                        <optgroup label="Video Formats">
                            <option value="mp4">MP4</option>
                            <option value="webm">WebM</option>
                            <option value="mov">MOV</option>
                            <option value="mkv">MKV</option>
                        </optgroup>

                        <optgroup label="Audio Formats (Extract)">
                            <option value="mp3">MP3</option>
                            <option value="opus">OPUS</option>
                            <option value="wav">WAV (Lossless)</option>
                            <option value="flac">FLAC (Lossless)</option>
                        </optgroup>

                        <optgroup label="Animated/Image">
                            <option value="gif">GIF (Animated)</option>
                        </optgroup>
                    </select>
                </div>
            </div>

            <!-- Buttons -->
            <div class="flex space-x-4 mb-4">
                <button
                    class="btn btn-primary flex-1"
                    onclick={handleConvertFile}
                    disabled={!selectedPath || !destinationPath || converting}
                >
                    {#if converting}
                        <span class="loading loading-spinner"></span>
                        Converting...
                    {:else}
                        Convert
                    {/if}
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
