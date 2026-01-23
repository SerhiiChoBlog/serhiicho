import type { Comment as CommentType } from '@shared/types/models'
import axios from 'axios'
import { onMounted, ref } from 'vue'

export default (postId: number) => {
    const comments = ref<CommentType[]>([])
    const commentsLoading = ref<boolean>(false)

    onMounted(() => fetchComments(true))

    function fetchComments(withLoading: boolean): void {
        if (withLoading) {
            commentsLoading.value = true
        }

        axios
            .get<CommentType[]>(`/api/comments`, {
                params: {
                    post_id: postId,
                    secret: Auth ? Auth.secret : null,
                },
            })
            .then(resp => (comments.value = resp.data))
            .catch(err => console.error(err))
            .finally(() => (commentsLoading.value = false))
    }

    return {
        comments,
        commentsLoading: commentsLoading,
        fetchComments,
    }
}
