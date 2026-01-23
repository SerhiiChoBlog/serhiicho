<script setup lang="ts">
import { onMounted, ref } from 'vue'
import Tip from '@shared/components/Tip.vue'
import EyeIcon from '@shared/components/Icons/EyeIcon.vue'
import EyeSlashIcon from '@shared/components/Icons/EyeSlashIcon.vue'

const { inputId } = defineProps<{ inputId: string }>()
const input = ref<HTMLInputElement | null>(null)
const showPassword = ref(false)

onMounted(() => {
    const elem = document.getElementById(inputId) as HTMLInputElement | null

    if (elem) {
        input.value = elem
    }
})

function changeInputType(): void {
    if (!input.value) {
        return
    }

    input.value.type = input.value.type === 'text' ? 'password' : 'text'
    showPassword.value = input.value.type === 'text'
}
</script>

<template>
    <tip content="Show / hide password">
        <button
            class="absolute right-2 top-1/2 -translate-y-1/2 px-1 text-gray-500 transition-colors hover:text-black"
            @click="changeInputType"
            type="button"
        >
            <eye-slash-icon v-if="showPassword" />
            <eye-icon v-else />
        </button>
    </tip>
</template>
