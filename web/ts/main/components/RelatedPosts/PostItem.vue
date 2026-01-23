<script setup lang="ts">
import { Post } from '@shared/types/models'

const { post } = defineProps<{ post: Post }>()
</script>

<template>
    <a
        :href="`/posts/${post.slug}`"
        class="p-2 transition-opacity hover:opacity-80 min-h-[160px]"
        :aria-label="`Go to a post ${post.title}`"
    >
        <img
            :src="`/${post.image_sm}`"
            :alt="`Photo of the post '${post.title}'`"
            class="rounded-lg shadow-lg mx-auto w-full"
            width="260"
            height="141"
        />

        <h2 class="text-lg leading-6 my-3">
            <b v-if="post.series && post.series.length > 0">
                Part {{ post.series[0]!.pivot!.part }} {{ post.series[0].title }}:
                <br />
            </b>

            {{ post.title }}
        </h2>
    </a>
</template>
