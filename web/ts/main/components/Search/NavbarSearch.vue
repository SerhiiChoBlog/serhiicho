<script setup lang="ts">
import { computed, ref } from 'vue'
import useSearch from '@/composables/useSearch'
import isTouchDevice from '@shared/modules/isTouchDevice'
import SearchIcon from '@shared/components/Icons/SearchIcon.vue'
import Dropdown from '@/components/Search/Dropdown.vue'
import AppearTransition from '@shared/components/Transitions/AppearTransition.vue'
import DropdownTransition from '@shared/components/Transitions/DropdownTransition.vue'

const { search, input, changedInput } = useSearch()
const isSmall = ref<boolean>(true)

const hideInput = computed<boolean>(() => {
    return isSmall.value && search.query === ''
})

const openSearchShortcutLabel =
    navigator.userAgent.indexOf('Mac OS X') != -1 ? '⌘+k' : 'ctrl+k'

document.addEventListener('keydown', (e: KeyboardEvent) => {
    if (e.key === 'Escape') {
        closeSearch()
    }

    // Toggle search when pressing cmd + k or ctrl + k
    if (e.key === 'k' && (e.metaKey || e.ctrlKey)) {
        e.preventDefault()
        toggleSearch()
    }
})

function closeSearch(): void {
    setTimeout(() => {
        search.query = ''
        isSmall.value = true
    }, 200)
}

function toggleSearch(): void {
    isSmall.value = !isSmall.value

    if (input.value) {
        input.value.focus()
    }
}
</script>

<template>
    <teleport to="body">
        <div
            v-if="!hideInput"
            @click="closeSearch"
            class="overlay bg-black/60 inset-0 fixed z-40"
        ></div>
    </teleport>

    <appear-transition>
        <div
            :class="hideInput ? 'w-[40px]' : 'w-[300px] md:w-[380px]'"
            class="absolute right-0 top-0 -translate-y-1/2 transition-all duration-300 ease-in-out z-50"
            title="Press cmd + k or ctrl + k to toggle search"
        >
            <label
                for="navbar-search-input"
                class="relative flex items-center justify-center rounded-full mx-auto transition-colors duration-500"
                :class="hideInput ? '' : 'bg-gray-900'"
            >
                <form @submit.prevent class="w-full m-0">
                    <input
                        ref="input"
                        @click="toggleSearch"
                        :type="hideInput ? 'button' : 'search'"
                        placeholder="Search article..."
                        v-model="search.query"
                        @input="changedInput"
                        autocomplete="off"
                        class="pl-7 w-full text-gray-400 rounded-full outline-hidden py-1 text-lg bg-transparent"
                        :class="hideInput ? 'cursor-pointer' : 'pr-12'"
                        id="navbar-search-input"
                    />

                    <div
                        class="inline-flex items-center gap-3 absolute top-1/2 right-3 -translate-y-1/2 cursor-pointer"
                    >
                        <div
                            v-if="!hideInput && !isTouchDevice()"
                            class="bg-gray-700 px-2 rounded-lg"
                        >
                            {{ openSearchShortcutLabel }}
                        </div>

                        <search-icon
                            class="w-5 h-5"
                            :class="{
                                'animate-blink': search.loading,
                                'text-white right-3': hideInput,
                                'text-gray-400 right-4': !hideInput,
                            }"
                        />
                    </div>
                </form>

                <dropdown-transition>
                    <dropdown v-if="search.query !== ''" :search />
                </dropdown-transition>
            </label>
        </div>
    </appear-transition>
</template>
