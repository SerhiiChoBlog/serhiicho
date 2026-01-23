<script setup lang="ts">
import type { Post } from '@shared/types/models'
import { onMounted, ref } from 'vue'
import handleServerError from '@shared/modules/handleServerError'
import axios from 'axios'
import Spinner from '@shared/components/Icons/Spinner.vue'
import Btn from '@/components/PrevNextPost/Btn.vue'

onMounted(() => {
    loading.value = true
    fetchPrevNextPosts()
})

const { postId } = defineProps<{ postId: number }>()
const loading = ref(false)
const prevPost = ref<Post | null>(null)
const nextPost = ref<Post | null>(null)

function fetchPrevNextPosts() {
    loading.value = false

    axios
        .get(`/api/posts/prev-next/${postId}`)
        .then(res => {
            prevPost.value = res.data.prev
            nextPost.value = res.data.next
        })
        .catch(handleServerError)
}
</script>

<template>
    <div v-if="loading" class="text-center pt-5">
        <spinner />
    </div>

    <div v-else class="grid grid-cols-1 md:grid-cols-2 mt-4">
        <btn :post="prevPost" class="text-left">Previous article</btn>
        <btn :post="nextPost" class="text-right">Next article</btn>
    </div>
</template>
