<script setup lang="ts">
import axios from 'axios'
import Tip from '@shared/components/Tip.vue'
import getSocialLinks from '@/modules/getSocialLinks'

interface Props {
    postTitle: string
    postId: number
}

const { postTitle, postId } = defineProps<Props>()
const links = getSocialLinks(postTitle)

function saveShare(social: string): void {
    const params = {
        social,
        secret: Auth ? Auth.secret : null,
    }

    axios
        .post(`/api/social-shares/${postId}`, params)
        .catch(err => console.error(err))
}
</script>

<template>
    <div class="flex gap-x-2 lg:gap-x-3 print:hidden">
        <a
            v-for="link in links"
            :key="link.title"
            :href="link.href ? link.href : 'javascript:'"
            :target="link.href ? '_blank' : ''"
            rel="noopener noreferrer"
            @mousedown="saveShare(link.title)"
            @click="link.action ? link.action() : null"
            :aria-label="link.tip"
        >
            <tip :content="link.tip" class="transition-opacity hover:opacity-80">
                <img
                    :src="`/storage/icons/${link.title}.webp`"
                    :alt="`Logo of the ${link.title}`"
                    width="27"
                    height="27"
                />
            </tip>
        </a>
    </div>
</template>
