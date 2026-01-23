<script setup lang="ts">
import type { Theme } from '@shared/types'
import { computed, ref, onMounted } from 'vue'
import axios from 'axios'
import Tip from '@shared/components/Tip.vue'

interface Props {
    theme: Theme
}

const { theme } = defineProps<Props>()
const currentTheme = ref<Theme>(theme)

const title = computed<string>(() => {
    let switchTo = 'light'

    if (currentTheme.value === 'light') {
        switchTo = 'dark'
    }

    return `Switch to ${switchTo} theme`
})

onMounted(() => setInitialTheme())

function setInitialTheme(): void {
    if (currentTheme.value !== 'system') {
        return
    }

    hasDarkSystemTheme() ? switchTheme('dark') : switchTheme('light')
}

function switchTheme(toTheme: 'dark' | 'light'): void {
    axios
        .get(`/theme/${toTheme}`)
        .then(() => changeInterfaceTo(toTheme))
        .catch(err => console.error(err))
}

function hasDarkSystemTheme(): boolean {
    return window.matchMedia('(prefers-color-scheme: dark)').matches
}

function changeInterfaceTo(toTheme: 'light' | 'dark'): void {
    currentTheme.value = toTheme
    document.body.classList.add(toTheme)
    document.body.classList.remove(toTheme === 'dark' ? 'light' : 'dark')
}
</script>

<template>
    <tip :content="title">
        <div class="flex items-center justify-center">
            <button
                type="button"
                :class="[
                    'cursor-pointer drop-shadow-md relative text-yellow-300',
                    'dark:text-main flex py-2 px-4 lg:px-2 bg-transparent',
                    'hover:bg-gray-800 hover:bg-opacity-50 rounded-full transition-colors'
                ]"
                @click="() => switchTheme(currentTheme === 'dark' ? 'light' : 'dark')"
                aria-label="Toggle between light and dark themes"
            >
                <svg
                    xmlns="http://www.w3.org/2000/svg"
                    xmlns:xlink="http://www.w3.org/1999/xlink"
                    x="0px"
                    y="0px"
                    width="27px"
                    height="27px"
                    viewBox="0 0 212.1 212.1"
                    xml:space="preserve"
                    fill="currentColor"
                >
                    <!-- Sun -->
                    <g class="origin-center transition-transform duration-500">
                        <path
                            d="M181.1,75V31.1h-43.9L106.1,0L75,31.1H31.1V75L0,106.1l31.1,31.1v43.9H75l31.1,31.1l31.1-31.1h43.9v-43.9l31.1-31.1
                            L181.1,75z M106.1,162.1c-30.9,0-56-25.1-56-56s25.1-56,56-56s56,25.1,56,56S137,162.1,106.1,162.1z"
                        />
                    </g>

                    <!-- Circle -->
                    <g
                        class="origin-center transition-transform translate-x-0 dark:-translate-x-[15%]"
                    >
                        <circle cx="106.1" cy="106.1" r="41" />
                    </g>
                </svg>

                <span class="block lg:hidden pl-4">
                    {{ title }}
                </span>
            </button>
        </div>
    </tip>
</template>
