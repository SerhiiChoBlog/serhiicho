<script setup lang="ts">
defineProps<{
    name: string
    required: boolean
    accept?: string
}>()

const emit = defineEmits<{
    (e: 'chosen', file: File): void
}>()

function fileSelected(e: Event): void {
    const $event = e as InputEvent
    const input = $event.target as HTMLInputElement

    if (!input.files || input.files.length === 0) {
        return
    }

    emit('chosen', input.files[0])
}
</script>

<template>
    <div class="inline-block mr-2">
        <label
            class="border-4 border-gray-500 text-font cursor-pointer border-dashed px-5 py-3 min-w-[250px] uppercase block text-center hover:border-gray-700 transition-colors">

            <input
                class="hidden"
                type="file"
                :name
                :required
                :accept
                @change="fileSelected"
            />

            <span>
                <slot />
            </span>
        </label>
    </div>
</template>
