<script setup lang="ts">
import type { Chapter as ChapterType } from '@shared/types'
import { computed } from 'vue'
import ChevronRightIcon from '@shared/components/Icons/ChevronRightIcon.vue'

const props = defineProps<{
    headers: NodeListOf<HTMLElement>
    chapter: ChapterType
}>()

const emit = defineEmits<{
    (e: 'clicked', title: string): void
}>()

const title = computed(() => {
    return props.chapter.title.replace(/\d+\. ?/, '')
})

const number = computed<string>(() => {
    const match = /\d+\./.exec(props.chapter.title)
    return match && match[0] ? match[0].replace('.', '') : ''
})
</script>

<template>
    <a
        :href="`#${chapter.slug}`"
        class="leading-5 py-2 border-b last:border-none border-border cursor-pointer relative flex justify-start text-font-second px-3 hover:bg-page-second select-none gap-2 items-center"
        :class="{
            'bg-page-second hover:!bg-page-second-hover': chapter.isActive,
            'ml-4': chapter.tag === 'H3',
        }"
        @click="emit('clicked', chapter.title)"
    >
        <span class="text-black dark:text-white" aria-hidden="true">
            <span v-if="number" class="mr-1">{{ number }}</span>
            <span v-else><chevron-right-icon class="w-4 h-4" /></span>
        </span>

        <span>{{ title }}</span>
    </a>
</template>
