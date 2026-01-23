import showToast from '@shared/modules/showToast'
import type { SocialLink } from '@shared/types'
import copy from 'copy-to-clipboard'

export default (postTitle: string): SocialLink[] => {
    const url = encodeURIComponent(window.location.href)

    return [
        {
            href: `https://www.linkedin.com/sharing/share-offsite/?url=${url}`,
            title: 'linkedin',
            tip: 'Share to LinkedIn',
        },
        {
            href: `https://www.reddit.com/submit?url=${url}&title=${postTitle}`,
            title: 'reddit',
            tip: 'Share to Reddit',
        },
        {
            href: `https://x.com/intent/tweet?text=${postTitle}&url=${url}`,
            title: 'x',
            tip: 'Share to X',
        },
        {
            href: `https://t.me/share/url?url=${url}&text=${postTitle}`,
            title: 'telegram',
            tip: 'Share to Telegram',
        },
        {
            href: null,
            title: 'clipboard',
            tip: 'Copy to a clipboard',
            action: () => {
                copy(window.location.href)

                showToast({
                    text: 'Current URL has been copied to your clipboard',
                    success: true,
                })
            },
        },
    ]
}
