<script setup lang="ts">
import useSearch from '@/composables/useSearch'
import SearchIcon from '@shared/components/Icons/SearchIcon.vue'
import Dropdown from '@/components/Search/Dropdown.vue'
import AppearTransition from '@shared/components/Transitions/AppearTransition.vue'
import DropdownTransition from '@shared/components/Transitions/DropdownTransition.vue'

const { search, input, changedInput } = useSearch()
</script>

<template>
    <appear-transition>
        <div class="max-w-[85%] md:max-w-[600px] mx-auto">
            <label
                for="search-input"
                class="relative flex items-center justify-center mt-5 shadow-lg rounded-full bg-page mx-auto dark:border dark:border-gray-600"
            >
                <form @submit.prevent class="w-full m-0">
                    <search-icon
                        :class="{ 'animate-blink': search.loading }"
                        class="w-4 md:w-5 h-4 md:h-6 text-gray-400 absolute top-1/2 -translate-y-1/2 left-4 lg:left-5"
                    />

                    <input
                        ref="input"
                        type="search"
                        placeholder="Search blog posts"
                        @input="changedInput"
                        v-model="search.query"
                        @blur="search.query = ''"
                        autocomplete="off"
                        class="pl-12 pr-7 md:pl-14 py-3 text-sm md:text-xl w-full text-font rounded-full outline-hidden bg-page dark:bg-page-second"
                        id="search-input"
                    />
                </form>

                <dropdown-transition>
                    <dropdown v-if="search.query !== ''" :search />
                </dropdown-transition>
            </label>
        </div>
    </appear-transition>
</template>
