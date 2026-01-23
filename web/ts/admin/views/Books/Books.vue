<script setup lang="ts">
import type { Book } from '@shared/types/models'
import type { Pagination } from '@shared/types/response'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import handleError from '@admin/modules/handleError'
import useUpdateBookField from '@admin/composables/books/useUpdateBookField'
import PageTitle from '@admin/components/PageTitle.vue'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import Spinner from '@shared/components/Icons/Spinner.vue'
import DataTable from '@admin/components/DataTable/DataTable.vue'
import HeadDetail from '@admin/components/DataTable/HeadDetail.vue'
import BookItem from '@admin/components/Books/BookItem.vue'

const books = ref<Book[]>([])
const loading = ref<boolean>(true)

onMounted(fetchBooks)

const { updateBookField } = useUpdateBookField()

function fetchBooks(): void {
    loading.value = true

    axios
        .get<Pagination<Book[]>>('/admin/ajax/books')
        .then(resp => (books.value = resp.data.data))
        .catch(handleError)
        .finally(() => (loading.value = false))
}
</script>

<template>
    <div class="container">
        <PageTitle>
            <div class="flex items-center justify-center gap-3">
                <span>Books</span>

                <MainButton
                    type="button"
                    :route="{ name: 'books-create' }"
                    title="Create a new book"
                >
                    <i class="fas fa-plus"></i>
                </MainButton>
            </div>
        </PageTitle>

        <div class="mt-5">
            <Spinner v-if="loading" />

            <h2 v-else-if="books.length === 0" class="text-center text-xl mt-6">
                There are not books yet
            </h2>

            <DataTable v-else>
                <template #head>
                    <HeadDetail
                        v-for="title in [
                            'Image',
                            'Title',
                            'Author',
                            'Rate',
                            'Read at',
                            '',
                        ]"
                        :key="title"
                    >
                        {{ title }}
                    </HeadDetail>
                </template>

                <BookItem
                    v-for="book in books"
                    :key="book.id"
                    :book="book"
                    @field-updated="updateBookField"
                    @refresh="fetchBooks"
                />
            </DataTable>
        </div>
    </div>
</template>
