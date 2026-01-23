<script setup lang="ts">
import { ref } from 'vue'
import contactModalSections from '@/modules/contactModalSections'
import Modal from '@shared/components/Modal.vue'

const isOpen = ref<boolean>(false)
</script>

<template>
    <button
        type="button"
        class="main-btn"
        @click="isOpen = true"
        aria-label="Open contacts modal"
    >
        Contacts
    </button>

    <modal :opened="isOpen" @close="isOpen = false">
        <div class="divide-y divide-border">
            <section
                v-for="section in contactModalSections"
                :key="section.title"
                class="py-5 first:pt-0 last:pb-0"
            >
                <h2 class="text-center text-2xl mb-3 font-bold">
                    {{ section.title }}
                </h2>

                <p v-html="section.content" class="text-center text-lg"></p>

                <div class="flex justify-center gap-x-5 mt-5">
                    <a
                        v-for="link in section.links"
                        :key="link.title"
                        :href="link.href"
                        target="_blank"
                        :title="link.title"
                        class="transition-transform hover:scale-125"
                        :aria-label="link.title"
                    >
                        <img
                            :src="link.img"
                            :alt="link.title"
                            width="30"
                            height="30"
                            class="rounded-md"
                        />
                    </a>
                </div>
            </section>
        </div>
    </modal>
</template>
