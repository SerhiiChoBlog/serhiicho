<script setup lang="ts">
import type { Post } from '@shared/types/models'
import { ref } from 'vue'
import handleError from '@admin/modules/handleError'
import axios from 'axios'
import MainButton from '@admin/components/Buttons/MainButton.vue'

const emit = defineEmits<{
    (e: 'postAdded', post: Post): void
}>()

const postSlug = ref<string>('')
const post = ref<Post | null>(null)
const loading = ref<boolean>(false)

function fetchPost(): void {
    if (postSlug.value === '') {
        return
    }

    loading.value = true

    if (postSlug.value.includes('/')) {
        postSlug.value = postSlug.value.split('/').pop() || ''
    }

    axios
        .get<Post | null>(`/admin/ajax/posts/${postSlug.value}`)
        .then(resp => {
            if (!resp.data) {
                return
            }

            post.value = resp.data
            emit('postAdded', resp.data)
            postSlug.value = ''
        })
        .catch(handleError)
        .finally(() => (loading.value = false))
}
</script>

<template>
    <div class="flex flex-col lg:flex-row gap-4 items-center mt-7">
        <MainButton>
            <i class="fas fa-plus mr-2"></i>
            Add post
        </MainButton>

        <input
            type="text"
            placeholder="Paste post slug or URL"
            class="px-5 py-2 rounded-md bg-page-second w-full max-w-[500px]"
            v-model="postSlug"
            @blur="fetchPost"
        />

        <i
            v-if="loading"
            class="fas fa-spinner animate-spin text-gray-400 text-2xl"
        ></i>
        <i v-else-if="post" class="fas fa-check text-green-500 text-2xl"></i>
        <i v-else class="fas fa-times text-red-500 text-2xl"></i>

        <span v-if="loading">Loading...</span>
        <span v-else-if="post">Added: "{{ post.title }}"</span>
        <span v-else>No post</span>
    </div>
</template>
