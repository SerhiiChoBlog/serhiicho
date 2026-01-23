<script setup lang="ts">
import type { Series as SeriesType } from '@shared/types/models'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import FormInput from '@admin/components/Form/FormInput.vue'
import handleError from '@admin/modules/handleError'
import { useRouter } from 'vue-router'

const loading = ref(false)
const router = useRouter()
const title = ref<string>('')
const seriesId = ref<number | null>(null)
const series = ref<SeriesType[]>([])

onMounted(() => fetchSeries())

function createPost(): void {
    loading.value = true

    axios.post<number>('/admin/ajax/posts', {
        title: title.value,
        series: seriesId.value,
    })
        .then(resp => {
            router.push({
                name: 'posts-edit',
                params: { id: resp.data },
            })
        })
        .catch(handleError)
        .finally(() => loading.value = false)
}

function fetchSeries(): void {
    axios.get<SeriesType[]>('/admin/ajax/series/names')
        .then(resp => series.value = resp.data)
        .catch(handleError)
}
</script>

<template>
    <div class="container">
        <h1 class="text-3xl font-bold text-font-second">Basic information</h1>
        <h2 class="text-xl">Enter title for the future blog post.</h2>

        <form @submit.prevent="createPost" class="mt-5">
            <div class="grid grid-cols-1 lg:grid-cols-[3fr_1fr] gap-5 max-w-3xl">
                <FormInput
                    type="text"
                    placeholder="Enter the title"
                    :length="190"
                    id="title"
                    @changed="(value: string) => (title = value)"
                >
                </FormInput>

                <MainButton>
                    <i
                        class="fas mr-2"
                        :class="loading ? 'fa-spinner animate-spin' : 'fa-plus'"
                    ></i>
                    <span>Create post</span>
                </MainButton>
            </div>

            <div class="mt-5">
                <label for="series">Select series</label>
                <select
                    name="series"
                    class="block px-4 py-1 w-full max-w-[200px] border border-border rounded-md shadow-xs focus:outline-hidden sm:text-md dark:bg-page-second"
                    v-model="seriesId"
                >
                    <option :value="null">No series</option>
                    <option v-for="s in series" :value="s.id">
                        {{ s.title }}
                    </option>
                </select>
            </div>
        </form>
    </div>
</template>
