<script setup lang="ts">
import type { Testimonial as TestimonialType } from '@shared/types'
import { ref, computed, onMounted } from 'vue'
import axios from 'axios'
import handleServerError from '@shared/modules/handleServerError'
import Spinner from '@shared/components/Icons/Spinner.vue'
import Testimonial from '@/components/Testimonials/Testimonial.vue'
import Modal from '@shared/components/Modal.vue'

const loading = ref<boolean>(false)
const testimonials = ref<TestimonialType[]>([])
const fullContent = ref<TestimonialType | null>(null)

onMounted(() => fetchTestimonials())

function fetchTestimonials(): void {
    if (loading.value) {
        return
    }

    loading.value = true

    axios
        .get<TestimonialType[]>('/api/about-me/testimonials')
        .then(rep => (testimonials.value = rep.data))
        .catch(handleServerError)
        .finally(() => (loading.value = false))
}

const groupedTestimonials = computed<TestimonialType[][]>(() => {
    if (testimonials.value.length === 0) {
        return []
    }

    const total = testimonials.value.length
    const testimonialsPerRow = Math.round(total / 3)

    return [
        testimonials.value.slice(0, testimonialsPerRow),
        testimonials.value.slice(testimonialsPerRow, testimonialsPerRow * 2),
        testimonials.value.slice(testimonialsPerRow * 2, total),
    ]
})
</script>

<template>
    <modal
        v-if="fullContent"
        :empty="true"
        :opened="fullContent !== null"
        @close="() => (fullContent = null)"
        max-width="700"
    >
        <testimonial :testimonial="fullContent" :show-full="true" />
    </modal>

    <spinner v-if="loading" />

    <div
        v-else
        class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 lg:gap-7"
    >
        <div
            v-for="(groupedTestimonial, i) in groupedTestimonials"
            :key="i"
            class="space-y-4 lg:space-y-7"
        >
            <testimonial
                v-for="testimonial in groupedTestimonial"
                :key="testimonial.name"
                :testimonial="testimonial"
                :show-full="false"
                @showFullContent="() => (fullContent = testimonial)"
            />
        </div>
    </div>
</template>
