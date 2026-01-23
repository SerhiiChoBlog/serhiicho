import type { User } from '@shared/types/models'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import { onMounted, ref } from 'vue'
import showToast from '@shared/modules/showToast'

export default (userSlug: string) => {
    const user = ref<User | null>(null)
    const disapproveComments = ref(false)
    const loading = ref(false)
    const comment = ref('')

    onMounted(fetchUser)

    function fetchUser(): void {
        if (!userSlug) {
            return
        }

        loading.value = true

        axios.get<User>(`/admin/ajax/users/${userSlug}`)
            .then(response => {
                user.value = response.data
            })
            .catch(handleError)
            .finally(() => loading.value = false)
    }

    function ban(): void {
        if (!user.value) {
            return
        }

        const msg = user.value.ban
            ? 'Are you sure you want to BAN this user?'
            : 'Are you sure you want to UNBAN this user?'

        if (!user.value || !confirm(msg))
            return

        if (!user.value.ban && comment.value === '') {
            showToast({ text: 'Please enter a comment', success: false })
            return
        }

        const url = user.value.ban
            ? `/admin/ajax/bans/unban`
            : `/admin/ajax/bans/ban`

        const params = {
            user_id: user.value.id,
            comment: comment.value,
            disapprove_comments: disapproveComments.value,
        }

        axios.post(url, params)
            .then(() => {
                fetchUser()
                comment.value = ''
            })
            .catch(handleError)
    }

    return {
        user,
        loading,
        comment,
        disapproveComments,
        ban,
        fetchUser,
    }
}
