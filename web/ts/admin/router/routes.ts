import { RouteRecordRaw } from 'vue-router'
import Dashboard from '@admin/views/Dashboard.vue'
import Posts from '@admin/views/Posts/Posts.vue'
import CreatePost from '@admin/views/Posts/CreatePost.vue'
import EditPost from '@admin/views/Posts/EditPost.vue'
import Tags from '@admin/views/Tags/Tags.vue'
import CreateTag from '@admin/views/Tags/CreateTag.vue'
import EditTag from '@admin/views/Tags/EditTag.vue'
import PostViews from '@admin/views/PostViews.vue'
import Visitors from '@admin/views/Visitors.vue'
import SocialShares from '@admin/views/SocialShares.vue'
import Referrers from '@admin/views/Referrers.vue'
import PostPublications from '@admin/views/PostPublications.vue'
import NotFound from '@shared/components/NotFound.vue'
import Users from '@admin/views/Users.vue'
import Comments from '@admin/views/Comments.vue'
import BanUser from '@admin/views/BanUser.vue'
import Series from '@admin/views/Series/Series.vue'
import CreateSeries from '@admin/views/Series/CreateSeries.vue'
import EditSeries from '@admin/views/Series/EditSeries.vue'
import Books from '@admin/views/Books/Books.vue'
import CreateBook from '@admin/views/Books/CreateBook.vue'

const routes: RouteRecordRaw[] = [
    { path: '/admin/dashboard', name: 'dashboard', component: Dashboard },
    { path: '/admin/posts', name: 'posts', component: Posts },
    { path: '/admin/posts/create', name: 'posts-create', component: CreatePost },
    { path: '/admin/posts/:id/edit', name: 'posts-edit', component: EditPost },
    { path: '/admin/tags', name: 'tags', component: Tags },
    { path: '/admin/tags/create', name: 'tags-create', component: CreateTag },
    { path: '/admin/tags/:id/edit', name: 'tags-edit', component: EditTag },
    { path: '/admin/post-views', name: 'post-views', component: PostViews },
    { path: '/admin/visitors', name: 'visitors', component: Visitors },
    { path: '/admin/social-shares', name: 'social-shares', component: SocialShares },
    { path: '/admin/referrers', name: 'referrers', component: Referrers },
    { path: '/admin/post-publications', name: 'post-publications', component: PostPublications },
    { path: '/admin/users', name: 'users', component: Users },
    { path: '/admin/ban-user/:slug', name: 'ban-user', component: BanUser },
    { path: '/admin/comments', name: 'comments', component: Comments },
    { path: '/admin/series', name: 'series', component: Series },
    { path: '/admin/series/create', name: 'series-create', component: CreateSeries },
    { path: '/admin/series/:id/edit', name: 'series-edit', component: EditSeries },
    { path: '/admin/books', name: 'books', component: Books },
    { path: '/admin/books/create', name: 'books-create', component: CreateBook },
    { path: '/:pathMatch(.*)', name: 'not-found', component: NotFound },
]

export default routes
