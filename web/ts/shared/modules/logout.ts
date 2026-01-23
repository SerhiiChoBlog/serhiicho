import axios from 'axios'
import sweetAlert from 'sweetalert2'
import PageRefresher from '@/modules/PageRefresher'

export default async (): Promise<void> => {
    const response = await sweetAlert.fire({
        title: 'Sign out?',
        html: 'Are you sure you want to sign out from your account?',
        showDenyButton: true,
        confirmButtonText: 'Yes, sign out',
        denyButtonText: 'No',
    })

    const wantsToStay = response.isDismissed || response.isDenied

    if (wantsToStay || !window.Auth) return

    try {
        await axios.post(
            '/logout',
            { email: window.Auth.email },
            {
                headers: { 'Content-Type': 'application/json' },
            },
        )
    } catch (err) {
        console.error(err)
    }

    PageRefresher.refresh("You've been logged out")
}
