<script setup lang="ts">
import SignInHeader from '@shared/components/SignIn/SignInHeader.vue'
import InputField from '@/components/Form/InputField.vue'
import FormLabel from '@/components/Form/FormLabel.vue'
import FormButton from '@/components/Form/FormButton.vue'
import Tip from '@shared/components/Tip.vue'
import SignInWith from '@shared/components/SignIn/SignInWith.vue'
import QuestionLink from '@shared/components/SignIn/QuestionLink.vue'
import useRegister from '@/composables/useRegister'

const emit = defineEmits<{
    (e: 'changed'): void
}>()

const { useAsSeparatePage = false } = defineProps<{
    useAsSeparatePage?: boolean
}>()

const { loading, formData, errors, register } = useRegister()

function switchToSignUp() {
    if (useAsSeparatePage) {
        window.location.href = '/login'
        return
    }

    emit('changed')
}
</script>

<template>
    <sign-in-header
        headline="Register"
        sub-headline="Create an account to unlock extra features only for registered users"
    />

    <form @submit.prevent="register" class="space-y-5 my-6">
        <div>
            <form-label for="name" id="name">
                Name
                <tip
                    :icon="true"
                    content="The name is public and will be visible on your comments"
                />
            </form-label>
            <input-field
                name="name"
                id="name"
                type="text"
                placeholder="Enter your name"
                v-model="formData.name"
                @changed="v => (formData.name = v)"
                :error-message="errors && errors.name ? errors.name[0] : null"
            />
        </div>

        <div>
            <form-label for="email" id="email">
                Email
                <tip
                    :icon="true"
                    content="We never share email addresses with anyone else"
                />
            </form-label>
            <input-field
                name="email"
                id="login-email"
                type="email"
                v-model="formData.email"
                placeholder="Enter your email"
                @changed="v => (formData.email = v)"
                :error-message="errors && errors.email ? errors.email[0] : null"
            />
        </div>

        <div>
            <form-label for="password" id="password">Password</form-label>
            <input-field
                name="password"
                id="register-password"
                type="password"
                placeholder="Enter your password"
                v-model="formData.password"
                @changed="v => (formData.password = v)"
                :error-message="
                    errors && errors.password ? errors.password[0] : null
                "
            />
        </div>

        <div class="pt-4">
            <form-button :loading="loading" type="submit">Register</form-button>
            <sign-in-with />
        </div>
    </form>

    <question-link
        question="Already have an account?"
        title="Sign in"
        @clicked="switchToSignUp"
    />
</template>
