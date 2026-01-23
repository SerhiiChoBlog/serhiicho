<script setup lang="ts">
import { User } from '@shared/types/models'
import SignedInMethodLogo from '@/components/SignedInMethodLogo.vue'
import AppearTransition from '@shared/components/Transitions/AppearTransition.vue'
import UploadProfilePhoto from '@user/components/Profile/UploadProfilePhoto.vue'

defineProps<{ user: User }>()
</script>

<template>
    <appear-transition>
        <div class="relative max-w-[250px] mx-auto lg:mx-0">
            <div class="relative group">
                <img
                    :src="`/storage/avatars/${user.avatar}`"
                    :alt="`${user.name}'s avatar`"
                    width="200"
                    height="200"
                    class="w-full lg:mx-0 rounded-full shadow-lg pointer-events-none"
                />

                <upload-profile-photo v-if="Auth && Auth.id === user.id" />
            </div>

            <signed-in-method-logo
                :src="`/storage/icons/${user.sign_in_method}.webp`"
                :method="user.sign_in_method"
                class="absolute bottom-0 right-5 w-14 h-14"
            />
        </div>
    </appear-transition>
</template>
