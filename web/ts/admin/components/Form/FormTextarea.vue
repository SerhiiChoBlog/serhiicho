<script setup lang="ts">
import LetterCounter from './LetterCounter.vue'
import { ref } from 'vue'

const emit = defineEmits<{
    (e: 'changed', value: string): void
}>()

const props = defineProps<{
    id: string
    placeholder: string
    value?: string
    name?: string
    length?: number
    classes?: string
}>()

const inputValue = ref<string>(props.value || '')
</script>

<template>
    <div class="inline-block w-full">
        <label :for="id" class="font-bold text-sm mb-2 text-font-second block">
            <slot />
        </label>

        <div class="inline-block relative w-full">
            <textarea
                @keyup="emit('changed', inputValue)"
                v-model="inputValue"
                class="border border-border bg-page-second py-3 px-4 rounded-md w-full"
                :class="classes"
                :placeholder
                :name
                :id
            ></textarea>

            <letter-counter v-if="length" :length :inputValue />
        </div>
    </div>
</template>
