<script setup lang="ts">
import { ref } from 'vue'
import handleServerError from '@shared/modules/handleServerError'
import showToast from '@shared/modules/showToast'
import axios from 'axios'
import PageRefresher from '@/modules/PageRefresher'

const loading = ref<boolean>(false)

function setPhoto(event: Event) {
    const input = event.target as HTMLInputElement
    const photo = input.files?.[0]

    if (!photo) {
        showToast({
            text: 'Please select a photo.',
            success: false,
        })

        return
    }

    uploadPhoto(photo)
}

function uploadPhoto(photo: File) {
    if (loading.value) {
        return
    }

    loading.value = true

    const formData = new FormData()
    formData.append('photo', photo)
    formData.append('secret', Auth.secret)

    axios
        .post('/api/users/upload-photo', formData)
        .then(() => PageRefresher.refresh('Photo uploaded successfully'))
        .catch(e => handleServerError(e))
        .finally(() => (loading.value = false))
}
</script>

<template>
    <label
        class="absolute w-full hidden group-hover:flex inset-0 bg-gray-900/80 items-center justify-center rounded-full text-white border-4 border-dashed border-gray-600 cursor-pointer"
    >
        <span class="text-lg select-none">Select photo</span>
        <input type="file" class="hidden" @change="setPhoto" />
    </label>
</template>
