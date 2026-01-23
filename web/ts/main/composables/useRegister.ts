import type { ServerError } from '@shared/types/response'
import { ref } from 'vue'
import axios from 'axios'
import handleServerError from '@shared/modules/handleServerError'
import PageRefresher from '@/modules/PageRefresher'

export default () => {
    const loading = ref<boolean>(false)
    const errors = ref<ServerError['errors'] | null>(null)
    const formData = ref({
        name: '',
        email: '',
        password: '',
    })

    function register(): void {
        if (loading.value) {
            return
        }

        loading.value = true

        axios
            .post('/register', formData.value, {
                headers: { 'Content-Type': 'application/json' },
            })
            .then(resp => {
                const msg =
                    'Thank you for signing up! Verify your email address to get all the benefits.'
                PageRefresher.refresh(msg)
            })
            .catch(async e => {
                const err = await handleServerError(e)
                errors.value = err ? err.errors : null
            })
            .finally(() => (loading.value = false))
    }

    return {
        loading,
        formData,
        errors,
        register,
    }
}
