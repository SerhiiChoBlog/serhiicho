<script setup lang="ts">
import type { Testimonial } from '@shared/types'

const emit = defineEmits<{ (e: 'showFullContent'): void }>()

defineProps<{
    testimonial: Testimonial
    showFull: boolean
}>()
</script>

<template>
    <figure class="relative flex flex-col rounded-lg p-4 md:p-6 shadow-lg bg-page">
        <figcaption class="flex items-center space-x-4">
            <img
                :src="testimonial.image_uri"
                :alt="testimonial.name"
                class="h-14 w-14 rounded-full shadow-md"
            />

            <div>
                <h3 class="font-bold text-xl text-gray-400">
                    <a
                        :href="testimonial.link ? testimonial.link : 'javascript:'"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        {{ testimonial.name }}
                    </a>
                </h3>

                <h3 class="text-[.8em] leading-6 text-gray-300">
                    {{ testimonial.role }}
                </h3>
            </div>
        </figcaption>

        <blockquote class="mt-3">
            <p
                v-html="
                    !showFull && testimonial.short
                        ? testimonial.short
                        : testimonial.content
                "
                class="[&>p]:mb-5 [&>p:last-child]:mb-0 text-gray-200 leading-6 text-[.9em]"
            ></p>

            <button
                v-if="!showFull && testimonial.short"
                type="button"
                @click="emit('showFullContent')"
                class="text-sm font-bold text-main hover:text-main-dark transition-colors block mt-4 uppercase"
            >
                Show full
            </button>
        </blockquote>
    </figure>
</template>
