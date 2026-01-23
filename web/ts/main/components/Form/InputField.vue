<script setup lang="ts">
import ShowPasswordButton from '@/components/ShowPasswordButton.vue'
import ValidationMessage from '@/components/Form/ValidationMessage.vue'

const modelValue = defineModel<string>()

type Props = {
    name: string
    id: string
    errorMessage?: string | null
    type?: 'text' | 'email' | 'password' | 'checkbox'
    placeholder?: string
}

withDefaults(defineProps<Props>(), {
    type: 'text',
    placeholder: '',
})
</script>

<template>
    <div class="relative">
        <input
            :type="type"
            :placeholder="placeholder"
            class="w-full px-4 py-3 border bg-page-second rounded-md focus:outline-hidden focus:ring-1 focus:ring-main"
            :class="{
                'border-red-600/50': errorMessage,
                'border-border': !errorMessage,
            }"
            :name="name"
            v-model="modelValue"
            :id="id"
            :ref="`ref-${name}`"
        />

        <show-password-button v-if="type === 'password'" :input-id="id" />
    </div>

    <validation-message :message="errorMessage" />
</template>
