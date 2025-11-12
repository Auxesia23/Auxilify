<script lang="ts">
    type ClickHandler = () => Promise<void>;
    import FolderButton from "./FolderButton.svelte";
    let {
        onClick,
        value = $bindable(),
        placeholderText,
        labelText,
    }: {
        onClick: ClickHandler;
        value: string;
        placeholderText?: string;
        labelText?: string;
    } = $props();
</script>

<div class="flex flex-col mb-4">
    <label for="input-file" class="label">
        <span class="label-text text-gray-400 font-medium">{labelText}:</span>
    </label>
    <div class="flex items-center space-x-3">
        <input
            id="input-file"
            type="text"
            placeholder={placeholderText}
            bind:value
            readonly
            class="input input-bordered w-full"
        />
        <FolderButton {onClick} />
    </div>
</div>
