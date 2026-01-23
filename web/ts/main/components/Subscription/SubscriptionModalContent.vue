<script setup lang="ts">
import type { SubscriptionTagsModal } from '@shared/types'
import type { Tag } from '@shared/types/models'
import { computed, onMounted, ref } from 'vue'
import CheckIcon from '@shared/components/Icons/CheckIcon.vue'
import CloseIcon from '@shared/components/Icons/CloseIcon.vue'
import SubscriptionButton from '@/modules/Subscription/SubscriptionButton'
import SubscriptionCache from '@/modules/Subscription/SubscriptionCache'

const props = defineProps<{ tags: SubscriptionTagsModal }>()
const selectedTags = ref<Tag[]>([])

const allTagsAreChosen = computed(() => {
    return props.tags.items.length === selectedTags.value.length
})

onMounted(() => {
    const cachedTags = SubscriptionCache.get()
    selectedTags.value = cachedTags ? cachedTags : props.tags.items
})

function toggleAll(): void {
    allTagsAreChosen.value ? removeAll() : selectAll()
}

function selectAll(): void {
    selectedTags.value = props.tags.items

    SubscriptionCache.set(props.tags.items)
    SubscriptionButton.setAllTags()
}

function removeAll(): void {
    selectedTags.value = []

    SubscriptionCache.set([])
    SubscriptionButton.setNoTags()
}

function chooseSubject(tag: Tag): void {
    if (selectedTagsHaveTag(tag)) {
        selectedTags.value = selectedTags.value.filter(subject => {
            return subject.id !== tag.id
        })
    } else {
        selectedTags.value.push(tag)
    }

    SubscriptionCache.set(selectedTags.value)
    SubscriptionButton.setTagsAmount(selectedTags.value.length)
}

function selectedTagsHaveTag(tag: Tag): boolean {
    return !!selectedTags.value.find(subject => subject.id === tag.id)
}
</script>

<template>
    <div>
        <div class="mb-5 text-center">
            <h2 class="text-2xl">Choose tags to subscribe</h2>

            <p class="text-xl">
                Tags chosen:
                <b class="text-green-700 font-sans">
                    {{
                        allTagsAreChosen
                            ? 'everything'
                            : `${selectedTags.length} out of ${tags.items.length}`
                    }}
                </b>
            </p>

            <div>
                <small
                    >You can choose only tags that I'm passionate to write
                    about</small
                >
            </div>

            <button
                type="button"
                class="px-7 py-0.5 border border-gray-400 rounded-full mt-3"
                @click="toggleAll"
            >
                <div v-if="allTagsAreChosen" class="flex items-center gap-2">
                    <close-icon class="w-5 h-5 text-red-800" />
                    Remove all
                </div>
                <div v-else class="flex items-center gap-2">
                    <check-icon class="w-5 h-5 text-green-700" />
                    Select all
                </div>
            </button>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
            <button
                type="button"
                v-for="tag in tags.items"
                :key="tag.id"
                class="flex items-center gap-7 border rounded-lg py-2 px-3 relative pointer outline-2 outline-transparent transition-all hover:outline-gray-400 dark:hover:outline-gray-600"
                :class="
                    selectedTagsHaveTag(tag)
                        ? 'border-green-500 dark:border-green-900 bg-green-50 dark:bg-page-second'
                        : 'border-border'
                "
                aria-label="Subscription subject"
                @click="chooseSubject(tag)"
            >
                <check-icon
                    v-if="selectedTagsHaveTag(tag)"
                    class="w-6 h-6 absolute top-1 right-1 text-green-700"
                />

                <img
                    :src="`/storage/tags/${tag.name}.webp`"
                    :alt="`The logo of ${tag.name}`"
                    width="200"
                    height="200"
                    class="w-[50px] rounded-md drop-shadow-md"
                />

                <div>
                    <h3 class="font-semibold font-sans text-lg">{{ tag.title }}</h3>
                    <p class="text-font-second text-md">{{ tag.description }}</p>
                </div>
            </button>
        </div>
    </div>
</template>
