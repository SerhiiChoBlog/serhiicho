<script lang="ts" setup>
import Popper from 'vue3-popper'

interface Props {
    content: string | null
    icon?: boolean
    classes?: string
}

const { content, icon = false, classes } = defineProps<Props>()
</script>

<template>
    <div v-if="content" class="inline-flex">
        <popper hover :content="content" class="border-none! m-0!">
            <div>
                <svg
                    v-if="icon"
                    xmlns="http://www.w3.org/2000/svg"
                    fill="none"
                    viewBox="0 0 24 24"
                    stroke-width="1.5"
                    stroke="currentColor"
                    :class="classes ? classes : 'w-5'"
                    class="aspect-square"
                >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        d="M9.879 7.519c1.171-1.025 3.071-1.025 4.242 0 1.172 1.025 1.172 2.687 0 3.712-.203.179-.43.326-.67.442-.745.361-1.45.999-1.45 1.827v.75M21 12a9 9 0 11-18 0 9 9 0 0118 0zm-9 5.25h.008v.008H12v-.008z"
                    />
                </svg>

                <slot />
            </div>

            <template #content>
                <div
                    class="bg-gray-800/90 px-4 py-2 text-white rounded-md text-sm lg:text-md shadow-md tracking-normal font-normal w-auto min-w-[120px]"
                >
                    {{ content }}
                </div>
            </template>
        </popper>
    </div>

    <slot v-else />
</template>
