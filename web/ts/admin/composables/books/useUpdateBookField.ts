import type { UpdateBookField } from '@admin/types'
import { ref } from 'vue'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import showToast from '@shared/modules/showToast'

export default () => {
    const loading = ref<boolean>(false)

    function updateBookField(bookSlug: string, field: UpdateBookField, value: string) {
        if (loading.value) {
            return
        }

        loading.value = true

        const params = {
            field_name: field,
            field_value: value,
        }

        axios.put(`/admin/ajax/books/${bookSlug}/update-field`, params)
            .then(() => showToast({ text: 'Book field updated successfully' }))
            .catch(handleError)
            .finally(() => loading.value = false)
    }

    return {
        updateBookField,
    }
}
