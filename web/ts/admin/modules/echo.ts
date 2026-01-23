import Echo from 'laravel-echo'
import { Channel } from 'laravel-echo/dist/channel'

export default <T>(channel: string, event: string, callback: (e: T) => void): Channel => {
    return new Echo({
        broadcaster: 'pusher',
        key: import.meta.env.VITE_PUSHER_APP_KEY,
        cluster: import.meta.env.VITE_PUSHER_APP_CLUSTER,
        forceTLS: true,
        encrypted: true,
    })
        .channel(channel)
        .listen(event, callback)
}
