<script setup lang="ts">
import { MenuItems } from '@headlessui/vue'
import DropdownTransition from '@shared/components/Transitions/DropdownTransition.vue'
</script>

<template>
    <dropdown-transition>
        <menu-items
            class="absolute right-0 mt-2 w-56 origin-top-right divide-y divide-border rounded-md bg-white shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-hidden"
        >
            <slot />
        </menu-items>
    </dropdown-transition>
</template>
