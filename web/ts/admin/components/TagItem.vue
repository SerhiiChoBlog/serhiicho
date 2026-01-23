<script setup lang="ts">
import type { Tag } from '@shared/types/models'

defineProps<{ tag: Tag }>()

const emit = defineEmits<{
    (e: 'deleteTag', id: number): void
    (e: 'updateTagId', id: number): void
}>()
</script>

<template>
    <span
        @click="emit('updateTagId', tag.id)"
        class="text-white py-1.5 px-3 rounded-lg font-light shadow-md text-xs border-2 border-transparent"
        :class="{
            'border-page! outline outline-black dark:outline-white': tag.pivot
                ? tag.pivot.is_main
                : false,
        }"
        :style="{ backgroundColor: tag.color }"
    >
        {{ tag.name }}

        <button @click="emit('deleteTag', tag.id)" type="button">
            <i class="fas fa-times ml-1 transition-transform hover:scale-125"></i>
        </button>

        <input type="hidden" name="tags[]" :value="tag.id" />
    </span>
</template>
