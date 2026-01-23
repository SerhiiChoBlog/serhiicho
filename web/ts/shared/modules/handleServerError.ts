import type { ServerError } from '@shared/types/response'
import showToast from '@shared/modules/showToast'

export default async (err: any): Promise<ServerError | null> => {
    if (!err.response || !err.response.data) {
        showToast({ text: 'Unknown server error', success: false })
        return null
    }

    const error = err.response.data as ServerError

    console.error(error.message)

    showToast({
        text: error.message || 'Server error',
        success: false,
    })

    return error
}
