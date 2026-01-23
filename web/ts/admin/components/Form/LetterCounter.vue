<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps<{
    length: number
    inputValue: string
}>()

const currLength = ref(0)

watch(props, () => {
    currLength.value = props.inputValue.length
})
</script>

<template>
    <span
        class="_counter absolute bg-page rounded-full border border-border p-1 top-1/2 -translate-y-1/2 -right-3 text-text-sm"
    >
        {{ length - currLength }}
    </span>
</template>
