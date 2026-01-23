<script setup lang="ts">
import type { Post } from '@shared/types/models'
import { ref } from 'vue'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import copy from 'copy-to-clipboard'
import showToast from '@shared/modules/showToast'
import useExportPostImages from '@admin/composables/useExportPostImages'
import useImportPostImages from '@admin/composables/useImportPostImages'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'
import FileInput from '@admin/components/Form/FileInput.vue'

const emit = defineEmits<{
    (e: 'uploaded'): void
    (e: 'removed'): void
}>()

const { post } = defineProps<{ post: Post }>()
const { downloadImages, exportLoading } = useExportPostImages(post.id)
const { importImages, importLoading } = useImportPostImages(post.id)
const loading = ref<boolean>(false)

function submitImage(file: File): void {
    loading.value = true

    const formData = new FormData()
    const params = { headers: { 'content-type': 'multipart/form-data' } }

    formData.append('image', file)

    axios
        .post<string>(`/admin/ajax/post-content-images/${post.id}`, formData, params)
        .then(resp => {
            showToast({ text: 'Image has been uploaded and copied into a buffer' })
            copy(resp.data)
            emit('uploaded')
        })
        .catch(handleError)
        .finally(() => (loading.value = false))
}

function deleteImage(id: number): void {
    if (!confirm('Sure you want to delete this content image?')) return

    loading.value = true

    axios
        .delete(`/admin/ajax/post-content-images/${id}`)
        .then(_ => {
            showToast({ text: 'Image has been removed' })
            emit('removed')
        })
        .catch(handleError)
        .finally(() => (loading.value = false))
}
</script>

<template>
    <div class="mt-6 pt-4">
        <h3 class="text-center text-2xl mb-10">Attached images for this post</h3>

        <DataTable>
            <template #head>
                <HeadDetail class="w-40">Image</HeadDetail>
                <HeadDetail>URI</HeadDetail>
                <HeadDetail class="w-40">Delete</HeadDetail>
            </template>

            <BodyRow v-if="post.images" v-for="image in post.images" :key="image.id">
                <BodyDetail>
                    <a :href="image.uri" target="_blank" rel="noreferrer">
                        <img :src="image.uri" alt="Post image" width="100" class="rounded-md shadow-xs" />
                    </a>
                </BodyDetail>

                <BodyDetail>
                    <input type="text" :value="'https://serhiicho.com' + image.uri"
                        class="border border-border bg-transparent text-font rounded-md p-2 w-full" disabled />
                </BodyDetail>

                <BodyDetail>
                    <form @submit.prevent="deleteImage(image.id)">
                        <MainButton type="submit" class="bg-red-500 hover:bg-red-600">
                            <i class="fas fa-trash mr-2"></i>
                            Delete
                        </MainButton>
                    </form>
                </BodyDetail>
            </BodyRow>
        </DataTable>

        <div class="mt-6 flex justify-between gap-3">
            <div class="inline-block mr-2">
                <FileInput @chosen="submitImage" name="image" :required="true">
                    {{ loading ? 'Uploading image...' : 'Choose a post image' }}
                </FileInput>
            </div>

            <div class="flex flex-col lg:flex-row gap-5 items-center">
                <MainButton @click="downloadImages" class="ml-3 !bg-gray-500 hover:!bg-gray-600"
                    :disabled="exportLoading">
                    <i :class="exportLoading
                        ? 'fa-spinner animate-spin'
                        : 'fa-arrow-down'
                        " class="fas mr-2"></i>
                    Export Images
                </MainButton>

                <FileInput name="import-images" :required="true" accept=".zip" @chosen="importImages">
                    {{ importLoading ? 'Importing images...' : 'Choose images archive' }}
                </FileInput>
            </div>
        </div>
    </div>
</template>
