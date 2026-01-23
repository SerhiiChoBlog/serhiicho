<script setup lang="ts">
import type { Section } from '@shared/types'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import Spinner from '@shared/components/Icons/Spinner.vue'
import PageTitle from '@admin/components/PageTitle.vue'

const sections = ref<Section[]>([])
const loading = ref<boolean>(true)

onMounted(() => fetchSections())

function fetchSections(): void {
    axios
        .get<Section[]>('/admin/ajax/dashboard/sections')
        .then(resp => (sections.value = resp.data))
        .catch(err => console.error(err))
        .finally(() => (loading.value = false))
}
</script>

<template>
    <PageTitle>Statistics</PageTitle>

    <Spinner v-if="loading" />

    <nav
        v-else
        class="grid grid-cols-2 md:grid-cols-3 xl:grid-cols-6 max-w-[1200px] mx-auto gap-3 lg:gap-6"
    >
        <div
            v-for="section in sections"
            :key="section.title"
            class="flex flex-col justify-between"
        >
            <div
                class="text-center uppercase text-xs lg:text-sm text-font-second tracking-widest"
            >
                {{ section.title }}
            </div>

            <div
                class="text-2xl lg:text-3xl text-center font-semibold text-font drop-shadow-md"
            >
                {{ section.value }}
            </div>
        </div>
    </nav>
</template>
