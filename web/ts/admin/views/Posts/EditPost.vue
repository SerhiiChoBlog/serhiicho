<script setup lang="ts">
import type { Post, Tag } from '@shared/types/models'
import type { EditPostFormElements } from '@shared/types'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'
import { tinymceConfig } from '@admin/modules/tinymce'
import SavePost from '@admin/modules/SavePost'
import PostContentImages from '@admin/components/Post/PostContentImages.vue'
import FormTextarea from '@admin/components/Form/FormTextarea.vue'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import FormInput from '@admin/components/Form/FormInput.vue'
import TagItem from '@admin/components/TagItem.vue'
import TagSuggestions from '@admin/components/Post/TagSuggestions.vue'
import FormImage from '@admin/components/FormImage.vue'
import DeletePost from '@admin/components/Post/DeletePost.vue'
import Editor from '@tinymce/tinymce-vue'
import KeywordSuggestions from '@admin/components/Post/KeywordSuggestions.vue'

onMounted(() => {
    fetchPost()
    document.addEventListener('keydown', savePostWithCorrectKeyBinds)
})

const route = useRoute()
const loading = ref(true)
const post = ref<Post | null>(null)
const tagInputValue = ref<string>('')
const form = ref<HTMLFormElement | null>(null)
const mainTagId = ref<number | null>(null)

function fetchPost(): void {
    loading.value = true

    axios
        .get<Post>(`/admin/ajax/posts/${route.params.id}/edit`)
        .then(resp => {
            post.value = resp.data
            setMainTagId()
        })
        .catch(handleError)
        .finally(() => (loading.value = false))
}

function setMainTagId(): void {
    if (!post.value || !post.value.tags || post.value.tags.length === 0) {
        return
    }

    const mainTag = post.value.tags.find(tag => tag.pivot && tag.pivot.is_main)

    if (mainTag) {
        mainTagId.value = mainTag.id
        return
    }

    mainTagId.value = post.value.tags[0].id
}

function savePost(form: HTMLFormElement | null): void {
    if (!post.value || !form) {
        return
    }

    loading.value = true

    if (!mainTagId.value) {
        setMainTagId()
    }

    const elements = form.elements as EditPostFormElements

    new SavePost(elements, post.value.id, mainTagId.value!)
        .save(post.value.content || '')
        .then(handleSavePostResponse)
        .finally(() => (loading.value = false))
}

function handleSavePostResponse(newPost: Post | null): void {
    if (!post.value || !newPost) {
        return
    }

    post.value = newPost
}

function updateTagId(id: number): void {
    if (!post.value || !form.value) {
        return
    }

    mainTagId.value = id

    savePost(form.value)
}

function deleteTag(id: number): void {
    if (!post.value || !post.value.tags) return

    post.value.tags = post.value.tags.filter(tag => tag.id !== id)
}

function addTag(tag: Tag): void {
    if (!post.value || !post.value.tags) return

    post.value.tags.push(tag)
    tagInputValue.value = ''
}

function savePostWithCorrectKeyBinds(e: KeyboardEvent): void {
    if ((e.ctrlKey || e.metaKey) && e.code === 'KeyS') {
        e.preventDefault()
        savePost(form.value)
    }
}
</script>

<template>
    <div class="container pb-6">
        <div v-if="post">
            <h1 class="text-4xl text-font-second">{{ post.title }}</h1>

            <div class="container">
                <form
                    @submit.prevent="savePost(form)"
                    enctype="multipart/form-data"
                    ref="form"
                >
                    <div class="flex flex-col lg:flex-row gap-6 mt-10">
                        <div class="flex-1">
                            <div>
                                <FormInput
                                    type="text"
                                    id="title"
                                    placeholder="Enter the title"
                                    name="title"
                                    :length="190"
                                    :value="post.title"
                                >
                                    Title
                                </FormInput>
                            </div>

                            <div class="mt-5">
                                <FormInput
                                    type="text"
                                    id="post-tags"
                                    name="post-tags"
                                    placeholder="Choose tags"
                                    @changed="(val: string) => tagInputValue = val"
                                    :value="tagInputValue"
                                >
                                    Tags

                                    <div class="ml-1 my-2 inline-block space-x-2">
                                        <template v-if="post && post.tags">
                                            <TagItem
                                                v-for="tag in post.tags"
                                                :key="tag.id"
                                                :tag="tag"
                                                @deleteTag="deleteTag"
                                                @updateTagId="updateTagId"
                                            />
                                        </template>
                                    </div>
                                </FormInput>

                                <TagSuggestions
                                    @addTag="addTag"
                                    :value="tagInputValue"
                                />
                            </div>

                            <div
                                class="flex items-center justify-between gap-8 mt-5"
                            >
                                <div class="w-full">
                                    <FormInput
                                        type="text"
                                        id="video"
                                        placeholder="Enter a YouTube video ID"
                                        classes=""
                                        :value="
                                            post.video ? post.video.id : undefined
                                        "
                                        :length="190"
                                        name="video"
                                    >
                                        YouTube video ID
                                    </FormInput>
                                </div>

                                <a
                                    v-if="post.video"
                                    :href="post.video.url"
                                    target="_blank"
                                    rel="noreferrer"
                                >
                                    <img
                                        :src="post.video.image"
                                        :alt="post.video.title"
                                        width="480"
                                        height="360"
                                        class="max-w-[140px] rounded-md shadow-md"
                                    />
                                </a>
                            </div>

                            <div class="mt-5">
                                <KeywordSuggestions
                                    @changed="fetchPost"
                                    :post="post"
                                />
                            </div>
                        </div>

                        <div class="flex-1">
                            <FormImage
                                :uri="post.image_lg"
                                defaultUri="/storage/posts/placeholder.jpg"
                            />
                        </div>
                    </div>

                    <div class="mt-5">
                        <FormTextarea
                            id="intro"
                            placeholder="Enter a short intro"
                            classes="min-h-[70px]"
                            :value="post.intro || undefined"
                            :length="190"
                            name="intro"
                        >
                            Short intro
                        </FormTextarea>
                    </div>

                    <div class="mt-5">
                        <Editor
                            :initialValue="post.content"
                            tinymce-script-src="/js/tinymce/tinymce.min.js"
                            :init="tinymceConfig"
                            v-model="post!.content"
                            @keydown="savePostWithCorrectKeyBinds"
                        />
                    </div>

                    <div>
                        <div
                            class="flex flex-col justify-center fixed right-[20px] top-20 z-[9999] space-y-2 opacity-85"
                        >
                            <MainButton class="bg-green-500 hover:bg-green-600">
                                <i
                                    v-if="loading"
                                    class="fas fa-spinner animate-spin mr-2"
                                ></i>
                                <i v-else class="fas fa-save mr-2"></i>
                                Save
                            </MainButton>

                            <MainButton
                                :href="`/posts/${post.slug}`"
                                target="_blank"
                            >
                                <i class="fas fa-eye mr-2"></i>
                                Preview
                            </MainButton>
                        </div>

                        <DeletePost :post="post" />
                    </div>
                </form>
            </div>

            <PostContentImages
                :post="post"
                @uploaded="fetchPost"
                @removed="fetchPost"
            />
        </div>
    </div>
</template>
