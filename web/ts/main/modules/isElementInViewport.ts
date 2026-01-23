export default (el: HTMLElement | null): boolean => {
    if (!el) {
        return false
    }

    let top = el.offsetTop
    let left = el.offsetLeft
    let width = el.offsetWidth
    let height = el.offsetHeight

    while (el && el.offsetParent) {
        el = el.offsetParent as HTMLElement
        top += el.offsetTop
        left += el.offsetLeft
    }

    return (
        top >= window.scrollY &&
        left >= window.scrollX &&
        top + height <= window.scrollY + window.innerHeight &&
        left + width <= window.scrollX + window.innerWidth
    )
}
