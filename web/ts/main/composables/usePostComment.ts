import handleServerError from '@shared/modules/handleServerError'
import axios from 'axios'
import { onMounted, ref } from 'vue'
import Cache from '@shared/modules/Cache'
import { events, storage } from '@shared/appConfig'
import dispatchEvent from '@shared/modules/dispatchEvent'
import listenEvent from '@shared/modules/listenEvent'
import showToast from '@shared/modules/showToast'

export default (postId: number, commentId?: number) => {
    const isLoading = ref<boolean>(false)
    const comment = ref<string>('')
    const canCommentIn = ref<number>(0)

    onMounted(() => restoreTimer())

    listenEvent<string>(events.commentIsPosted, startTimer)

    async function postComment(): Promise<void> {
        if (isLoading.value) {
            return
        }

        if (!Auth) {
            dispatchEvent(events.openLoginModal)
            return
        }

        isLoading.value = true

        const params = {
            comment: comment.value,
            comment_id: commentId || null,
            secret: Auth.secret,
            post_id: postId,
        }

        try {
            const response = await axios.post<{ date: string; message: string }>(
                '/api/comments',
                params,
            )
            handleResponse(response.data)
        } catch (err) {
            handleServerError(err)
        }

        isLoading.value = false
    }

    function handleResponse(data: { date: string; message: string }): void {
        showToast({ text: data.message })
        comment.value = ''
        dispatchEvent(events.commentIsPosted, data.date)
    }

    function startTimer(dateStr: string): void {
        const date = new Date(dateStr)
        const canComment = Math.floor(date.getTime() / 1000)
        const now = Math.floor(Date.now() / 1000)

        canCommentIn.value = canComment - now

        if (canCommentIn.value <= 0) {
            Cache.delete(storage.allowCommentAt)
            return
        }

        Cache.set<string>(storage.allowCommentAt, dateStr)

        const interval = window.setInterval(() => {
            canCommentIn.value--

            if (canCommentIn.value <= 0) {
                window.clearInterval(interval)
            }
        }, 1000)
    }

    function restoreTimer(): void {
        const dateStr = Cache.get<string>(storage.allowCommentAt)

        if (!dateStr) {
            return
        }

        startTimer(dateStr)
    }

    return {
        comment,
        postComment,
        canCommentIn,
        isLoading,
    }
}
