<script setup lang="ts">
import type { Comment } from '@shared/types/models'
import moment from 'moment'
import PageTitle from '@admin/components/PageTitle.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'
import useComments from '@admin/composables/useComments'
import useToggleApprove from '@admin/composables/useToggleApprove'

const {
    comments,
    totalComments,
    loading,
    nextPageUrl,
    deleteComment,
    loadMoreComments,
    fetchComments,
} = useComments()

const { toggleApproved } = useToggleApprove()

function toggleApprovingComment(c: Comment): void {
    toggleApproved(c)
    fetchComments('/admin/ajax/comments', true)
}
</script>

<template>
    <div class="container">
        <PageTitle>Comments <b v-if="totalComments">{{ totalComments }}</b></PageTitle>

        <Spinner v-if="loading" />

        <DataTable v-else>
            <template #head>
                <HeadDetail>#</HeadDetail>
                <HeadDetail class="w-64">Post name</HeadDetail>
                <HeadDetail class="min-w-[170px]">User</HeadDetail>
                <HeadDetail>Reply to</HeadDetail>
                <HeadDetail>Comment</HeadDetail>
                <HeadDetail class="w-20">Likes</HeadDetail>
                <HeadDetail class="w-20">Created</HeadDetail>
                <HeadDetail class="w-8"></HeadDetail>
            </template>

            <BodyRow v-for="comment in comments" :key="comment.id">
                <BodyDetail class="text-center">
                    <h2>{{ comment.id }}</h2>

                    <button @click="toggleApprovingComment(comment)">
                        <i
                            class="fas text-md lg:text-2xl"
                            :class="comment.is_approved ? 'fa-toggle-on text-green-600' : 'fa-toggle-off text-red-600'"
                        ></i>
                    </button>
                </BodyDetail>
                <BodyDetail>
                    <a :href="`/posts/${comment.post?.slug}`" target="_blank">
                        {{ comment.post?.title }}
                    </a>
                </BodyDetail>
                <BodyDetail>
                    <a :href="`/@${comment.user?.slug}`" rel="noreferrer" class="flex items-center gap-3">
                        <img
                            class="inline-block w-10 h-10 rounded-full shadow-md"
                            :src="`/storage/avatars/${comment.user?.avatar}`"
                            :alt="comment.user?.name"
                        />

                        <div>
                            <h2>{{ comment.user?.name }}</h2>
                            <h3>Visitor {{ comment.visitor_id || 'no' }}</h3>
                        </div>
                    </a>
                </BodyDetail>
                <BodyDetail class="align-baseline">
                    <span v-if="comment.parent">{{ comment.parent.comment }}</span>
                </BodyDetail>
                <BodyDetail class="align-baseline">
                    {{ comment.comment }}
                </BodyDetail>
                <BodyDetail>
                    {{ comment.likes_count }}
                </BodyDetail>
                <BodyDetail>
                    {{ moment(comment.created_at).format('D MMM YYYY H:mm') }}
                </BodyDetail>
                <BodyDetail>
                    <router-link
                        :to="{ name: 'ban-user', params: { slug: comment.user?.slug } }"
                        class="transition-transform text-red-600 hover:scale-125"
                    >
                        <i class="fas fa-ban"></i>
                    </router-link>

                    <button
                        @click="deleteComment(comment.id)"
                        type="button"
                        class="transition-transform hover:scale-125 mt-2"
                    >
                        <i class="fas fa-trash"></i>
                    </button>
                </BodyDetail>
            </BodyRow>
        </DataTable>

        <!-- Load more button -->
        <div class="text-center mt-7">
            <button
                v-if="nextPageUrl"
                class="border border-border px-7 py-2 rounded-full uppercase"
                @click="loadMoreComments"
            >
                Load more
            </button>
        </div>
    </div>
</template>
