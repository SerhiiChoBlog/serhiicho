<script setup lang="ts">
import type { Series } from '@shared/types/models'
import MainButton from '@admin/components/Buttons/MainButton.vue'
import handleError from '@admin/modules/handleError'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { ref } from 'vue'
import showToast from '@shared/modules/showToast'

const { series } = defineProps<{ series: Series }>()
const router = useRouter()
const loading = ref(false)

function deleteSeries(): void {
    if (
        !confirm(
            `STEP 1/2. Are you sure you want to delete series "${series.title}"`,
        )
    )
        return

    if (!confirm(`STEP 2/2. Last confirmation. Are you sure?`)) return

    loading.value = true

    axios
        .delete(`/admin/ajax/series/${series.id}`)
        .then(() => {
            router.push({ name: 'series' })
            showToast({ text: `The series "${series.title}" has been removed` })
        })
        .catch(handleError)
        .finally(() => (loading.value = false))
}
</script>

<template>
    <div class="mt-7">
        <MainButton
            @click="deleteSeries"
            class="bg-red-500 hover:bg-red-600"
            type="button"
        >
            <i class="fas fa-trash mr-2"></i>
            Delete
        </MainButton>
    </div>
</template>
