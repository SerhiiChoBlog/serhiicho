import { ref } from 'vue'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import downloadFile from '@shared/modules/downloadFile'

export default (postId: number) => {
    const exportLoading = ref(false)

    async function downloadImages(): Promise<void> {
        exportLoading.value = true

        try {
            const uri = `/admin/ajax/post-content-images/${postId}/export`
            const resp = await axios.post<string>(uri)
            downloadFile(resp.data, 'post-files.zip')
        } catch (err) {
            handleError(err)
        }

        exportLoading.value = false
    }

    return {
        exportLoading,
        downloadImages,
    }
}
