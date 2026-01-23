import type { Search, SearchResult } from '@shared/types'
import { onMounted, reactive, ref } from 'vue'
import { debounce } from 'lodash'
import axios from 'axios'
import SuggestionControls from '@/modules/SuggestionControls'

export default () => {
    const input = ref<HTMLInputElement | null>(null)
    const search = reactive<Search>({
        query: '',
        results: [],
        loading: false,
        suggestionsIndex: -1,
    })

    onMounted(() => {
        window.addEventListener('keydown', (e: KeyboardEvent) => {
            if (e.key === 'Escape') {
                search.results = []
                search.query = ''
            }
        })

        window.addEventListener('keyup', (e: KeyboardEvent) => {
            if (search.query !== '') {
                new SuggestionControls(search).handle(e.key)
            }
        })
    })

    const makeRequestToFetchItems = debounce<(q: string) => void>(q => {
        if (q === '') {
            search.results = []
            search.loading = false
            return
        }

        axios
            .get<SearchResult[]>(`/api/search/${q}`)
            .then(resp => {
                const results = resp.data

                if (results[0]) {
                    results[0].selected = true
                }

                search.results = results
                search.suggestionsIndex = 0
            })
            .catch(err => console.error(err))
            .finally(() => (search.loading = false))
    }, 100)

    function fetchSearchSuggestions(e: KeyboardEvent): void {
        if (['ArrowDown', 'ArrowUp'].includes(e.key)) {
            return
        }

        search.loading = true
        search.query = search.query.trimStart()

        makeRequestToFetchItems(search.query)
    }

    function changedInput(e: Event): void {
        search.query = (e.target as HTMLInputElement).value
        fetchSearchSuggestions(e as KeyboardEvent)
    }

    return {
        input,
        search,
        changedInput,
    }
}
