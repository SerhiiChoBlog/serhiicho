import type { BookAuthor } from '@shared/types/models'
import type { BookAuthorSuggestionsEmits } from '@shared/types/emits'
import { ref } from 'vue'
import axios from 'axios'
import { debounce } from 'lodash'
import handleError from '@admin/modules/handleError'

export default (emit: BookAuthorSuggestionsEmits) => {
    const searchQuery = ref('')
    const loading = ref<boolean>(false)
    const accepted = ref<string | null>(null)
    const suggestions = ref<BookAuthor[]>([])

    function addAuthor(): void {
        if (searchQuery.value === '') {
            return
        }

        loading.value = true

        axios.post<BookAuthor>(`/admin/ajax/book-authors`, {
            name: searchQuery.value,
        })
            .then(resp => chooseSuggestion(resp.data))
            .catch(handleError)
            .finally(() => loading.value = false)
    }

    const fetchSuggestions = debounce(() => {
        if (searchQuery.value === '') {
            suggestions.value = []
            return
        }

        loading.value = true

        axios.get<BookAuthor[]>(`/admin/ajax/book-authors/search/${searchQuery.value}`)
            .then(resp => suggestions.value = resp.data || [])
            .catch(handleError)
            .finally(() => loading.value = false)
    }, 200)

    function inputChanged(val: string): void {
        searchQuery.value = val
        fetchSuggestions()
    }

    function chooseSuggestion(suggestion: BookAuthor): void {
        accepted.value = suggestion.name
        searchQuery.value = suggestion.name
        suggestions.value = []
        emit('changed', suggestion.id)
    }

    return {
        searchQuery,
        loading,
        inputChanged,
        addAuthor,
        accepted,
        suggestions,
        chooseSuggestion,
    }
}
