<script setup lang="ts">
import type { Tag } from '@shared/types/models'
import { onMounted, ref } from 'vue'
import handleError from '@admin/modules/handleError'
import axios from 'axios'
import PageTitle from '@admin/components/PageTitle.vue'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'

const tags = ref<Tag[]>([])
const loading = ref(true)

onMounted(() => fetchTags())

function fetchTags(): void {
    loading.value = true

    axios.get<Tag[]>('/admin/ajax/tags')
        .then(resp => tags.value = resp.data)
        .catch(handleError)
        .finally(() => loading.value = false)
}
</script>

<template>
    <div class="container">
        <PageTitle>
            <div class="flex items-center justify-center gap-3">
                <span>Tags</span>

                <MainButton type="button" :route="{ name: 'tags-create' }" title="Add new tag">
                    <i class="fas fa-plus"></i>
                </MainButton>
            </div>
        </PageTitle>

        <div class="mt-5">
            <Spinner v-if="loading" />

            <h2 v-else-if="tags.length === 0" class="text-center text-xl mt-6">
                There are not tags
            </h2>

            <DataTable v-else>
                <template #head>
                    <HeadDetail
                        v-for="title in ['Img', 'Title', 'Slug', 'Posts', 'Color', 'Tag', 'Description', '']"
                        :key="title"
                    >
                        {{ title }}
                    </HeadDetail>
                </template>

                <BodyRow v-for="tag in tags" :key="tag.id">
                    <BodyDetail>
                        <div class="bg-white w-[28px] aspect-square flex items-center justify-center rounded-full">
                            <img
                                :src="`/storage/tags/${tag.name}.webp`"
                                :alt="`Logo of ${tag.name}`"
                                width="200"
                                height="200"
                                class="w-[18px]"
                            />
                        </div>
                    </BodyDetail>
                    <BodyDetail>{{ tag.title }}</BodyDetail>
                    <BodyDetail class="min-w-[165px]">{{ tag.name }}</BodyDetail>
                    <BodyDetail>{{ tag.posts_count }}</BodyDetail>
                    <BodyDetail>
                        <input
                            type="text"
                            :value="tag.color"
                            class="w-[80px] border border-border bg-transparent text-center rounded-md"
                        >
                    </BodyDetail>
                    <BodyDetail class="min-w-[140px]">
                        <a
                            :href="`/posts?tag=${tag.name}`"
                            class="py-[3px] inline-block font-white px-2 rounded-md text-white text-xs"
                            :style="{ backgroundColor: tag.color }"
                        >
                            {{ tag.name }}
                        </a>
                    </BodyDetail>
                    <BodyDetail>
                        {{ tag.description }}
                    </BodyDetail>
                    <BodyDetail>
                        <RouterLink
                            :to="{ name: 'tags-edit', params: { id: tag.id } }"
                            class="text-font-second inline-block text-xs transition-colors hover:text-main p-1"
                            title="Edit tag"
                        >
                            <i class="fas fa-pen"></i>
                        </RouterLink>
                    </BodyDetail>
                </BodyRow>
            </DataTable>
        </div>
    </div>
</template>
