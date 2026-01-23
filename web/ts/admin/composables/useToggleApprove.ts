import type { Comment } from '@shared/types/models'
import axios from 'axios'
import handleError from '@admin/modules/handleError'

export default () => {
    function toggleApproved(c: Comment): void {
        const url = c.is_approved
            ? `/admin/ajax/comments/disapprove/${c.id}`
            : `/admin/ajax/comments/approve/${c.id}`

        axios.put(url).catch(handleError)
    }

    return {
        toggleApproved,
    }
}
