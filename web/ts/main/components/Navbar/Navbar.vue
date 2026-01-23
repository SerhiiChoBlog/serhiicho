<script setup lang="ts">
import Hamburger from '@/components/Navbar/Hamburger.vue'
import NavbarMenu from '@/components/Navbar/NavbarMenu.vue'
import { Theme } from '@shared/types'
import AppearTransition from '@shared/components/Transitions/AppearTransition.vue'
import useScroll from '@/composables/useScroll'
import { computed } from 'vue'
import NavbarSearch from '@/components/Search/NavbarSearch.vue'

const { theme } = defineProps<{ theme: Theme }>()
const { isScrolledDown } = useScroll()
const isHome = computed<boolean>(() => window.location.pathname === '/')
</script>

<template>
    <appear-transition>
        <nav
            class="fixed top-0 left-0 right-0 z-50 transition-all duration-300 print:hidden"
            :class="{
                'bg-gray-800/95 shadow-md': isScrolledDown,
                'bg-transparent': !isScrolledDown,
            }"
        >
            <div
                class="container flex justify-between items-center min-h-[3.4rem] gap-4"
            >
                <div class="flex items-center justify-between gap-x-1 w-full">
                    <div
                        v-if="isHome"
                        class="print:hidden text-lg transition-opacity hover:opacity-80 text-lato drop-shadow-[1px_1px_0_rgba(0,0,0,.4)]"
                    >
                        <span class="tracking-wider text-white">{S} SERHII</span>
                    </div>

                    <a
                        v-else
                        tabindex="0"
                        href="/"
                        class="print:hidden text-lg transition-opacity hover:opacity-80 text-lato drop-shadow-[1px_1px_0_rgba(0,0,0,.4)]"
                        aria-label="Visit main page"
                        rel="noreferrer"
                    >
                        <span class="tracking-wider text-white">{S} SERHII</span>
                    </a>

                    <div class="flex items-center">
                        <div class="relative">
                            <navbar-search />
                        </div>

                        <div class="relative">
                            <hamburger />
                        </div>
                    </div>
                </div>

                <navbar-menu :theme />
            </div>
        </nav>
    </appear-transition>
</template>
