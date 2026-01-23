<script setup lang="ts">
defineProps<{
    headline: string
    subHeadline: string
}>()
</script>

<template>
    <h2 class="text-2xl lg:text-4xl text-main">{{ headline }}</h2>
    <h3 class="text-font-second max-w-[430px]">{{ subHeadline }}</h3>
</template>
