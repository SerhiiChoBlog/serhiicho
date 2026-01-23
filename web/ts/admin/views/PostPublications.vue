<script setup lang="ts">
import type { PostPublication } from '@shared/types/models'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import PageTitle from '@admin/components/PageTitle.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'

const publications = ref<PostPublication[]>([])
const loading = ref(false)

onMounted(() => fetchVisitors())

function fetchVisitors(): void {
    axios.get<PostPublication[]>('/admin/ajax/post-publications')
        .then(resp => publications.value = resp.data)
        .catch(handleError)
        .finally(() => loading.value = false)
}
</script>

<template>
    <div class="container max-w-3xl">
        <PageTitle>Post publications</PageTitle>

        <Spinner v-if="loading" />

        <div v-else-if="publications.length === 0">
            <h2 class="text-center text-xl">
                There are no post publications yet
            </h2>
        </div>

        <DataTable v-else>
            <template #head>
                <HeadDetail
                    v-for="title in ['#', 'Notified', 'Skipped', 'Post', 'Created']"
                    :key="title"
                >
                    {{ title }}
                </HeadDetail>
            </template>

            <BodyRow v-for="publication in publications" :key="publication.id">
                <BodyDetail>{{ publication.id }}</BodyDetail>
                <BodyDetail>{{ publication.people_notified }}</BodyDetail>
                <BodyDetail>{{ publication.people_skipped }}</BodyDetail>
                <BodyDetail>{{ publication.post ? publication.post.title : '-no-' }}</BodyDetail>
                <BodyDetail class="min-w-[160px]">{{ publication.created_at }}</BodyDetail>
            </BodyRow>
        </DataTable>
    </div>
</template>
