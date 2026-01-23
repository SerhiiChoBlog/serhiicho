<script setup lang="ts">
import type { Visitor } from '@shared/types/models'
import type { Pagination } from '@shared/types/response'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import { debounce } from 'lodash'
import handleError from '@admin/modules/handleError'
import PageTitle from '@admin/components/PageTitle.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'
import VisitorComponent from '@admin/components/Visitors/Visitor.vue'
import SortBtn from '@admin/components/Visitors/SortBtn.vue'

type RequestParams = {
    search: string|null
    order: keyof Visitor
    direction: 'asc' | 'desc'
}

const visitors = ref<Visitor[]>([])
const nextPageUrl = ref<string | null>(null)
const loading = ref(true)
const defaultUri = '/admin/ajax/visitors'
const requestParams = ref<RequestParams>({
    search: null,
    order: 'id',
    direction: 'desc',
})

onMounted(() => {
    loading.value = true

    fetchVisitors(defaultUri)

    window.addEventListener('scroll', () => {
        const currPageHeight = window.innerHeight + window.scrollY + 100

        if (currPageHeight >= document.body.offsetHeight) {
            loadMoreVisitors()
        }
    })
})

function fetchVisitors(url: string, refresh = true): void {
    axios.get<Pagination<Visitor[]>>(url, {
        params: requestParams.value,
    })
        .then(resp => {
            const data = resp.data

            nextPageUrl.value = data.next_page_url

            refresh
                ? visitors.value = data.data
                : visitors.value.push(...data.data)
        })
        .catch(handleError)
        .finally(() => loading.value = false)
}

function searchVisitors(): void {
    fetchVisitors(defaultUri, true)
}

const loadMoreVisitors = debounce((): void => {
    if (nextPageUrl.value) {
        fetchVisitors(nextPageUrl.value, false)
    }
}, 200)

function sortVisitors(orderBy: keyof Visitor): void {
    requestParams.value.order = orderBy
    fetchVisitors(defaultUri)
}

function changeDescendingOrder(orderBy: keyof Visitor): void {
    requestParams.value.direction = requestParams.value.direction === 'asc'
        ? 'desc'
        : 'asc'

    sortVisitors(orderBy)
}
</script>

<template>
    <div class="container max-w-3xl">
        <PageTitle>Visitors</PageTitle>

        <div class="text-center mb-4">
            <input
                type="search"
                class="px-5 py-2 border border-border rounded-md w-full mx-auto max-w-[500px]"
                placeholder="Enter you search query..."
                v-model="requestParams.search"
                @keyup="searchVisitors"
            />
        </div>

        <Spinner v-if="loading" />

        <DataTable v-else>
            <template #head>
                <HeadDetail class="w-20">
                    <SortBtn
                        @order="sortVisitors('id')"
                        @direction="changeDescendingOrder('id')"
                        :class="{ 'text-yellow-200': requestParams.order === 'id' }"
                    >
                        ID
                    </SortBtn>
                </HeadDetail>

                <HeadDetail>
                    <SortBtn
                        @order="sortVisitors('is_robot')"
                        @direction="changeDescendingOrder('is_robot')"
                        :class="{ 'text-yellow-200': requestParams.order === 'is_robot' }"
                    >
                        🤖
                    </SortBtn>
                </HeadDetail>
                <HeadDetail class="w-36">
                    <SortBtn
                        @order="sortVisitors('ip')"
                        @direction="changeDescendingOrder('ip')"
                        :class="{ 'text-yellow-200': requestParams.order === 'ip' }"
                    >
                        IP
                    </SortBtn>
                </HeadDetail>
                <HeadDetail class="w-14">
                    <SortBtn
                        @order="sortVisitors('code')"
                        @direction="changeDescendingOrder('code')"
                        :class="{ 'text-yellow-200': requestParams.order === 'code' }"
                    >
                        Flag
                    </SortBtn>
                </HeadDetail>
                <HeadDetail>
                    <SortBtn
                        @order="sortVisitors('country')"
                        @direction="changeDescendingOrder('country')"
                        :class="{ 'text-yellow-200': requestParams.order === 'country' }"
                    >
                        Country
                    </SortBtn>
                </HeadDetail>
                <HeadDetail>
                    <SortBtn
                        @order="sortVisitors('city')"
                        @direction="changeDescendingOrder('city')"
                        :class="{ 'text-yellow-200': requestParams.order === 'city' }"
                    >
                        City
                    </SortBtn>
                </HeadDetail>
                <HeadDetail class="w-16">
                    <SortBtn
                        @order="sortVisitors('post_views_count')"
                        @direction="changeDescendingOrder('post_views_count')"
                        :class="{ 'text-yellow-200': requestParams.order === 'post_views_count' }"
                    >
                        Views
                    </SortBtn>
                </HeadDetail>
                <HeadDetail>
                    <SortBtn
                        @order="sortVisitors('is_blocked')"
                        @direction="changeDescendingOrder('is_blocked')"
                        :class="{ 'text-yellow-200': requestParams.order === 'is_blocked' }"
                    >
                        Ban
                    </SortBtn>
                </HeadDetail>
            </template>

            <VisitorComponent
                v-for="visitor in visitors"
                :key="visitor.id"
                :visitor="visitor"
                @refreshVisitors="fetchVisitors(defaultUri, true)"
            />
        </DataTable>
    </div>
</template>
