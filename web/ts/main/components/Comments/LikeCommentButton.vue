<script setup lang="ts">
import type { Comment } from '@shared/types/models'
import axios from 'axios'
import { computed, ref } from 'vue'
import { events } from '@shared/appConfig'
import dispatchEvent from '@shared/modules/dispatchEvent'
import LikeIcon from '@shared/components/Icons/LikeIcon.vue'
import Control from '@/components/Comments/Controls/Control.vue'

const emit = defineEmits<{ (e: 'likedComment'): void }>()
const props = defineProps<{ comment: Comment }>()

const loading = ref<boolean>(false)
const isLiked = computed(() =>
    props.comment.likes ? props.comment.likes.length > 0 : false,
)

function toggleLike(): void {
    if (!Auth) {
        dispatchEvent(events.openLoginModal)
        return
    }

    if (loading.value || Auth.id === props.comment.user_id) {
        return
    }

    loading.value = true

    const params = {
        comment_id: props.comment.id,
        secret: Auth.secret,
    }

    axios
        .post(`/api/comments/like`, params)
        .then(res => emit('likedComment'))
        .catch(err => console.error(err))
        .finally(() => (loading.value = false))
}
</script>

<template>
    <control
        @click="toggleLike"
        :class="$auth && $auth.id === comment.user_id ? 'cursor-default' : 'group'"
    >
        <like-icon
            :fill="isLiked ? 'currentColor' : 'none'"
            :class="{
                'text-red-500 dark:text-red-400': isLiked,
                'text-gray-700 dark:text-gray-400': !isLiked,
            }"
            class="w-5 h-5 translate-y-[1px] transition-colors group-hover:text-red-500"
        />

        <span>
            {{ comment.likes_count }} like{{ comment.likes_count === 1 ? '' : 's' }}
        </span>
    </control>
</template>
