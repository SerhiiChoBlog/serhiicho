<script setup lang="ts">
import { ref } from 'vue'
import listenEvent from '@shared/modules/listenEvent'
import { events } from '@shared/appConfig'
import Modal from '@shared/components/Modal.vue'
import RegisterForm from '@shared/components/SignIn/RegisterForm.vue'
import LoginForm from '@shared/components/SignIn/LoginForm.vue'

const isOpen = ref<boolean>(false)
const showRegisterForm = ref<boolean>(false)

listenEvent(events.openLoginModal, () => (isOpen.value = true))
</script>

<template>
    <modal :opened="isOpen" @close="isOpen = false" max-width="550">
        <register-form v-if="showRegisterForm" @changed="showRegisterForm = false" />
        <login-form v-else @changed="showRegisterForm = true" />
    </modal>
</template>
