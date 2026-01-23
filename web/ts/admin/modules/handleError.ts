import showToast from '@shared/modules/showToast'

export default (err: any) => {
    console.error(err)

    if (err && err.response && err.response.data) {
        if (err.response.data.message) {
            showToast({ text: err.response.data.message, success: false })
            return
        }

        if (err.response.data.exception) {
            showToast({ text: err.response.data.exception, success: false })
            return
        }
    }

    showToast({ text: 'Unknown error', success: false })
}
