<script setup lang="ts">
import { reactive, ref } from 'vue'
import handleError from '@admin/modules/handleError'
import showToast from '@shared/modules/showToast'
import axios from 'axios'
import PageTitle from '@admin/components/PageTitle.vue'
import FormInput from '@admin/components/Form/FormInput.vue'
import FormTextarea from '@admin/components/Form/FormTextarea.vue'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import { useRouter } from 'vue-router'

const loading = ref(false)
const router = useRouter()
const form = reactive({
    title: '',
    name: '',
    color: '',
    description: '',
})

function createTag(): void {
    loading.value = true

    axios.post('/admin/ajax/tags', form)
        .then(resp => {
            showToast({ text: `Tag "${form.name}" has been created` })
            router.push({ name: 'tags' })
        })
        .catch(handleError)
        .finally(() => loading.value = false)
}
</script>

<template>
    <div class="container">
        <PageTitle>Create a new tag</PageTitle>

        <Spinner v-if="loading" />

        <form v-else @submit="createTag" class="max-w-[1000px] mx-auto">
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-5">
                <FormInput
                    @changed="title => form.title = title"
                    id="title"
                    type="text"
                    placeholder="Enter the title"
                >
                    Full title
                </FormInput>

                <FormInput
                    @changed="name => form.name = name"
                    id="name"
                    type="text"
                    placeholder="Enter the slug"
                >
                    Short name (slug)
                </FormInput>

                <FormInput
                    @changed="color => form.color = color"
                    id="color"
                    type="color"
                    classes="min-h-[50px]"
                >
                    Choose the color
                </FormInput>
            </div>

            <div class="mt-4">
                <FormTextarea
                    type="text"
                    placeholder="Enter the description"
                    :length="190"
                    id="description"
                    @changed="desc => form.description = desc"
                >
                    Enter the description
                </FormTextarea>
            </div>

            <div class="mt-7">
                <MainButton type="submit">
                    <i class="fas fa-plus mr-2"></i>
                    Create tag
                </MainButton>
            </div>
        </form>
    </div>
</template>
