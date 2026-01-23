<script setup lang="ts">
import type { User } from '@shared/types/models'
import type { Pagination } from '@shared/types/response'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import moment from 'moment'
import handleError from '@admin/modules/handleError'
import PageTitle from '@admin/components/PageTitle.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'
import LoadMoreButton from '@shared/components/LoadMoreButton.vue'

const users = ref<User[]>([])
const totalUsers = ref<number | null>(null)
const loading = ref(false)
const nextPageUrl = ref<string | null>(null)

onMounted(() => {
    loading.value = true
    fetchUsers('/admin/ajax/users', true)
})

function fetchUsers(url: string, refresh = true): void {
    axios
        .get<{ users: Pagination<User[]>; total: number }>(url)
        .then(resp => {
            const data = resp.data

            nextPageUrl.value = data.users.next_page_url
            totalUsers.value = data.total

            refresh
                ? (users.value = data.users.data)
                : users.value.push(...data.users.data)
        })
        .catch(handleError)
        .finally(() => (loading.value = false))
}

function loadMoreUsers(): void {
    if (nextPageUrl.value) {
        fetchUsers(nextPageUrl.value, false)
    }
}
</script>

<template>
    <div class="container max-w-7xl">
        <PageTitle>
            Users <b v-if="totalUsers">{{ totalUsers }}</b>
        </PageTitle>

        <Spinner v-if="loading" />

        <DataTable v-else>
            <template #head>
                <HeadDetail class="w-12">ID</HeadDetail>
                <HeadDetail class="w-12">Avatar</HeadDetail>
                <HeadDetail class="w-12">Platform</HeadDetail>
                <HeadDetail>Name</HeadDetail>
                <HeadDetail>Email</HeadDetail>
                <HeadDetail>Confirmed</HeadDetail>
                <HeadDetail>Comments</HeadDetail>
                <HeadDetail>Likes</HeadDetail>
                <HeadDetail>Last update</HeadDetail>
                <HeadDetail>Created</HeadDetail>
                <HeadDetail></HeadDetail>
            </template>

            <BodyRow v-for="user in users" :key="user.id">
                <BodyDetail>{{ user.id }}</BodyDetail>
                <BodyDetail>
                    <a :href="`/@${user.slug}`" rel="noreferrer">
                        <img
                            :src="`/storage/avatars/${user.avatar}`"
                            width="200"
                            height="200"
                            class="w-8 h-8 rounded-full"
                            :class="user.ban ? 'outline outline-red-500' : ''"
                            :alt="user.name"
                        />
                    </a>
                </BodyDetail>
                <BodyDetail>
                    <img
                        :src="`/storage/icons/${user.sign_in_method}.webp`"
                        width="100"
                        height="100"
                        class="w-8 h-8 rounded-full"
                    />
                </BodyDetail>
                <BodyDetail>{{ user.name }}</BodyDetail>
                <BodyDetail>{{ user.email }}</BodyDetail>
                <BodyDetail
                    :class="
                        user.email_verified_at ? 'text-green-700' : 'text-red-700'
                    "
                >
                    {{ user.email_verified_at ? 'Yes' : 'No' }}
                </BodyDetail>
                <BodyDetail
                    class="font-black"
                    :class="
                        user.comments_count && user.comments_count > 0
                            ? 'text-green-700'
                            : 'text-red-700'
                    "
                >
                    {{ user.comments_count }}
                </BodyDetail>
                <BodyDetail class="font-black"> 0 </BodyDetail>
                <BodyDetail>
                    {{ moment(user.updated_at).format('D MMM YYYY H:mm') }}
                </BodyDetail>
                <BodyDetail>
                    {{ moment(user.created_at).format('D MMM YYYY H:mm') }}
                </BodyDetail>
                <BodyDetail>
                    <router-link
                        :to="{ name: 'ban-user', params: { slug: user.slug } }"
                        class="transition-transform text-red-600 hover:scale-125"
                    >
                        <i class="fas fa-ban"></i>
                    </router-link>
                </BodyDetail>
            </BodyRow>
        </DataTable>

        <LoadMoreButton v-if="nextPageUrl" @clicked="loadMoreUsers" />
    </div>
</template>
