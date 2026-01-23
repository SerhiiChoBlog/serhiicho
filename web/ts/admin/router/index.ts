import { createRouter, createWebHistory } from 'vue-router'
import routes from './routes'

const router = createRouter({
    linkActiveClass: 'bg-main',
    history: createWebHistory(),
    routes,
})

router.afterEach((to, from) => {
    // @ts-ignore
    window.scrollTo({ top: 0, behavior: 'instant' })
})

export default router
