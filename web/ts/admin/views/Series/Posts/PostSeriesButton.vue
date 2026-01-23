<script setup lang="ts">
type Props = {
    icon: string
}

const { icon } = defineProps<Props>()
</script>

<template>
    <button
        type="button"
        class="text-xl hover:scale-150 transition-transform"
    >
        <i class="fas" :class="icon" />
    </button>
</template>
