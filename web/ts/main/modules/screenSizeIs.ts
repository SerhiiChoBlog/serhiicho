type Size = 'sm' | 'md' | 'lg' | 'xl' | '2xl'

export default (screen: Size | Size[]): boolean => {
    if (Array.isArray(screen)) {
        return screen.some(size => screenSize(size))
    }

    return screenSize(screen)
}

function screenSize(screen: Size): boolean {
    switch (screen) {
        case 'sm':
            return window.innerWidth <= 640
        case 'md':
            return window.innerWidth <= 768
        case 'lg':
            return window.innerWidth <= 1024
        case 'xl':
            return window.innerWidth <= 1280
        default:
            return window.innerWidth <= 1536
    }
}
