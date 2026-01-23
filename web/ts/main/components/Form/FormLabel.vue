<script setup lang="ts">
type Props = {
    id: string
    isRequired?: boolean
}

const { id, isRequired = false } = defineProps<Props>()
</script>

<template>
    <label
        class="flex items-center gap-1.5 leading-4 text-sm text-font-second font-bold mb-1"
        :for="id"
    >
        <i v-if="isRequired" class="text-red-500">*</i>
        <slot />
    </label>
</template>
