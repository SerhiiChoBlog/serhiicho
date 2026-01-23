<script setup lang="ts">
import type { Tag } from '@shared/types/models'
import { ref } from 'vue'
import showEverythingTag from '@/modules/showEverythingTag'
import TagsFilter from './TagsFilter.vue'
import AppearTransition from '@shared/components/Transitions/AppearTransition.vue'

type Props = {
    currentTag: Tag | null
}

const { currentTag } = defineProps<Props>()
const tag = ref<Tag>(currentTag || showEverythingTag)

function setTag(newTag: Tag): void {
    tag.value = newTag
}
</script>

<template>
    <appear-transition>
        <div class="mb-2 flex flex-col md:flex-row items-center gap-1">
            <div class="w-full max-w-[320px]">
                <tags-filter :tag="tag" />
            </div>

            <div class="text-md leading-5">
                <h2>
                    <b :style="{ color: tag.color }">{{ tag.title }}.</b>
                    <span class="ml-2 text-font-second">{{ tag.description }}</span>
                </h2>
            </div>
        </div>
    </appear-transition>
</template>
