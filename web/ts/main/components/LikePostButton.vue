<script setup lang="ts">
import type { Post } from '@shared/types/models'
import type { GetPostLikesResponse } from '@shared/types/response'
import dispatchEvent from '@shared/modules/dispatchEvent'
import { events } from '@shared/appConfig'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import LikeIcon from '@shared/components/Icons/LikeIcon.vue'

const { post } = defineProps<{ post: Post }>()

onMounted(() => fetchIsLiked())

const btn = ref<HTMLButtonElement | null>(null)
const isLiked = ref<boolean>(false)
const likes = ref<number>(0)
const loading = ref<boolean>(false)

function likePost(): void {
    if (!Auth) {
        dispatchEvent(events.openLoginModal)
        return
    }

    if (isLiked.value) {
        return
    }

    loading.value = true

    const params = {
        post_id: post.id,
        secret: Auth.secret,
    }

    axios
        .post<GetPostLikesResponse>(`/api/post-likes`, params)
        .then(res => {
            isLiked.value = true
            likes.value = res.data.likes
        })
        .catch(err => console.error(err))
        .finally(() => (loading.value = false))
}

function fetchIsLiked(): void {
    loading.value = true

    const params = {
        post_id: post.id,
        secret: Auth ? Auth.secret : null,
    }

    axios
        .get<GetPostLikesResponse>('/api/post-likes/is-liked', { params })
        .then(res => {
            isLiked.value = res.data.liked
            likes.value = res.data.likes
        })
        .catch(err => console.error(err))
        .finally(() => (loading.value = false))
}
</script>

<template>
    <button
        v-if="!loading"
        type="button"
        @click="likePost"
        ref="btn"
        :aria-label="`${isLiked ? 'Unlike' : 'Like'} this post`"
        class="print:hidden fixed bottom-2 right-2 md:bottom-6 md:right-6 z-30 p-2 rounded-full bg-page border-2 transition-all duration-300"
        :class="{
            'border-transparent shadow-red-400 cursor-default': isLiked,
            'border-gray-700 dark:border-gray-200 hover:shadow-[0_0_20px_red,inset_0_0_10px_red] hover:text-red-400 hover:animate-pulse':
                !isLiked,
        }"
    >
        <like-icon
            :fill="isLiked ? 'currentColor' : 'none'"
            class="w-7 h-7"
            :class="{
                'text-red-500 dark:text-red-400': isLiked,
                'text-gray-700 dark:text-gray-400': !isLiked,
            }"
        />

        <span
            v-if="likes > 0"
            class="absolute text-[.9rem] font-bold bg-page px-1.5 rounded-full transition-all"
            :class="{
                'text-red-500 dark:text-red-400 right-1/2 translate-x-1/2': isLiked,
                'text-gray-700 dark:text-gray-400 border-gray-700 dark:border-gray-400 border -top-2 right-[34px]':
                    !isLiked,
            }"
        >
            {{ likes }}
        </span>
    </button>
</template>
