<script setup lang="ts">
import { Tag } from '@shared/types/models'
import { ref, watch } from 'vue'
import axios from 'axios'
import handleError from '@admin/modules/handleError'

const emit = defineEmits<{
    (e: 'addTag', tag: Tag): void
}>()

const props = defineProps<{ value: string }>()
const suggestions = ref<Tag[]>([])

watch(props, () => {
    fetchTags(props.value)
})

function fetchTags(input: string): void {
    if (input === '') {
        suggestions.value = []
        return
    }

    axios
        .post<Tag[]>(`/api/tags/search/${input}`)
        .then(resp => (suggestions.value = resp.data))
        .catch(handleError)
}

function chooseTag(tag: Tag): void {
    emit('addTag', tag)
    suggestions.value = []
}
</script>

<template>
    <ul
        v-if="suggestions.length > 0"
        class="bg-white shadow-md mt-2 rounded-md overflow-hidden divide-y divide-border"
    >
        <li
            v-for="tag in suggestions"
            :key="tag.id"
            class="transition-colors hover:bg-page-second px-3 py-1"
            @click="chooseTag(tag)"
        >
            {{ tag.name }}
        </li>
    </ul>
</template>
