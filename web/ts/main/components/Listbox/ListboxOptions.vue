<script setup lang="ts">
import { ListboxOptions } from '@headlessui/vue'
</script>

<template>
    <transition
        leave-active-class="transition duration-100 ease-in"
        leave-from-class="opacity-100"
        leave-to-class="opacity-0"
    >
        <listbox-options
            class="absolute mt-1 z-10 max-h-96 w-full overflow-auto rounded-md bg-page py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:hidden-none sm:text-sm"
        >
            <slot />
        </listbox-options>
    </transition>
</template>
