<script setup lang="ts">
import type { User } from '@shared/types/models'
import type { FirebaseSignUpRequest } from '@shared/types'
import { ref } from 'vue'
import getFirebaseUser from '@/firebase/getFirebaseUser'
import firebaseAuth from 'firebase/auth'
import getGoogleProvider from '@/firebase/getGoogleProvider'
import getGithubProvider from '@/firebase/getGithubProvider'
import showToast from '@shared/modules/showToast'
import handleServerError from '@shared/modules/handleServerError'
import axios from 'axios'
import PageRefresher from '@/modules/PageRefresher'
import Spinner from '@shared/components/Icons/Spinner.vue'
import Tip from '@shared/components/Tip.vue'

const { serviceName } = defineProps<{
    serviceName: 'google' | 'github'
}>()

const isLoading = ref<boolean>(false)

function signIn(): void {
    if (isLoading.value) {
        return
    }

    isLoading.value = true

    let provider: firebaseAuth.AuthProvider | null = null

    if (serviceName === 'github') {
        provider = getGithubProvider()
    } else if (serviceName === 'google') {
        provider = getGoogleProvider()
    }

    if (!provider) {
        showToast({
            text: `Something went wrong with ${serviceName} service`,
            success: false,
        })
        return
    }

    getFirebaseUser(makeRequestToSignInUser, handleError, provider, serviceName)
}

function makeRequestToSignInUser(params: FirebaseSignUpRequest): void {
    axios
        .post<User | string>('/firebase-sign-in', params)
        .then(resp => handleResponse(resp.data))
        .catch(e => handleServerError(e))
        .finally(() => (isLoading.value = false))
}

function handleResponse(user: User | string): void {
    if (typeof user === 'string') {
        showToast({ text: user })
        return
    }

    if (!user) {
        return
    }

    const msg = `Signed in successfully as ${user.name}. Click here to go to your profile`
    PageRefresher.refresh(msg, '/users/profile')
}

function handleError(err: any): void {
    console.error(err)
    isLoading.value = false
}
</script>

<template>
    <button
        @click="signIn"
        type="button"
        :disabled="isLoading"
        :class="[
            'cursor-pointer rounded-full w-8 h-8 transition-transform',
            'transform hover:scale-110 flex items-center',
        ]"
    >
        <spinner v-if="isLoading" class="w-8 h-8" />

        <tip v-else :content="`Sign in with ${serviceName}`" class="inline-flex">
            <img
                :src="`/storage/icons/${serviceName}.webp`"
                :alt="`${serviceName} logo`"
                width="100"
                height="100"
                class="w-full"
            />
        </tip>
    </button>
</template>
