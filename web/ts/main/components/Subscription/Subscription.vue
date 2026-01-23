<script setup lang="ts">
import type { Tag } from '@shared/types/models'
import type {
    SubscriptionDemandResponse,
    SubscriptionForm,
    SubscriptionTagsModal,
} from '@shared/types'
import AppearTransition from '@shared/components/Transitions/AppearTransition.vue'
import { onMounted, reactive, ref } from 'vue'
import axios from 'axios'
import showToast from '@shared/modules/showToast'
import CheckIcon from '@shared/components/Icons/CheckIcon.vue'
import Modal from '@shared/components/Modal.vue'
import SubscriptionModalContent from './SubscriptionModalContent.vue'
import ChooseTagsButton from './ChooseTagsButton.vue'
import SubscriptionTagIcon from './SubscriptionTagIcon.vue'
import SubscriptionButton from '@/modules/Subscription/SubscriptionButton'
import SubscriptionCache from '@/modules/Subscription/SubscriptionCache'
import listenEvent from '@shared/modules/listenEvent'
import { events } from '@shared/appConfig'

const selectedTags = ref<Tag[]>([])

const modalForm = reactive<SubscriptionForm>({
    loading: false,
    email: '',
})

const tags = reactive<SubscriptionTagsModal>({
    showModal: false,
    loading: false,
    items: [],
})

const advantages = [
    'Receive emails only for specific topics',
    'Your email is kept private and safe',
    'You can unsubscribe any time',
    'No advertisements',
    'No spam',
]

onMounted(() => {
    if (SubscriptionCache.isEmpty()) {
        SubscriptionCache.deleteAll()
    }

    SubscriptionButton.rename()

    updateChosenTags()

    listenEvent(events.subscribeTagsButtonTitle, () => {
        updateChosenTags()
    })
})

function updateChosenTags(): void {
    const cachedTags = SubscriptionCache.get()

    if (!cachedTags) {
        return
    }

    selectedTags.value = cachedTags
}

function showModalWithTags(): void {
    tags.showModal = true

    if (tags.items.length > 0) {
        return
    }

    tags.loading = true

    axios
        .get<Tag[]>('/api/tags/subscription')
        .then(resp => tags.items.push(...resp.data))
        .catch(err => console.error(err))
        .finally(() => (tags.loading = false))
}

function subscribe(): void {
    if (modalForm.email === '' || modalForm.loading) {
        return
    }

    modalForm.loading = true

    const uri = `/api/subscriptions?email=${
        modalForm.email
    }&tags=${generateTagsUrlParam()}`

    axios
        .post<SubscriptionDemandResponse>(uri)
        .then(resp => handleSubscribeResponse(resp.data))
        .catch(err => console.error(err))
        .finally(() => (modalForm.loading = false))
}

function handleSubscribeResponse(data: SubscriptionDemandResponse): void {
    modalForm.email = ''

    showToast({
        text: data.message,
        success: data.success,
        duration: 7000,
    })

    if (data.success) {
        SubscriptionCache.deleteAll()
        SubscriptionButton.setAllTags()
        selectedTags.value = []
    }
}

function generateTagsUrlParam() {
    if (SubscriptionCache.wantsAllTags() || allTagsSelected()) {
        return 'all'
    }

    if (SubscriptionCache.isEmpty()) {
        return 'none'
    }

    return selectedTags.value.map(t => t.id).join(',')
}

function allTagsSelected(): boolean {
    return selectedTags.value.length === tags.items.length
}
</script>

<template>
    <div>
        <appear-transition>
            <footer class="print:hidden">
                <div
                    class="container grid grid-cols-1 lg:grid-cols-[3fr_2fr] gap-7 md:gap-20 relative z-10"
                >
                    <form
                        @submit.prevent="subscribe"
                        class="w-full lg:w-[550px] mx-auto"
                    >
                        <label
                            for="subscription"
                            class="flex flex-col mb-4 drop-shadow-md pl-4"
                        >
                            <span class="text-lg lg:text-2xl"
                                >Don't miss the Next article</span
                            >
                            <span class="text-sm lg:text-lg">
                                You can subscribe to specific tags and be notified
                                when it's out
                            </span>
                        </label>

                        <div
                            class="flex h-[45px] lg:h-[53px] bg-gray-600 p-1 rounded-full shadow-xs"
                        >
                            <choose-tags-button
                                class="hidden md:inline-flex"
                                @clicked="showModalWithTags"
                            />

                            <input
                                type="email"
                                id="subscription"
                                class="px-6 h-full bg-transparent rounded-full w-full focus:outline-hidden text-white placeholder-gray-300"
                                placeholder="Your email address"
                                v-model="modalForm.email"
                            />

                            <button
                                type="submit"
                                class="transition-all hover:px-10 h-full px-4 lg:px-7 rounded-full shadow-md"
                                :class="
                                    modalForm.loading
                                        ? 'bg-main text-white'
                                        : 'bg-white text-gray-700'
                                "
                            >
                                <span v-if="modalForm.loading" class="flex gap-4">
                                    <b class="animate-spin inline-block">/</b>
                                    <span>Processing...</span>
                                </span>
                                <span v-else>Subscribe</span>
                            </button>
                        </div>

                        <div class="md:hidden text-center">
                            <choose-tags-button
                                class="w-full h-10 py-2.5 mt-2 text-center mx-auto"
                                @clicked="showModalWithTags"
                            />
                        </div>

                        <div class="text-center">
                            <div
                                class="inline-flex gap-2 flex-wrap mt-4 justify-center"
                            >
                                <subscription-tag-icon
                                    v-for="tag in selectedTags"
                                    :key="tag.id"
                                    :tag="tag"
                                />
                            </div>
                        </div>
                    </form>

                    <div class="mt-4 text-md lg:text-lg flex flex-col gap-3">
                        <ul class="space-y-1">
                            <li
                                v-for="(item, index) in advantages"
                                :key="item"
                                class="flex gap-2 items-center"
                            >
                                <span
                                    v-if="index === 0"
                                    class="mx-1 text-xl"
                                >❤️</span>

                                <check-icon
                                    v-else
                                    class="w-5 lg:w-7 h-5 lg:h-7 text-green-500"
                                />
                                <span>{{ item }}</span>
                            </li>
                        </ul>
                    </div>
                </div>
            </footer>
        </appear-transition>

        <modal
            v-if="!tags.loading"
            :opened="tags.showModal"
            @close="tags.showModal = false"
        >
            <subscription-modal-content :tags="tags" />
        </modal>
    </div>
</template>
