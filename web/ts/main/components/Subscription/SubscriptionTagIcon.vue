<script setup lang="ts">
import type { Tag } from '@shared/types/models'
import Tip from '@shared/components/Tip.vue'

type Props = {
    tag: Tag
}

const { tag } = defineProps<Props>()
</script>

<template>
    <tip :content="tag.title">
        <a :href="`/posts?tag=${tag.name}`" rel="noreferrer">
            <img
                :src="`/storage/tags/${tag.name}.webp`"
                :alt="`The ${tag.name} logo`"
                class="w-[25px] h-[25px] drop-shadow-md transition-transform hover:scale-125 duration-300"
                width="200"
                height="200"
            />
        </a>
    </tip>
</template>
