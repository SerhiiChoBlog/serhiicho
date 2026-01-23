<script setup lang="ts">
import type { Keyword, Post } from '@shared/types/models'
import { ref } from 'vue'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import FormInput from '@admin/components/Form/FormInput.vue'

const emit = defineEmits<{ (e: 'changed'): void }>()

const props = defineProps<{ post: Post }>()
const searchQuery = ref('')
const loading = ref<boolean>(false)

function addKeyword(): void {
    if (searchQuery.value === '') {
        return
    }

    loading.value = true

    axios
        .post(`/admin/ajax/keywords`, {
            name: searchQuery.value,
            post_id: props.post.id,
        })
        .then(_ => {
            searchQuery.value = ''
            emit('changed')
        })
        .catch(handleError)
        .finally(() => (loading.value = false))
}

function inputChanged(val: string): void {
    searchQuery.value = val
}

function deleteKeyword(keyword: Keyword): void {
    if (!confirm(`Are you sure you want to delete the keyword "${keyword.name}"?`)) {
        return
    }

    const params = {
        data: {
            keyword_id: keyword.id,
            post_id: props.post.id,
        },
    }

    axios
        .delete('/admin/ajax/keywords', params)
        .then(_ => emit('changed'))
        .catch(handleError)
}
</script>

<template>
    <div class="relative">
        <FormInput
            type="text"
            id="post-keywords"
            name="post-keywords"
            placeholder="Start typing to add a keyword..."
            @changed="inputChanged"
            @keydown.enter.prevent="addKeyword"
            :value="searchQuery"
        >
            Keywords:

            <template #input>
                <div
                    v-if="loading"
                    class="absolute bg-white inset-1 flex items-center pl-5 rounded-xs"
                >
                    <span>Loading... Please wait...</span>
                </div>
            </template>
        </FormInput>

        <div class="mt-2 flex flex-wrap gap-1">
            <div
                v-for="keyword in props.post.keywords"
                :key="keyword.name"
                class="inline-block bg-main text-white px-2 py-0.5 rounded-md"
            >
                <span>{{ keyword.name }}</span>

                <button
                    @click="deleteKeyword(keyword)"
                    type="button"
                    class="ml-2 hover:scale-110 hover:text-green-200"
                >
                    <i class="fas fa-times"></i>
                </button>
            </div>
        </div>
    </div>
</template>
