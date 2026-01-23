<script setup lang="ts">
import type { Tag } from '@shared/types/models'
import { onMounted, ref, watch } from 'vue'
import axios from 'axios'
import { Listbox } from '@headlessui/vue'
import ListboxButton from '@/components/Listbox/ListboxButton.vue'
import ListboxOptions from '@/components/Listbox/ListboxOptions.vue'
import ListboxOption from '@/components/Listbox/ListboxOption.vue'
import replaceQueryParam from '@/modules/replaceQueryParam'
import showEverythingTag from '@/modules/showEverythingTag'
import CloseIcon from '@shared/components/Icons/CloseIcon.vue'
import TagImage from './TagImage.vue'

const { tag } = defineProps<{ tag: Tag }>()
const tags = ref<Tag[]>([showEverythingTag])
const selectedTag = ref<Tag>(tag)

watch(selectedTag, () => {
    filterPosts(selectedTag.value)
})

onMounted(() => fetchTags())

function filterPosts(tag: Tag): void {
    if (tag.name === 'everything') {
        clearFilters()
        return
    }

    const query = replaceQueryParam('tag', tag.name)

    if (window.location.search != query) {
        window.location.href = `/posts${query}`
    }
}

function clearFilters(): void {
    window.location.href = `/posts`
}

function fetchTags(): void {
    axios
        .get<Tag[]>('/api/tags')
        .then(resp => {
            tags.value.push(...resp.data)
        })
        .catch(err => console.error(err))
}
</script>

<template>
    <div
        class="grid"
        :class="
            tag.name === showEverythingTag.name
                ? 'grid-cols-1 mr-3'
                : 'grid-cols-[280px_35px] gap-1'
        "
    >
        <listbox v-model="selectedTag">
            <div class="relative">
                <listbox-button>
                    <div class="flex items-center gap-2">
                        <tag-image :tagName="tag.name" />
                        {{ selectedTag.title }}
                    </div>
                </listbox-button>

                <listbox-options>
                    <listbox-option v-for="tag in tags" :key="tag.name" :value="tag">
                        <div class="flex items-center gap-2">
                            <tag-image :tagName="tag.name" />
                            {{ tag.title }}
                        </div>
                    </listbox-option>
                </listbox-options>
            </div>
        </listbox>

        <button
            @click="clearFilters"
            type="button"
            class="p-2 rounded-lg transition-colors hover:bg-page-second hover:text-red-700"
            v-if="tag.name !== showEverythingTag.name"
        >
            <close-icon class="w-5 h-5 text-font-second" />
        </button>
    </div>
</template>
