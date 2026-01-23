<script setup lang="ts">
import type { Feed } from '@shared/types/models'
import type { Component } from 'vue'

defineProps<{
    title: string
    content: string
    image: string
    feed: Feed
    url?: string
    icon?: Component
}>()
</script>

<template>
    <div
        class="flex items-center gap-5 bg-page-second rounded-lg px-6 py-2 md:py-4 space-y-1.5"
    >
        <div v-if="image" class="flex items-center relative min-w-[60px]">
            <a :href="url ? url : 'javascript:'">
                <img
                    :src="image"
                    class="w-[60px] h-[60px] rounded-md object-cover object-right"
                    alt="Feed image"
                    width="60"
                    height="60"
                />
            </a>

            <div
                v-if="icon"
                class="absolute -left-5 -top-3 rounded-full bg-page-second p-0.5"
            >
                <component :is="icon" class="w-8 h-8 fill-page-second" />
            </div>
        </div>

        <div>
            <span class="text-md opacity-70">
                <strong v-if="feed.user">{{ feed.user.name }}</strong>
                {{ title }} · {{ feed.pretty_date }}
            </span>

            <h2 class="text-md md:text-lg opacity-80 mt-1">{{ content }}</h2>
        </div>
    </div>
</template>
