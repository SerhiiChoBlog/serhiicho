import { ref } from 'vue'

export default () => {
    const isScrolledDown = ref<boolean>(window.scrollY > 10)

    document.addEventListener('scroll', () => {
        isScrolledDown.value = window.scrollY > 10
    })

    return {
        isScrolledDown,
    }
}
