<script setup lang="ts">
import type { Series } from '@shared/types/models'
import type { EditSeriesFormElements } from '@shared/types'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { tinymceConfig } from '@admin/modules/tinymce'
import SaveSeries from '@admin/modules/SaveSeries'
import FormTextarea from '@admin/components/Form/FormTextarea.vue'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import FormInput from '@admin/components/Form/FormInput.vue'
import FormImage from '@admin/components/FormImage.vue'
import DeleteSeries from '@admin/components/Series/DeleteSeries.vue'
import SeriesPosts from '@admin/views/Series/Posts/SeriesPosts.vue'
import Editor from '@tinymce/tinymce-vue'

onMounted(() => {
    fetchSeries()
    document.addEventListener('keydown', saveSeriesWithCorrectKeyBinds)
})

const route = useRoute()
const loading = ref(true)
const series = ref<Series | null>(null)
const form = ref<HTMLFormElement | null>(null)

function fetchSeries(): void {
    loading.value = true

    axios.get<Series>(`/admin/ajax/series/${route.params.id}/edit`)
        .then(resp => series.value = resp.data)
        .catch(handleError)
        .finally(() => loading.value = false)
}

function saveSeries(form: HTMLFormElement | null): void {
    if (!series.value || !form)
        return

    loading.value = true

    const elements = form.elements as EditSeriesFormElements

    new SaveSeries(elements, series.value.id)
        .save(series.value.description)
        .then(handleSavePostResponse)
        .finally(() => loading.value = false)
}

function handleSavePostResponse(newSeries: Series | null): void {
    if (!series.value || !newSeries) {
        return
    }

    series.value = newSeries
}

function saveSeriesWithCorrectKeyBinds(e: KeyboardEvent): void {
    if ((e.ctrlKey || e.metaKey) && e.code === 'KeyS') {
        e.preventDefault()
        saveSeries(form.value)
    }
}
</script>

<template>
    <div class="container pb-6">
        <div v-if="series">
            <h1 class="text-4xl text-font-second">{{ series.title }}</h1>

            <div class="container">
                <form
                    @submit.prevent="saveSeries(form)"
                    enctype="multipart/form-data"
                    ref="form"
                >
                    <input type="hidden" name="tag_id">

                    <div class="flex flex-col lg:flex-row gap-6 mt-10">
                        <div class="flex-1">
                            <div>
                                <FormInput
                                    type="text"
                                    id="title"
                                    placeholder="Enter the title"
                                    name="title"
                                    :length="190"
                                    :value="series.title"
                                >
                                    Title
                                </FormInput>
                            </div>

                            <div class="mt-5">
                                <FormTextarea
                                    id="intro"
                                    placeholder="Enter a short intro"
                                    classes="min-h-[70px]"
                                    :value="series.intro"
                                    :length="190"
                                    name="intro"
                                >
                                    Intro
                                </FormTextarea>
                            </div>

                            <div class="grid grid-cols-1 lg:grid-cols-2 gap-6 mt-3">
                                <div>
                                    <FormInput
                                        type="color"
                                        id="color_from"
                                        name="color_from"
                                        :value="series.color_from"
                                    >
                                        Color from
                                    </FormInput>
                                </div>
                                <div>
                                    <FormInput
                                        type="color"
                                        id="color_to"
                                        name="color_to"
                                        :value="series.color_to"
                                    >
                                        Color to
                                    </FormInput>
                                </div>
                            </div>
                        </div>

                        <div
                            class="flex-1 px-4 rounded-lg flex items-center"
                            :style="{ 'background': `linear-gradient(90deg, ${series.color_from} 0%, ${series.color_to} 100%)` }"
                        >
                            <FormImage
                                :uri="series.image_lg"
                                defaultUri="/storage/series/placeholder.jpg"
                            />
                        </div>
                    </div>

                    <div class="mt-8">
                        <Editor
                            :initialValue="series.description"
                            tinymce-script-src="/js/tinymce/tinymce.min.js"
                            :init="tinymceConfig"
                            v-model="series.description"
                            @keydown="saveSeriesWithCorrectKeyBinds"
                        />
                    </div>

                    <div>
                        <div class="flex gap-2 fixed right-[20px] top-14 z-[9999] opacity-90">
                            <MainButton class="bg-green-500 hover:bg-green-600">
                                <i v-if="loading" class="fas fa-spinner animate-spin mr-2"></i>
                                <i v-else class="fas fa-save mr-2"></i>
                                Save
                            </MainButton>

                            <MainButton :href="`/series/${series.slug}`" target="_blank">
                                <i class="fas fa-eye mr-2"></i>
                                Preview
                            </MainButton>
                        </div>

                        <DeleteSeries :series="series" />
                    </div>

                    <div class="mt-6">
                        <SeriesPosts
                            v-if="series.posts"
                            :posts="series.posts"
                            :seriesTitle="series.title"
                        />
                    </div>
                </form>
            </div>
        </div>
    </div>
</template>
