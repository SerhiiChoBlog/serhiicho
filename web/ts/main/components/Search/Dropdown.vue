<script setup lang="ts">
import { Search } from '@shared/types'
import NotFound from './NotFound.vue'
import DropdownItems from './DropdownItems.vue'
import ReloadIcon from '@shared/components/Icons/ReloadIcon.vue'

defineProps<{ search: Search }>()
</script>

<template>
    <div
        :class="[
            'bg-page dark:bg-page-second shadow-xl top-[calc(100%+9px)] absolute',
            'left-0 right-0 z-50 rounded-lg dark:border-2 dark:border-border',
            'min-h-12.5 max-h-150 overflow-y-auto',
        ]"
    >
        <div>
            <div v-if="search.loading" class="py-8 flex items-center justify-center">
                <reload-icon class="animate-spin size-6" />
            </div>

            <div
                v-else-if="!search.loading && search.results.length === 0"
                class="py-8 flex items-center justify-center"
            >
                <not-found />
            </div>

            <div v-else>
                <dropdown-items
                    v-if="search.results.length > 0"
                    :items="search.results"
                />
            </div>
        </div>
    </div>
</template>
