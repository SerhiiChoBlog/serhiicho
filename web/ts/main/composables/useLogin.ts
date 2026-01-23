import type { ServerError } from '@shared/types/response'
import { ref } from 'vue'
import axios from 'axios'
import handleServerError from '@shared/modules/handleServerError'
import PageRefresher from '@/modules/PageRefresher'

export default () => {
    const loading = ref<boolean>(false)
    const errors = ref<ServerError['errors'] | null>(null)
    const formData = ref({
        email: '',
        password: '',
        remember: true,
    })

    const redirectToReset = () => (window.location.href = '/password/reset')

    function login(): void {
        if (loading.value) {
            return
        }

        loading.value = true

        axios
            .post<ServerError | null>('/login', formData.value, {
                headers: { 'Content-Type': 'application/json' },
            })
            .then(resp => {
                const msg = `You've signed in! Click here to go to a profile`
                PageRefresher.refresh(msg, '/users/profile')
            })
            .catch(async e => {
                const err = await handleServerError(e)
                errors.value = err ? err.errors : null
                formData.value.password = ''
            })
            .finally(() => (loading.value = false))
    }

    return {
        loading,
        formData,
        errors,
        login,
        redirectToReset,
    }
}
