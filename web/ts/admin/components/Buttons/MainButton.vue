<script setup lang="ts">
import type { RouteLocationRaw } from 'vue-router'

const { route, type, href } = defineProps<{
    route?: RouteLocationRaw
    type?: 'button' | 'submit' | 'reset'
    href?: string
}>()

const classes =
    'inline-block uppercase text-white text-center text-sm px-8 py-3 rounded-full transition-all bg-main hover:bg-main-hover shadow-md hover:shadow-lg'
</script>

<template>
    <a v-if="href" :href :class="classes" rel="noreferrer">
        <slot />
    </a>

    <RouterLink v-else-if="route" :class="classes" :to="route">
        <slot />
    </RouterLink>

    <button v-else :type :class="classes">
        <slot />
    </button>
</template>
