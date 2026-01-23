<script setup lang="ts">
import type { Post } from '@shared/types/models'
import AddPostToSeries from '@admin/views/Series/Posts/AddPostToSeries.vue'
import PostSeriesButton from '@admin/views/Series/Posts/PostSeriesButton.vue'

const { posts, seriesTitle } = defineProps<{
    seriesTitle: string
    posts: Post[]
}>()

function addPostToSeries(post: Post): void {
    posts.push(post)
}

function movePostUpInOrder(post: Post): void {
    const index = posts.indexOf(post)

    if (index === 0) {
        return
    }

    posts.splice(index, 1)
    posts.splice(index - 1, 0, post)
}

function movePostDownInOrder(post: Post): void {
    const index = posts.indexOf(post)

    if (index === posts.length - 1) {
        return
    }

    posts.splice(index, 1)
    posts.splice(index + 1, 0, post)
}

function deletePost(post: Post): void {
    const index = posts.indexOf(post)
    posts.splice(index, 1)
}
</script>

<template>
    <h2 class="text-3xl mb-5">Related posts</h2>

    <h3 v-if="posts.length === 0" class="text-xl">🙅🏻‍♂️ No related posts yet</h3>

    <div v-else class="space-y-3">
        <article
            v-for="(post, index) in posts"
            :key="post.id"
            class="bg-page-second p-3 rounded-lg flex flex-col justify-between lg:flex-row gap-6"
        >
            <div class="flex flex-col lg:flex-row items-center gap-4">
                <input type="hidden" name="posts[]" :value="post.id" />

                <img
                    :src="`/${post.image_sm}`"
                    class="w-full lg:w-48 rounded-md shadow-md"
                    width="550"
                    height="275"
                />

                <div class="flex flex-col justify-center">
                    <a :href="`/posts/${post.slug}`" class="hover:underline">
                        <h2 class="text-xl lg:text-2xl">
                            {{ seriesTitle }}: {{ post.title }}
                        </h2>
                    </a>
                    <h3 class="text-md lg:text-lg opacity-80">Part {{ ++index }}</h3>
                    <p class="text-sm lg:text-md">{{ post.intro }}</p>
                </div>
            </div>

            <div
                class="mx-2 flex flex-row lg:flex-col justify-center gap-6 lg:gap-2"
            >
                <PostSeriesButton
                    v-if="index !== 1"
                    @click="movePostUpInOrder(post)"
                    icon="fa-arrow-up"
                    class="text-gray-500"
                />

                <PostSeriesButton
                    @click="deletePost(post)"
                    icon="fa-times"
                    class="text-red-500"
                />

                <PostSeriesButton
                    v-if="index !== posts.length"
                    @click="movePostDownInOrder(post)"
                    icon="fa-arrow-down"
                    class="text-gray-500"
                />
            </div>
        </article>
    </div>

    <AddPostToSeries @postAdded="addPostToSeries" />
</template>
