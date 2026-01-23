<script setup lang="ts">
import type { Post } from '@shared/types/models'
import { onMounted, reactive, ref } from 'vue'
import axios from 'axios'
import echo from '@admin/modules/echo'
import handleError from '@admin/modules/handleError'
import { Tab, TabGroup, TabList, TabPanel, TabPanels } from '@headlessui/vue'
import PageTitle from '@admin/components/PageTitle.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'

type Views = {
    unique: { post_id: number, amount: number, post: Post }[] | null
    daily: { date: string, amount:number }[] | null
}

type EchoData = {
    post_id: number | null
    country: string | null
}

const loading = ref(true)

const lastNewView = reactive<EchoData>({
    post_id: null,
    country: null,
})

const views = reactive<Views>({
    unique: null,
    daily: null,
})

onMounted(() => {
    fetchViews()

    echo<EchoData>('notification', 'NewPostView', ({ post_id, country }) => {
        fetchViews()
        lastNewView.post_id = post_id
        lastNewView.country = country

        new Audio('/storage/new-message.wav')
            .play()
            .catch(err => console.error(err))
    })
})

function fetchViews(): void {
    axios.get<Views>('/admin/ajax/post-views')
        .then(resp => {
            views.unique = resp.data.unique
            views.daily = resp.data.daily
        })
        .catch(handleError)
        .finally(() => loading.value = false)
}
</script>

<template>
    <div class="container max-w-3xl">
        <PageTitle>Post views</PageTitle>

        <Spinner v-if="loading" />

        <TabGroup v-else>
            <TabList class="flex gap-2 rounded-xl bg-gray-800 p-1 text-white">
                <Tab
                    v-for="title in ['Unique views', 'Daily views']"
                    :key="title"
                    as="template"
                    v-slot="{ selected }"
                >
                    <button
                        :class="[
                          'w-full rounded-lg py-3 text-md font-medium leading-5',
                          'ring-white ring-opacity-60 ring-offset-2 ring-offset-gray-400 focus:outline-hidden focus:ring-2',
                          selected ? 'bg-gray-600 shadow' : 'hover:bg-white/30 text-white',
                        ]"
                      >
                          {{ title }}
                    </button>
                </Tab>
            </TabList>

            <TabPanels class="mt-5">
                <TabPanel>
                    <DataTable>
                        <template #head>
                            <HeadDetail
                                v-for="title in ['Image', 'Views', 'Post title']"
                                :key="title"
                            >
                                {{ title }}
                            </HeadDetail>
                        </template>

                        <BodyRow
                            v-if="views.unique"
                            v-for="view in views.unique"
                            :key="view.post_id"
                            class="transition-colors duration-500"
                            :class="{ 'bg-yellow-200! animate-pulse': lastNewView.post_id === view.post_id }"
                        >
                            <BodyDetail>
                                <img
                                    :src="'/' + view.post.image_sm"
                                    :alt="view.post.title"
                                    class="w-20 rounded-xs shadow-xs"
                                />
                            </BodyDetail>
                            <BodyDetail>{{ view.amount }}</BodyDetail>
                            <BodyDetail class="px-3 py-1 flex items-center">
                                <span class="first-letter:uppercase leading-5">
                                    {{ view.post.title }}

                                    <template v-if="lastNewView.country && lastNewView.post_id === view.post_id">
                                        <img
                                            :src="`/storage/admin/flags/${lastNewView.country.toLowerCase()}.png`"
                                            :alt="lastNewView.country"
                                            width="20"
                                            class="inline-block ml-1 rounded-xs shadow-xs"
                                        >

                                        <span class="ml-1 text-xs font-bold">
                                            {{ lastNewView.country }}
                                        </span>
                                    </template>
                                </span>
                            </BodyDetail>
                        </BodyRow>
                    </DataTable>
                </TabPanel>

                <TabPanel>
                    <DataTable>
                        <template #head>
                            <HeadDetail
                                v-for="title in ['Date', 'Views']"
                                :key="title"
                            >
                                {{ title }}
                            </HeadDetail>
                        </template>

                        <BodyRow
                            v-if="views.daily"
                            v-for="view in views.daily"
                            :key="view.amount"
                            class="first:bg-yellow-200! first:text-black"
                        >
                            <BodyDetail>{{ view.date }}</BodyDetail>
                            <BodyDetail>{{ view.amount }}</BodyDetail>
                        </BodyRow>
                    </DataTable>
                </TabPanel>
            </TabPanels>
        </TabGroup>
    </div>
</template>
