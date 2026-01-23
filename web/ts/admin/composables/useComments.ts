import type { Pagination } from '@shared/types/response'
import type { Comment } from '@shared/types/models'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import showToast from '@shared/modules/showToast'

export default () => {
    const comments = ref<Comment[]>([])
    const totalComments = ref<number | null>(null)
    const loading = ref<boolean>(false)
    const nextPageUrl = ref<string | null>(null)

    onMounted(() => {
        loading.value = true
        fetchComments('/admin/ajax/comments', true)
    })

    function deleteComment(commentId: number): void {
        if (!confirm(`Are you sure you want to delete comment #${commentId}?`)) {
            return
        }

        axios.delete(`/admin/ajax/comments/${commentId}`)
            .then(() => {
                fetchComments('/admin/ajax/comments', true)
                showToast({ text: `Comment #${commentId} has been deleted` })
            })
            .catch(handleError)
    }

    function fetchComments(url: string, refresh): void {
        axios.get<{ comments: Pagination<Comment[]>, total: number }>(url)
            .then(resp => {
                const data = resp.data

                nextPageUrl.value = data.comments.next_page_url
                totalComments.value = data.total

                refresh
                    ? comments.value = data.comments.data
                    : comments.value.push(...data.comments.data)
            })
            .catch(handleError)
            .finally(() => loading.value = false)
    }

    function loadMoreComments(): void {
        if (nextPageUrl.value) {
            fetchComments(nextPageUrl.value, false)
        }
    }

    return {
        comments,
        totalComments,
        loading,
        nextPageUrl,
        deleteComment,
        loadMoreComments,
        fetchComments,
    }
}
