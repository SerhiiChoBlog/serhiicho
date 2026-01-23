<script setup lang="ts">
import FirebaseSignInButton from '@/components/FirebaseSignInButton.vue'
</script>

<template>
    <div class="flex justify-center items-center gap-3 pt-7">
        <span>Or login with</span>

        <div class="min-w-[70px] flex gap-4 items-center">
            <firebase-sign-in-button service-name="google" />
            <firebase-sign-in-button service-name="github" />
        </div>
    </div>
</template>
