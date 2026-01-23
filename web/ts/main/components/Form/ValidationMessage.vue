<script setup lang="ts">
defineProps<{ message?: string | null }>()
</script>

<template>
    <span
        v-if="message"
        class="text-sm text-red-600 leading-4 block mt-1 ml-1"
        role="alert"
    >
        {{ message }}
    </span>
</template>
