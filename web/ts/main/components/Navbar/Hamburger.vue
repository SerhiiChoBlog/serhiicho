<script setup lang="ts">
import { ref } from 'vue'
import { events } from '@shared/appConfig'
import dispatchEvent from '@shared/modules/dispatchEvent'
import listenEvent from '@shared/modules/listenEvent'

const isActive = ref(false)

listenEvent<boolean>(events.openMobileNavbar, isOpened => {
    isActive.value = isOpened
})

function toggleNavbar(): void {
    dispatchEvent(events.openMobileNavbar, !isActive.value)
}
</script>

<template>
    <a
        role="button"
        class="block lg:hidden relative size-13 transition-transform duration-300 delay-300 z-20 cursor-pointer"
        :class="isActive ? 'rotate-45' : ''"
        @click="toggleNavbar"
        aria-label="Show mobile menu"
    >
        <span
            aria-hidden="true"
            class="w-[19px] h-[2px] block bg-white rounded-full absolute left-[calc(50%-8px)] transition-all duration-500"
            :class="
                isActive
                    ? 'rotate-[-90deg] top-[calc(50%-1.5px)]'
                    : 'top-[calc(50%-6px)]'
            "
        ></span>

        <span
            aria-hidden="true"
            class="w-[19px] h-[2px] block bg-white rounded-full absolute top-[calc(50%-1px)] left-[calc(50%-8px)] transition-all duration-500"
        ></span>

        <span
            aria-hidden="true"
            class="w-[19px] h-[2px] block bg-white rounded-full absolute top-[calc(50%+4px)] left-[calc(50%-8px)] transition-all duration-500"
            :class="isActive ? 'translate-y-[40px] opacity-0' : ''"
        ></span>
    </a>
</template>
