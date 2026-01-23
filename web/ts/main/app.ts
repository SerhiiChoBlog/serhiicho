import { onMounted } from 'vue'
import { createApp } from 'vue/dist/vue.esm-bundler'
import smoothLoader from 'smooth-loader'
import insertTagsForHighlightJs from '@/modules/insertTagsForHighlightJs'
import globalValues from '@/plugins/globalValues'
import PageRefresher from '@/modules/PageRefresher'
import ContactModal from '@/components/ContactModal.vue'
import Navbar from '@/components/Navbar/Navbar.vue'
import SharePostButtons from '@/components/SharePostButtons.vue'
import Chapters from '@/components/Chapters/Chapters.vue'
import Tip from '@shared/components/Tip.vue'
import Search from '@/components/Search/Search.vue'
import Filters from '@/components/Filters/Filters.vue'
import Subscription from '@/components/Subscription/Subscription.vue'
import ShowPasswordButton from '@/components/ShowPasswordButton.vue'
import Comments from '@/components/Comments/Comments.vue'
import LazyLoading from '@shared/components/LazyLoading.vue'
import FirebaseSignInButton from '@/components/FirebaseSignInButton.vue'
import RelatedPosts from '@/components/RelatedPosts/RelatedPosts.vue'
import LikePostButton from '@/components/LikePostButton.vue'
import LoginForm from '@shared/components/SignIn/LoginForm.vue'
import RegisterForm from '@shared/components/SignIn/RegisterForm.vue'
import ZoomedImage from '@/components/ZoomedImage/ZoomedImage.vue'
import SignInModal from '@shared/components/SignIn/SignInModal.vue'
import Testimonials from '@/components/Testimonials/Testimonials.vue'
import PrevNextPost from '@/components/PrevNextPost/PrevNextPost.vue'

const app = createApp({
    components: {
        ContactModal,
        Navbar,
        SharePostButtons,
        Chapters,
        Search,
        Tip,
        Filters,
        Subscription,
        ShowPasswordButton,
        Comments,
        LazyLoading,
        FirebaseSignInButton,
        RelatedPosts,
        LikePostButton,
        SignInModal,
        LoginForm,
        RegisterForm,
        ZoomedImage,
        Testimonials,
        PrevNextPost,
    },

    setup() {
        onMounted(() => {
            insertTagsForHighlightJs()
            PageRefresher.scrollToTheLastPosition()
        })
    },
})

app.use(globalValues).mount('#main')

smoothLoader('[data-src]')
