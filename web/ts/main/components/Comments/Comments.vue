<script setup lang="ts">
import { computed } from 'vue'
import useComments from '@/composables/useComments'
import Spinner from '@shared/components/Icons/Spinner.vue'
import Comment from '@/components/Comments/Comment.vue'
import CommentForm from '@/components/Comments/CommentForm.vue'
import CommentsWrapper from '@/components/Comments/CommentsWrapper.vue'
import CommentAvatar from '@/components/Comments/CommentAvatar.vue'

const { postId } = defineProps<{ postId: number }>()
const { comments, commentsLoading, fetchComments } = useComments(postId)

const commentsCount = computed(() => {
    return comments.value.reduce((a, c) => a + c.comments!.length + 1, 0)
})
</script>

<template>
    <div class="mb-5 lg:mb-12 mx-auto print:hidden">
        <div class="flex gap-5">
            <comment-avatar :user="$auth" />

            <div class="w-full">
                <h2
                    class="text-lg mb-1 lg:text-2xl font-bold text-gray-900 dark:text-white"
                >
                    Discussion ({{ commentsCount }})
                </h2>

                <comment-form @posted="fetchComments(false)" :post-id="postId" />
            </div>
        </div>

        <div v-if="commentsLoading" class="flex gap-6 items-center justify-center">
            <spinner class="mx-0" />
            <h2 class="text-2xl">Loading comments...</h2>
        </div>

        <comments-wrapper v-else-if="comments.length > 0">
            <comment
                v-for="comment in comments"
                :key="comment.id"
                :comment="comment"
                :post-id="postId"
                @posted="fetchComments(false)"
                @likedComment="fetchComments(false)"
            />
        </comments-wrapper>

        <h2 v-else class="text-xl text-center text-font-second mb-8 mt-5">
            🙏 Be the first to leave a comment
        </h2>
    </div>
</template>
