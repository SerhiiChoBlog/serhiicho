<script setup lang="ts">
import { ref } from 'vue'

type Props = {
    uri: string
    defaultUri: string
}

const { uri, defaultUri } = defineProps<Props>()
const imageSrc = ref<string>(uri ? `/${uri}` : defaultUri)

function updateImagePreview(e: Event): void {
    const fileReader = new FileReader()
    const src = e.target as HTMLInputElement

    if (src.files && src.files.length !== 0) {
        fileReader.readAsDataURL(src.files[0])

        fileReader.onload = function () {
            imageSrc.value = this.result as string
        }

        return
    }

    imageSrc.value = defaultUri
}
</script>

<template>
    <div class="relative">
        <img :src="imageSrc" alt="image" class="rounded-md shadow-lg w-full" />

        <label
            for="post-image-input"
            class="opacity-0 flex items-center hover:opacity-100 transition-opacity border-4 lg:border-4 border-dashed text-center rounded-lg cursor-pointer border-gray-300 absolute inset-2 lg:inset-6"
        >
            <b
                class="text-gray-600 px-4 py-2 rounded-xl bg-gray-300 text-xl lg:text-2xl drop-shadow-xs mx-auto shadow-lg opacity-95"
            >
                Choose image
            </b>
        </label>

        <input
            @change="updateImagePreview"
            type="file"
            id="post-image-input"
            name="image"
            class="hidden"
        />
    </div>
</template>
