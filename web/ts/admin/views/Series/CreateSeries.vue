<script setup lang="ts">
import type { Series } from '@shared/types/models'
import { ref } from 'vue'
import axios from 'axios'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import FormInput from '@admin/components/Form/FormInput.vue'
import handleError from '@admin/modules/handleError'
import { useRouter } from 'vue-router'

const loading = ref(false)
const router = useRouter()
const formData = ref({
    title: '',
    intro: '',
})

function createSeries(): void {
    loading.value = true

    axios.post<Series>('/admin/ajax/series', formData.value)
        .then(resp => {
            router.push({
                name: 'series-edit',
                params: { id: resp.data.id },
            })
        })
        .catch(handleError)
        .finally(() => loading.value = false)
}
</script>

<template>
    <div class="container">
        <h1 class="text-3xl font-bold text-font-second">Create series</h1>
        <h2 class="text-xl">Enter the information for the future series</h2>

        <form @submit.prevent="createSeries" class="mt-5">
            <div class="max-w-[700px] space-y-5">
                <FormInput
                    type="text"
                    placeholder="Enter a title"
                    :length="190"
                    id="title"
                    @changed="(v: string) => (formData.title = v)"
                >
                </FormInput>

                <FormInput
                    type="text"
                    placeholder="Enter a short introduction"
                    :length="190"
                    id="intro"
                    @changed="(v: string) => (formData.intro = v)"
                >
                </FormInput>

                <MainButton>
                    <i
                        class="fas mr-2"
                        :class="loading ? 'fa-spinner animate-spin' : 'fa-plus'"
                    ></i>
                    <span>Create series</span>
                </MainButton>
            </div>
        </form>
    </div>
</template>
