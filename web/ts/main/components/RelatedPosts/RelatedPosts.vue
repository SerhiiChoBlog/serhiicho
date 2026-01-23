<script setup lang="ts">
import type { Post } from '@shared/types/models'
import handleServerError from '@shared/modules/handleServerError'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import Spinner from '@shared/components/Icons/Spinner.vue'
import PostItem from '@/components/RelatedPosts/PostItem.vue'

type Props = {
    postId: number
}

const { postId } = defineProps<Props>()
const posts = ref<Post[]>([])
const loading = ref(false)

onMounted(() => {
    loading.value = true
    fetchPost()
})

function fetchPost(): void {
    axios
        .get<Post[]>(`/api/posts/related/${postId}`)
        .then(resp => (posts.value = resp.data))
        .catch(err => handleServerError(err))
        .finally(() => (loading.value = false))
}
</script>

<template>
    <div v-if="loading" class="text-center pt-5">
        <spinner />
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-3 gap-2">
        <post-item v-for="post in posts" :key="post.id" :post="post" />
    </div>
</template>
