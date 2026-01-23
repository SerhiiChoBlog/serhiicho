<script setup lang="ts">
import type { Post } from '@shared/types/models'
import type { Pagination } from '@shared/types/response'
import { onMounted, ref } from 'vue'
import handleError from '@admin/modules/handleError'
import PageTitle from '@admin/components/PageTitle.vue'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import PostItem from '@admin/components/Post/PostItem.vue'
import { debounce } from 'lodash'
import axios from 'axios'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'

const posts = ref<Post[]>([])
const loading = ref(true)
const defaultUrl = '/admin/ajax/posts'
const nextPageUrl = ref<string | null>(null)

onMounted(() => {
    loading.value = true

    fetchPosts(defaultUrl)

    window.addEventListener('scroll', () => {
        const currPageHeight = window.innerHeight + window.scrollY + 100

        if (currPageHeight >= document.body.offsetHeight) {
            loadMorePosts()
        }
    })
})

function fetchPosts(url: string, refresh = true): void {
    axios.get<Pagination<Post[]>>(url)
        .then(resp => {
            const data = resp.data

            nextPageUrl.value = data.next_page_url

            refresh
                ? posts.value = data.data
                : posts.value.push(...data.data)
        })
        .catch(handleError)
        .finally(() => loading.value = false)
}

const loadMorePosts = debounce((): void => {
    if (nextPageUrl.value) {
        fetchPosts(nextPageUrl.value, false)
    }
}, 200)
</script>

<template>
    <div class="container">
        <PageTitle>
            <div class="flex items-center justify-center gap-3">
                <span>Posts</span>

                <MainButton :route="{ name: 'posts-create' }">
                    <i class="fas fa-plus"></i>
                </MainButton>
            </div>
        </PageTitle>

        <Spinner v-if="loading" />

        <DataTable v-else-if="posts.length > 0">
            <template #head>
                <HeadDetail>ID</HeadDetail>
                <HeadDetail></HeadDetail>
                <HeadDetail>Image</HeadDetail>
                <HeadDetail class="min-w-[300px]">Title</HeadDetail>
                <HeadDetail>Intro</HeadDetail>
                <HeadDetail>Shares</HeadDetail>
                <HeadDetail>Views</HeadDetail>
                <HeadDetail>View</HeadDetail>
                <HeadDetail>Edit</HeadDetail>
            </template>

            <PostItem
                v-for="post in posts"
                :key="post.id"
                :post="post"
            />
        </DataTable>

        <h2 v-else class="text-xl text-center pt-4">There are no posts</h2>
    </div>
</template>
