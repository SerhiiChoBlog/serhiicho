<script setup lang="ts">
const emit = defineEmits<{
    (e: 'order'): void
    (e: 'direction'): void
}>()
</script>

<template>
    <div class="flex gap-2 items-center">
        <button
            @click="emit('direction')"
            type="button"
            aria-label="Set order direction"
            class="hover:text-yellow-200"
        >
            <i class="fas fa-sort" />
        </button>

        <button
            @click="emit('order')"
            type="button"
            class="hover:text-yellow-200"
            aria-label="Sort visitors"
        >
            <slot />
        </button>
    </div>
</template>
