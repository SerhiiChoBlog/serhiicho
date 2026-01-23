<script setup lang="ts">
import { ListboxButton } from '@headlessui/vue'
import SelectorIcon from '@shared/components/Icons/SelectorIcon.vue'
</script>

<template>
    <listbox-button
        class="relative w-full cursor-default rounded-lg bg-page py-2 pl-3 pr-9 text-left focus:outline-hidden border border-border text-xs sm:text-sm"
    >
        <span class="truncate block">
            <slot />
        </span>

        <span
            class="pointer-events-none absolute inset-y-0 right-0 flex items-center pr-2"
        >
            <selector-icon class="h-5 w-5 text-gray-400" aria-hidden="true" />
        </span>
    </listbox-button>
</template>
