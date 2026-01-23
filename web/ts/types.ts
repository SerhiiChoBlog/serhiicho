import type { User } from '@shared/types/models'

export {}

declare module '@vue/runtime-core' {
    interface ComponentCustomProperties {
        $auth: User | null
    }
}

declare global {
    const Auth: User | null
    var tinymce: any
    var Pusher: any
}
