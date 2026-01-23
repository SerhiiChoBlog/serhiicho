<script setup lang="ts">
const emit = defineEmits<{
    (e: 'clicked'): void
}>()

defineProps<{
    question: string
    title: string
}>()
</script>

<template>
    <span class="flex flex-col text-center">
        <span class="text-sm text-font">{{ question }}</span>

        <button
            @click="emit('clicked')"
            class="text-sm text-main hover:underline uppercase"
        >
            {{ title }}
        </button>
    </span>
</template>
