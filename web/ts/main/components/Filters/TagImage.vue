<script setup lang="ts">
defineProps<{ tagName: string }>()
</script>

<template>
    <img
        :src="`/storage/tags/${tagName}.webp`"
        width="200"
        height="200"
        class="w-[20px] h-[20px]"
        :alt="`Logo of ${tagName}`"
    />
</template>
