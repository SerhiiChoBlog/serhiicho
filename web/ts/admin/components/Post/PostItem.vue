<script setup lang="ts">
import type { Post } from '@shared/types/models'
import IconButton from '@admin/components/Buttons/IconButton.vue'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'

defineProps<{ post: Post }>()
</script>

<template>
    <BodyRow>
        <BodyDetail>{{ post.id }}</BodyDetail>
        <BodyDetail>
            <i v-if="post.is_published" class="fas fa-check text-green-600"></i>
            <i v-else class="fas fa-times text-red-600"></i>
        </BodyDetail>
        <BodyDetail>
            <a :href="`/${post.image_lg}`" target="_blank" rel="noreferrer">
                <img
                    :src="`/${post.image_sm}`"
                    width="300"
                    class="rounded-md shadow-md"
                    :alt="post.title"
                />
            </a>
        </BodyDetail>
        <BodyDetail class="text-xl leading-6">{{ post.title }}</BodyDetail>
        <BodyDetail class="text-sm">{{ post.intro || '-' }}</BodyDetail>
        <BodyDetail class="text-center">{{ post.social_shares_count }}</BodyDetail>
        <BodyDetail class="text-center">{{ post.post_views_count }}</BodyDetail>
        <BodyDetail class="text-center">
            <icon-button
                icon="fa-eye"
                :href="`/posts/${post.slug}`"
                title="View post"
            />
        </BodyDetail>
        <BodyDetail class="text-center">
            <icon-button
                icon="fa-pen"
                :route="{ name: 'posts-edit', params: { id: post.id.toString() } }"
                title="Edit post"
            />
        </BodyDetail>
    </BodyRow>
</template>
