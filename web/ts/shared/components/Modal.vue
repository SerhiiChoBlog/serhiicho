<script setup lang="ts">
import { onMounted, ref, watch } from 'vue'
import AppearTransition from '@shared/components/Transitions/AppearTransition.vue'
import CloseIcon from '@shared/components/Icons/CloseIcon.vue'

type Props = {
    opened: boolean
    maxWidth?: string
    empty?: boolean
}

const emit = defineEmits<{ (e: 'close'): void }>()
const props = defineProps<Props>()

const isOpen = ref<boolean>(props.opened)

watch(props, (props: Props) => {
    isOpen.value = props.opened
})

onMounted(() => {
    window.addEventListener('keydown', e => {
        if (e.key === 'Escape') {
            emit('close')
        }
    })
})
</script>

<template>
    <appear-transition>
        <aside v-show="isOpen" class="fixed inset-0 z-50 bg-black/80">
            <div
                v-if="isOpen"
                :class="[
                    'overflow-y-auto absolute top-1/2 left-1/2 -translate-x-1/2',
                    '-translate-y-1/2 w-full max-w-250 mx-auto pt-10',
                    'max-h-[calc(100vh-40px)]',
                ]"
                :style="{ maxWidth: maxWidth + 'px' }"
            >
                <div
                    class="bg-page text-font shadow-xl rounded-lg h-full relative"
                    :class="empty ? '' : 'md:p-10 p-5'"
                >
                    <button
                        @click="emit('close')"
                        type="button"
                        aria-label="Close contacts modal button"
                        :class="[
                            'absolute top-2 right-2 p-3 text-font-second',
                            'hover:text-red-600 transition-colors z-10 cursor-pointer',
                        ]"
                    >
                        <close-icon class="w-6 h-6" />
                    </button>

                    <slot />
                </div>
            </div>
        </aside>
    </appear-transition>
</template>
