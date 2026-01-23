<script setup lang="ts">
import type { Goal } from '@shared/types'
import { onMounted, ref } from 'vue'
import axios from 'axios'
import Spinner from '@shared/components/Icons/Spinner.vue'
import PageTitle from '@admin/components/PageTitle.vue'

const goals = ref<Goal[]>([])
const loading = ref<boolean>(true)

onMounted(() => fetchGoals())

function fetchGoals(): void {
    axios
        .get<Goal[]>('/admin/ajax/dashboard/goals')
        .then(resp => (goals.value = resp.data))
        .catch(err => console.error(err))
        .finally(() => (loading.value = false))
}
</script>

<template>
    <PageTitle>Goals</PageTitle>

    <Spinner v-if="loading" />

    <div class="flex flex-col md:flex-row justify-center gap-10">
        <div
            v-for="goal in goals"
            :key="goal.title"
            class="w-52 aspect-square border-4 border-dashed rounded-full flex items-center justify-center"
            :class="goal.min >= goal.max ? 'border-green-500' : 'border-border'"
        >
            <div class="text-center">
                <h2 class="text-xl">{{ goal.title }}</h2>
                <h3 class="text-sm">{{ goal.goal }}</h3>
                <p class="text-4xl">{{ goal.min }} / {{ goal.max }}</p>
                <small>{{ goal.daysLeft }} days left</small>
            </div>
        </div>
    </div>
</template>
