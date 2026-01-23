<script setup lang="ts">
import type { Badge } from '@shared/types/models'
import Tip from '@shared/components/Tip.vue'

type Props = {
    badge: Badge
}

const { badge } = defineProps<Props>()
</script>

<template>
    <Tip :content="badge.description">
        <img
            :src="`/storage/badges/${badge.icon}`"
            :alt="`Badge for ${badge.name}`"
            width="270"
            height="270"
            class="drop-shadow-md pointer-events-none"
        />
    </Tip>
</template>
