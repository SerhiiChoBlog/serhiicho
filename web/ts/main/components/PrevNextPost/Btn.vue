<script setup lang="ts">
import type { Post } from '@shared/types/models'

defineProps<{ post: Post | null }>()
</script>

<template>
    <div v-if="post" class="w-full">
        <a
            :href="`/posts/${post.slug}`"
            class="px-3 py-4 block w-full h-full rounded-xl text-main-hover select-none hover:bg-page-second transition-colors"
        >
            <span class="text-2xl">
                <slot />
            </span>

            <p class="text-font-second text-sm">
                {{ post.title }}
            </p>
        </a>
    </div>

    <span v-else></span>
</template>
