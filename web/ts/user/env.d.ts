import type { User } from '@shared/types/models'

/// <reference types="vite/client" />

interface Window {
    Auth: User
}

interface ImportMetaEnv {
    readonly VITE_PUSHER_APP_KEY: string
    readonly VITE_PUSHER_APP_CLUSTER: string
}

interface ImportMeta {
    readonly env: ImportMetaEnv
}
