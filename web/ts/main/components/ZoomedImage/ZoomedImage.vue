<script setup lang="ts">
import { onMounted, ref } from 'vue'
import ScaleTransition from '@shared/components/Transitions/ScaleTransition.vue'

const images = ref<HTMLImageElement[]>([])
const selectedImage = ref<HTMLImageElement | null>(null)

onMounted(() => {
    const elems = document.querySelectorAll<HTMLImageElement>('#post-content img')
    images.value = Array.from(elems)
    addClickEventsToImages()
})

function addClickEventsToImages(): void {
    images.value.forEach(img => {
        img.classList.add('cursor-zoom-in')

        img.addEventListener('click', () => {
            selectedImage.value = img
        })
    })
}

function hideZoomedImage(): void {
    selectedImage.value = null
}
</script>

<template>
    <scale-transition>
        <div
            v-if="selectedImage"
            class="fixed inset-0 z-[100] bg-black/80 flex items-center justify-center px-0 md:px-3"
            @click="hideZoomedImage"
        >
            <img
                :src="selectedImage.src"
                class="w-auto max-h-screen cursor-zoom-out shadow-lg shadow-black/30 rounded-xl"
                :alt="selectedImage.alt"
            />
        </div>
    </scale-transition>
</template>
