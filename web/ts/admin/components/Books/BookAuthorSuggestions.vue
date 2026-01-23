<script setup lang="ts">
import type { BookAuthorSuggestionsEmits } from '@shared/types/emits'
import FormInput from '@admin/components/Form/FormInput.vue'
import useBookAuthorSuggestions from '@admin/composables/books/useBookAuthorSuggestions'

const emit = defineEmits<BookAuthorSuggestionsEmits>()

const {
    searchQuery,
    loading,
    inputChanged,
    addAuthor,
    accepted,
    suggestions,
    chooseSuggestion,
} = useBookAuthorSuggestions(emit)
</script>

<template>
    <div class="relative">
        <FormInput
            type="text"
            id="book-authors"
            name="book-authors"
            placeholder="Start typing to add an author..."
            @changed="inputChanged"
            @keydown.enter.prevent="addAuthor"
            :value="searchQuery"
        >
            <span v-if="accepted">Book author is {{ accepted }}</span>
            <span v-else>Enter a book's author</span>

            <i
                v-if="loading"
                class="fas fa-spinner fa-spin text-orange-600 ml-2"
            ></i>
            <i v-else-if="accepted" class="fas fa-check text-green-600 ml-2"></i>
            <i v-else class="fas fa-arrow-left text-green-600 ml-2"></i>
        </FormInput>

        <div class="mt-2">
            <div v-if="suggestions.length === 0">
                <span>No suggestions yet</span>
            </div>
            <div v-else>
                <span>Suggestions:</span>

                <div class="inline-flex gap-1.5 ml-2">
                    <button
                        v-for="suggestion in suggestions"
                        :key="suggestion.id"
                        @click="chooseSuggestion(suggestion)"
                        class="text-main hover:underline"
                        type="button"
                    >
                        {{ suggestion.name }}
                    </button>
                </div>
            </div>
        </div>
    </div>
</template>
