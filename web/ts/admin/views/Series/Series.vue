<script setup lang="ts">
import type { Series as SeriesType } from '@shared/types/models'
import type { Pagination } from '@shared/types/response'
import { onMounted, ref } from 'vue'
import { debounce } from 'lodash'
import axios from 'axios'
import moment from 'moment'
import handleError from '@admin/modules/handleError'
import IconButton from '@admin/components/Buttons/IconButton.vue'
import PageTitle from '@admin/components/PageTitle.vue'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'
import Tip from '@shared/components/Tip.vue'

const series = ref<SeriesType[]>([])
const totalSeries = ref<number | null>(null)
const loading = ref(false)
const nextPageUrl = ref<string | null>(null)

onMounted(() => {
    loading.value = true
    fetchSeries('/admin/ajax/series', true)

    window.addEventListener('scroll', () => {
        const currPageHeight = window.innerHeight + window.scrollY + 100

        if (currPageHeight >= document.body.offsetHeight) {
            loadMoreSeries()
        }
    })
})

function fetchSeries(url: string, refresh = true): void {
    axios.get<{ series: Pagination<SeriesType[]>, total: number }>(url)
        .then(resp => {
            const data = resp.data

            nextPageUrl.value = data.series.next_page_url
            totalSeries.value = data.total

            refresh
                ? series.value = data.series.data
                : series.value.push(...data.series.data)
        })
        .catch(handleError)
        .finally(() => loading.value = false)
}

const loadMoreSeries = debounce((): void => {
    if (nextPageUrl.value) {
        fetchSeries(nextPageUrl.value, false)
    }
}, 200)
</script>

<template>
    <div class="container max-w-7xl">
        <PageTitle>
            <div class="flex items-center justify-center gap-3">
                <span>Series <b v-if="totalSeries">{{ totalSeries }}</b></span>

                <MainButton :route="{ name: 'series-create' }">
                    <i class="fas fa-plus"></i>
                </MainButton>
            </div>
        </PageTitle>

        <Spinner v-if="loading" />

        <h2 v-else-if="series.length === 0" class="text-center text-xl">
            There are no series yet
        </h2>

        <DataTable v-else>
            <template #head>
                <HeadDetail class="w-12">ID</HeadDetail>
                <HeadDetail class="w-56">Image</HeadDetail>
                <HeadDetail>Title</HeadDetail>
                <HeadDetail>Intro</HeadDetail>
                <HeadDetail>Posts</HeadDetail>
                <HeadDetail>Created at</HeadDetail>
                <HeadDetail>Edit</HeadDetail>
            </template>

            <BodyRow v-for="s in series" :key="s.id">
                <BodyDetail>{{ s.id }}</BodyDetail>
                <BodyDetail :style="{ 'background': `linear-gradient(90deg, ${s.color_from} 0%, ${s.color_to} 100%)` }">
                    <router-link :to="{ name: 'series-edit', params: { id: s.id.toString() } }">
                        <img
                            :src="`/${s.image_sm}`"
                            width="200"
                            class="rounded-md shadow-md"
                            :alt="s.title"
                        />
                    </router-link>
                </BodyDetail>
                <BodyDetail>{{ s.title }}</BodyDetail>
                <BodyDetail>{{ s.intro }}</BodyDetail>
                <BodyDetail>
                    <span v-if="s.posts.length === 0">No posts</span>

                    <div v-else class="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-1 max-w-[200px]">
                        <div v-for="post in s.posts" :key="post.id">
                            <a
                                :href="`/admin/posts/${post.id}/edit`"
                                target="_blank"
                                rel="noreferrer"
                                class="inline-block"
                            >
                                <Tip :content="post.title">
                                    <img
                                        :src="`/${post.image_sm}`"
                                        class="rounded-md shadow-md w-full"
                                        :alt="post.title"
                                    />
                                </Tip>
                            </a>
                        </div>
                    </div>
                </BodyDetail>
                <BodyDetail>
                    {{ moment(s.created_at).format('D MMM YYYY H:mm') }}
                </BodyDetail>
                <BodyDetail class="text-center">
                    <icon-button
                        icon="fa-pen"
                        :route="{ name: 'series-edit', params: { id: s.id.toString() } }"
                        title="Edit series"
                    />
                </BodyDetail>
            </BodyRow>
        </DataTable>
    </div>
</template>
