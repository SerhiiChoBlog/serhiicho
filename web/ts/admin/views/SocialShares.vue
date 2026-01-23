<script setup lang="ts">
import type { Post } from '@shared/types/models'
import { onMounted, ref } from 'vue'
import handleError from '@admin/modules/handleError'
import axios from 'axios'
import PageTitle from '@admin/components/PageTitle.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'

const posts = ref<Post[]>([])
const loading = ref(true)

onMounted(() => fetchPosts())

function fetchPosts(): void {
    axios.get<Post[]>('/admin/ajax/social-shares')
        .then(resp => posts.value = resp.data)
        .catch(handleError)
        .finally(() => loading.value = false)
}
</script>

<template>
    <div class="container">
        <PageTitle>Social shares</PageTitle>

        <Spinner v-if="loading" />

        <h2 v-else-if="posts.length === 0" class="text-2xl text-center">There are no posts</h2>

        <DataTable v-else>
            <template #head>
                <HeadDetail
                    v-for="title in ['Image', 'Article', 'Social']"
                    :key="title"
                >
                    {{ title }}
                </HeadDetail>
            </template>

            <BodyRow v-for="post in posts" :key="post.id">
                <BodyDetail>
                    <img
                        :src="`/${post.image_sm}`"
                        width="120"
                        class="rounded-md shadow-md"
                        :alt="post.title"
                    />
                </BodyDetail>

                <BodyDetail>
                    <a :href="`/posts/${post.slug}`" target="_blank" class="text-lg" rel="noreferrer">
                        {{ post.title }}
                    </a>
                </BodyDetail>

                <BodyDetail>
                    <div class="flex gap-3">
                        <div
                            class="flex flex-col lg:flex-row gap-x-1 items-center"
                            v-if="post.social"
                            v-for="(items, social) in post.social"
                            :key="social"
                        >
                            <img
                                :src="`/storage/icons/${social.toString()}.webp`"
                                :alt="`social`"
                                width="20"
                            />

                            <span class="text-md font-bold text-font-second">
                                {{ items.length }}
                            </span>
                        </div>
                    </div>
                </BodyDetail>
            </BodyRow>
        </DataTable>
    </div>
</template>
