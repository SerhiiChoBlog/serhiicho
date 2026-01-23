<script setup lang="ts">
import Spinner from '@shared/components/Icons/Spinner.vue'

type Props = {
    type: 'submit' | 'button'
    loading?: boolean
}

const props = defineProps<Props>()
</script>

<template>
    <div class="text-center">
        <button
            :type="type"
            class="cursor-pointer inline-block min-w-[200px] mx-auto px-5 py-3 text-main font-bold uppercase border-2 border-main rounded-full hover:border-main-hover hover:text-main-hover transition-colors text-sm lg:text-md"
            :class="{
                'cursor-not-allowed opacity-80 pointer-events-none': loading,
            }"
        >
            <span v-if="loading" class="flex items-center">
                <spinner class="size-5!" /> Processing...
            </span>

            <slot v-else />
        </button>
    </div>
</template>
