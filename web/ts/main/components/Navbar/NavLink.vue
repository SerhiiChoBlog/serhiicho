<script setup lang="ts">
import trimSlashes from '@shared/modules/trimSlashes'
import { computed } from 'vue'

type Props = {
    href: string
    classes?: string
}

const { href, classes } = defineProps<Props>()
const isSelfHref = computed(() => trimSlashes(href) === trimSlashes(window.location.pathname))
const className = `relative tracking-wider w-full min-w-fit lg:w-max border-b
    border-border lg:border-none text-sm before:content-none px-7 py-5 lg:p-0
    lg:before:content-[''] navbar-link uppercase lg:drop-shadow-md ${classes ?? ''}`
</script>

<template>
    <small v-if="isSelfHref" :class="className">
        <slot />
    </small>

    <a v-else :href="href" :class="className" rel="noreferrer">
        <slot />
    </a>
</template>
