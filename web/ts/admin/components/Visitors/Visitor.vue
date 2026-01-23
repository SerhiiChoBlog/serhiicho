<script setup lang="ts">
import type { Visitor } from '@shared/types/models'
import BodyDetail from '@admin/components/DataTable/BodyDetail.vue'
import BodyRow from '@admin/components/DataTable/BodyRow.vue'
import useBlockVisitor from '@admin/composables/useBlockVisitor'

const emit = defineEmits<{ (e: 'refreshVisitors'): void }>()
const props = defineProps<{ visitor: Visitor }>()
const { blockVisitor } = useBlockVisitor()

function blockVisitorHandler(): void {
    const message = props.visitor.is_blocked
        ? 'Are you sure you want to unblock this visitor?'
        : 'Are you sure you want to block this visitor?'

    if (!confirm(message)) return

    blockVisitor(props.visitor.id, !props.visitor.is_blocked)
    emit('refreshVisitors')
}
</script>

<template>
    <BodyRow>
        <BodyDetail>{{ visitor.id }}</BodyDetail>
        <BodyDetail class="text-center">
            <span v-if="visitor.is_robot">🤖</span>
            <span v-else>👦🏻</span>
        </BodyDetail>
        <BodyDetail>{{ visitor.ip }}</BodyDetail>
        <BodyDetail>
            <img
                v-if="visitor.code"
                :src="`/storage/admin/flags/${visitor.code.toLowerCase()}.png`"
                width="80"
                height="42"
                class="w-9 h-5 rounded-xs shadow-xs border border-border"
            />
            <img
                v-else
                src="/storage/admin/flags/unknown.png"
                width="80"
                height="56"
            />
        </BodyDetail>
        <BodyDetail>{{ visitor.country || '-' }}</BodyDetail>
        <BodyDetail>{{ visitor.city || '-' }}</BodyDetail>
        <BodyDetail>{{ visitor.post_views_count }}</BodyDetail>
        <BodyDetail>
            <button @click="blockVisitorHandler" type="button">
                <span v-if="visitor.is_blocked"
                    ><i class="fas fa-check text-green-700"></i
                ></span>
                <span v-else><i class="fas fa-ban text-red-700"></i></span>
            </button>
        </BodyDetail>
    </BodyRow>
</template>
