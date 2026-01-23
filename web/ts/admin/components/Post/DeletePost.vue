<script setup lang="ts">
import type { Post } from '@shared/types/models'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import handleError from '@admin/modules/handleError'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { ref } from 'vue'
import showToast from '@shared/modules/showToast'

const { post } = defineProps<{ post: Post }>()
const router = useRouter()
const loading = ref(false)

function deletePost(): void {
    if (!confirm(`STEP 1/3. Are you sure you want to delete post "${post.title}"`)) {
        return
    }

    const msg = `STEP 2/3. You will not be able to restore this post. Are you sure you want to remove it forever?`

    if (!confirm(msg)) {
        return
    }

    if (!confirm(`STEP 3/3. Last confirmation. Are you sure?`)) {
        return
    }

    loading.value = true

    axios
        .delete(`/admin/ajax/posts/${post.id}`)
        .then(() => {
            router.push({ name: 'posts' })
            showToast({ text: `The post "${post.title}" has been removed` })
        })
        .catch(handleError)
        .finally(() => (loading.value = false))
}
</script>

<template>
    <div class="mt-7">
        <MainButton
            @click="deletePost"
            class="bg-red-500 hover:bg-red-600"
            type="button"
        >
            <i class="fas fa-trash mr-2"></i>
            Delete
        </MainButton>

        <label class="checkbox">
            <input
                type="checkbox"
                name="is-published"
                class="ml-3 mr-1"
                :checked="post.is_published"
            />

            <span>publish this post</span>
        </label>
    </div>
</template>
