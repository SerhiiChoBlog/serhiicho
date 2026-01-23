import { createApp } from "vue/dist/vue.esm-bundler"
import Avatar from '@user/components/Profile/Avatar.vue'
import Tip from '@shared/components/Tip.vue'
import Feed from '@user/components/Profile/Feed.vue'
import SignInModal from '@shared/components/SignIn/SignInModal.vue'

const app = createApp({
    components: {
        Avatar,
        Tip,
        Feed,
        SignInModal,
    },
})

app.config.globalProperties.Auth = Auth

app.mount('#user-main')
