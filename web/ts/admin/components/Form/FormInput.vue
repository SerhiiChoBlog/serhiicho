<script setup lang="ts">
import LetterCounter from './LetterCounter.vue'
import { ref, watchEffect } from 'vue'

const emit = defineEmits<{
    (e: 'changed', value: string): void
}>()

const props = defineProps<{
    id: string
    type: string
    placeholder?: string
    value?: string
    length?: number
    classes?: string
    name?: string
    min?: number
    max?: number
}>()

const inputValue = ref<string>('')

watchEffect(() => {
    inputValue.value = props.value || ''
})
</script>

<template>
    <div class="inline-block w-full">
        <label :for="id" class="font-bold text-sm mb-2 text-font-second block">
            <slot />
        </label>

        <div class="inline-block relative w-full">
            <input
                :type
                class="border border-border bg-page-second py-3 px-4 rounded-md w-full"
                :class="classes"
                :placeholder="placeholder"
                v-model="inputValue"
                :name
                :id
                autocomplete="off"
                @input="emit('changed', inputValue)"
                :min
                :max
            />

            <letter-counter v-if="length" :length :inputValue />

            <slot name="input" />
        </div>
    </div>
</template>
