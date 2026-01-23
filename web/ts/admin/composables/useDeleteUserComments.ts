import axios from 'axios';
import handleError from '@admin/modules/handleError'
import showToast from '@shared/modules/showToast'

export default (userSlug: string) => {
    async function deleteComments(): Promise<void> {
        if (!confirm('Are you sure you want to delete all the comments?'))
            return

        try {
            await axios.delete(`/admin/ajax/comments/delete-all/${userSlug}`)
            showToast({ text: 'All the comments were deleted' })
        } catch (err) {
            handleError(err)
        }
    }

    return {
        deleteComments,
    }
}
