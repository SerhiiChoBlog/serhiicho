<script setup lang="ts">
import type { RouteLocationRaw } from 'vue-router'

type Props = {
    icon: string
    title: string
    href?: string
    route?: RouteLocationRaw
}

const { icon, href, title, route } = defineProps<Props>()
const classes =
    'text-sm border border-border rounded-xs w-10 h-10 flex justify-center items-center transition-colors hover:border-main hover:text-main'
</script>

<template>
    <a v-if="href" :href :class="classes" :title>
        <i class="fas" :class="icon"></i>
    </a>

    <RouterLink v-if="route" :to="route" :class="classes" :title>
        <i class="fas" :class="icon"></i>
    </RouterLink>
</template>
