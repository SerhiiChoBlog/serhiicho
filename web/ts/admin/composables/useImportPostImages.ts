import { ref } from 'vue'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import showToast from '@shared/modules/showToast'

export default (postId: number) => {
    const importLoading = ref(false)

    async function importImages(file: File): Promise<void> {
        importLoading.value = true

        const formData = new FormData()
        formData.append('file', file)

        try {
            const uri = `/admin/ajax/post-content-images/${postId}/import`

            const resp = await axios.postForm<string>(uri, formData)

            showToast({ text: resp.data })
        } catch (err) {
            handleError(err)
        }

        importLoading.value = false
    }

    return {
        importLoading,
        importImages,
    }
}
