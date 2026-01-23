<script setup lang="ts">
import type { UpdateBookField } from '@admin/types'
import type { Book } from '@shared/types/models'
import { ref } from 'vue'
import axios from 'axios'
import showToast from '@shared/modules/showToast'
import handleError from '@admin/modules/handleError'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'

const emit = defineEmits<{
    (e: 'field-updated', slug: string, field: UpdateBookField, value: string): void
    (e: 'refresh'): void
}>()

const { book } = defineProps<{ book: Book }>()
const loading = ref<boolean>(false)

function fieldUpdated(e: Event, fieldName: UpdateBookField): void {
    const value = (e.target as HTMLInputElement).value
    emit('field-updated', book.slug, fieldName, value)
}

function deleteBook(): void {
    if (!confirm('Are you sure you want to delete this book?')) {
        return
    }

    loading.value = true

    axios
        .delete<string>(`/admin/ajax/books/${book.slug}`)
        .then(resp => {
            showToast({ text: resp.data })
            emit('refresh')
        })
        .catch(handleError)
        .finally(() => (loading.value = false))
}
</script>

<template>
    <BodyRow>
        <BodyDetail class="w-28">
            <img
                :src="`/storage/books/${book.image}`"
                :alt="`Logo of ${book.title}`"
                width="378"
                height="300"
                class="h-20 w-auto rounded-xs"
            />
        </BodyDetail>
        <BodyDetail>
            <input
                type="text"
                :value="book.title"
                class="w-full bg-transparent border-none py-1 px-2 text-lg"
                @blur="e => fieldUpdated(e, 'title')"
            />
        </BodyDetail>
        <BodyDetail>
            {{ book.author!.name }}
        </BodyDetail>
        <BodyDetail>
            <input
                type="number"
                min="1"
                max="5"
                :value="book.rate"
                class="w-14 bg-transparent border-none text-center p-1"
                @input="e => fieldUpdated(e, 'rate')"
            />
        </BodyDetail>
        <BodyDetail>
            <input
                type="date"
                :value="book.read_at"
                class="bg-transparent border-none py-1 px-3"
                pattern="\d{4}-\d{2}-\d{2}"
                @blur="e => fieldUpdated(e, 'read_at')"
            />
        </BodyDetail>
        <BodyDetail>
            <button @click="deleteBook" type="button">
                <i class="fas fa-trash text-red-600 hover:text-red-700"></i>
            </button>
        </BodyDetail>
    </BodyRow>
</template>
