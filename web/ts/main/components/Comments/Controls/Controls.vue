<script setup lang="ts">
import type { Comment } from '@shared/types/models'
import LikeCommentButton from '@/components/Comments/LikeCommentButton.vue'
import Control from '@/components/Comments/Controls/Control.vue'

const emit = defineEmits<{
    (e: 'reply'): void
    (e: 'likedComment'): void
}>()

defineProps<{
    comment: Comment
    reply: boolean
}>()
</script>

<template>
    <div class="flex gap-2 items-center">
        <like-comment-button :comment @likedComment="emit('likedComment')" />

        <span v-if="$auth && comment.user_id !== $auth.id">·</span>

        <control v-if="$auth && comment.user_id !== $auth.id" @click="emit('reply')">
            <span v-if="reply">Cancel</span>
            <span v-else>Reply</span>
        </control>

        <span>·</span>

        <small class="text-gray-400">
            {{ comment.pretty_created_at }}
        </small>
    </div>
</template>
