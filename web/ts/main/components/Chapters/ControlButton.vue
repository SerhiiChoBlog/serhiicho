<script setup lang="ts">
import Tip from '@shared/components/Tip.vue'

interface Props {
    tipContent: string
    disabled: boolean
    label: string
}

const { tipContent, disabled, label } = defineProps<Props>()
</script>

<template>
    <tip :content="tipContent" :aria-label="tipContent">
        <button
            type="button"
            class="flex items-center transition-transform"
            :aria-label="label"
            :class="{
                'opacity-40 cursor-not-allowed': disabled,
                'hover:scale-125': !disabled,
            }"
        >
            <slot />
        </button>
    </tip>
</template>
