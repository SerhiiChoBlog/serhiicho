<script setup lang="ts">
import type { User } from '@shared/types/models'
import SignedInMethodLogo from '@/components/SignedInMethodLogo.vue'

type Props = {
    user: User | null
}

defineProps<Props>()
</script>

<template>
    <div class="relative w-[38px] lg:w-[43px] mt-2">
        <img
            :src="`/storage/avatars/${user ? user.avatar : 'default.svg'}`"
            :alt="`${user ? user.name : 'Default'} avatar`"
            class="w-8 lg:w-10 aspect-square rounded-full shadow-md"
            width="400"
            height="400"
        />

        <signed-in-method-logo
            v-if="user"
            :method="user.sign_in_method"
            class="absolute w-4 lg:w-5 aspect-square left-6 lg:left-7 top-6 lg:top-7"
        />
    </div>
</template>
