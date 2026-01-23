<script setup lang="ts">
import { computed, onMounted } from 'vue'
import usePostComment from '@/composables/usePostComment'
import CommentButton from '@/components/Comments/CommentButton.vue'
import useCachedComment from '@/composables/useCachedComment'
import FirebaseSignInButton from '@/components/FirebaseSignInButton.vue'

const emit = defineEmits<{ (e: 'posted'): void }>()

const { postId, commentId } = defineProps<{
    postId: number
    commentId?: number
}>()

const { comment, postComment, isLoading, canCommentIn } = usePostComment(
    postId,
    commentId,
)

const { getCommentFromCache, saveCommentToCache, deleteCommentFromCache } =
    useCachedComment(postId, commentId)

onMounted(() => {
    comment.value = getCommentFromCache()
})

const placeholder = computed(() => {
    return canCommentIn.value > 0
        ? `To prevent spam, next comment you can write in ${canCommentIn.value} seconds`
        : 'Write a comment ...'
})

async function submitForm(): Promise<void> {
    await postComment()
    emit('posted')
    deleteCommentFromCache()
}
</script>

<template>
    <form @submit.prevent="submitForm" class="w-full mb-5 lg:mb-8 relative">
        <div>
            <div class="mb-2">
                <textarea
                    v-if="$auth"
                    v-model="comment"
                    :placeholder="placeholder"
                    class="p-3 bg-page focus:outline-hidden dark:bg-page-second rounded-lg border border-border w-full min-h-[90px]"
                    @blur="saveCommentToCache(comment)"
                    :disabled="canCommentIn > 0"
                    minlength="4"
                    maxlength="1000"
                ></textarea>

                <p v-else class="text-xl py-3">Please, sign in to comment</p>

                <span
                    v-if="$auth"
                    class="text-xs lg:text-lg absolute right-1 bottom-4 lg:bottom-6 opacity-60"
                >
                    Letters: {{ comment.length }}/1000
                </span>
            </div>

            <div class="flex gap-4 items-center">
                <comment-button
                    type="submit"
                    :disabled="isLoading || canCommentIn > 0"
                    :class="{ 'opacity-50': isLoading || canCommentIn > 0 }"
                >
                    <span v-if="isLoading">Processing...</span>
                    <span v-else-if="!$auth">Sign in with email</span>
                    <span v-else>Comment</span>
                </comment-button>

                <div v-if="!$auth" class="inline-flex gap-4 items-center">
                    <firebase-sign-in-button service-name="google" />
                    <firebase-sign-in-button service-name="github" />
                </div>
            </div>
        </div>
    </form>
</template>
