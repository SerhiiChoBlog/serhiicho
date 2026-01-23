<script setup lang="ts">
import useLogin from '@/composables/useLogin'
import SignInHeader from '@shared/components/SignIn/SignInHeader.vue'
import InputField from '@/components/Form/InputField.vue'
import FormLabel from '@/components/Form/FormLabel.vue'
import FormButton from '@/components/Form/FormButton.vue'
import Tip from '@shared/components/Tip.vue'
import SignInWith from '@shared/components/SignIn/SignInWith.vue'
import QuestionLink from '@shared/components/SignIn/QuestionLink.vue'

const emit = defineEmits<{
    (e: 'changed'): void
}>()

const { useAsSeparatePage = false } = defineProps<{
    useAsSeparatePage?: boolean
}>()

const { loading, formData, errors, login, redirectToReset } = useLogin()

function switchToSignUp() {
    if (useAsSeparatePage) {
        window.location.href = '/register'
        return
    }

    emit('changed')
}
</script>

<template>
    <sign-in-header
        headline="Sign in"
        sub-headline="Welcome back! Please sign in to continue"
    />

    <form @submit.prevent="login" class="space-y-5 my-6">
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
                placeholder="Enter your email"
                v-model="formData.email"
                @changed="v => (formData.email = v)"
                :error-message="errors && errors.email ? errors.email[0] : null"
            />
        </div>

        <div>
            <form-label for="password" id="password">Password</form-label>
            <input-field
                name="password"
                id="login-password"
                type="password"
                v-model="formData.password"
                placeholder="Enter your password"
                @changed="v => (formData.password = v)"
            />
        </div>

        <div class="flex gap-2 items-center">
            <input type="checkbox" v-model="formData.remember" id="remember-me" />
            <label for="remember-me" id="remember"> Keep me signed in </label>
        </div>

        <div class="pt-4">
            <form-button :loading type="submit">Sign in</form-button>
            <sign-in-with />
        </div>
    </form>

    <div class="flex items-center justify-center gap-9">
        <question-link
            question="Don't have an account?"
            title="Sign up"
            @clicked="switchToSignUp"
        />

        <question-link
            question="Forgot your password?"
            title="Reset password"
            @clicked="redirectToReset"
        />
    </div>
</template>
