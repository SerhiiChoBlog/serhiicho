import toast from 'toastify-js'

type ShowToastParams = {
    text: string
    success?: boolean
    url?: string
    duration?: number
}

export default (params: ShowToastParams) => {
    if (typeof params.success === 'undefined') {
        params.success = true
    }

    const config: toast.Options = {
        text: params.text,
        destination: params.url || 'javascript:',
        gravity: 'top',
        position: 'right',
        close: true,
        duration: params.duration || 4500,
        className:
            'flex shadow-md fixed p-5 justify-between z-[10000] gap-2 rounded-md right-0 md:right-10 max-w-full rounded-none sm:rounded-lg',
        style: {
            background: params.success ? '#3E8FD1' : '#c44d45',
        },
    }

    if (window.screen.availWidth < 992) {
        config.gravity = 'bottom'
    }

    toast(config).showToast()
}
