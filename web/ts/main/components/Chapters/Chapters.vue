<script setup lang="ts">
import type { Chapter as ChapterType, SavedChapter } from '@shared/types'
import { onMounted, ref } from 'vue'
import isElementInViewport from '@/modules/isElementInViewport'
import screenSizeIs from '@/modules/screenSizeIs'
import sweetAlert from 'sweetalert2'
import showToast from '@shared/modules/showToast'
import toSlug from '@shared/modules/toSlug'
import Chapter from './Chapter.vue'
import ChevronDownIcon from '@shared/components/Icons/ChevronDownIcon.vue'
import ChevronUpIcon from '@shared/components/Icons/ChevronUpIcon.vue'
import SaveIcon from '@shared/components/Icons/SaveIcon.vue'
import DeleteIcon from '@shared/components/Icons/DeleteIcon.vue'
import LoadIcon from '@shared/components/Icons/LoadIcon.vue'
import ControlButton from './ControlButton.vue'
import SlideUpTransition from '@shared/components/Transitions/SlideUpTransition.vue'

interface Props {
    slug: string
}

const { slug } = defineProps<Props>()
const headers = ref<NodeListOf<HTMLElement> | null>(null)
const chapters = ref<ChapterType[]>([])
const savedChapter = ref<SavedChapter | null>(null)
const currentChapter = ref<ChapterType | null>(null)
const showChapters = ref(false)

onMounted(() => {
    if (!screenSizeIs('lg')) {
        showChapters.value = true
    }

    headers.value = document.querySelectorAll<HTMLElement>(
        '#post-content h2, #post-content h3',
    )

    headers.value.forEach(elem => {
        createChapterFromElement(elem)
        window.addEventListener('scroll', () => markVisibleChapters(elem))
    })

    setSavedChapter()
    detectCurrentChapter()
    showAlertIfHasSavedChapter()

    scrollToChapterWhenHasHashInUrl()
    window.addEventListener('hashchange', () => scrollToChapterWhenHasHashInUrl())
})

function scrollToChapterWhenHasHashInUrl(): void {
    const currHash = window.location.hash

    if (!currHash) {
        return
    }

    const hashName = currHash.replace('#', '')
    const chapter = chapters.value.find(chap => chap.slug === hashName)

    if (chapter) {
        scrollToChapter(chapter.title)
    }
}

function setSavedChapter(): void {
    const saved = localStorage.getItem(`chapter:${slug}`) as string | null

    if (!saved) {
        return
    }

    savedChapter.value = JSON.parse(saved) as SavedChapter
}

async function showAlertIfHasSavedChapter(): Promise<void> {
    if (!savedChapter.value) {
        return
    }

    const response = await sweetAlert.fire({
        title: 'Welcome back',
        html: `You have a saved chapter <b>"${savedChapter.value.chapter.title}"</b>. Do you want to continue reading where you've left off?`,
        showDenyButton: true,
        confirmButtonText: 'Continue reading',
        denyButtonText: 'Delete save',
    })

    const wantsLoadSave = response.isConfirmed
    const wantsDeleteSave = response.isDenied

    if (wantsDeleteSave) {
        deleteSave()
    }

    if (wantsLoadSave) {
        setTimeout(() => goToSavedChapter(), 300)
    }
}

function deleteSave(): void {
    localStorage.removeItem(`chapter:${slug}`)
    savedChapter.value = null

    showToast({
        text: `Your saved position in the post has been deleted!`,
        success: true,
        duration: 6000,
    })
}

function goToSavedChapter(): void {
    if (!savedChapter.value) {
        return
    }

    if (screenSizeIs('lg')) {
        showChapters.value = false
    }

    window.scrollTo({ top: savedChapter.value!.top, behavior: 'smooth' })
}

function scrollToChapter(title: string): void {
    if (!headers.value) {
        return
    }

    const neededElement = Array.from(headers.value).find(header => {
        return title === header.innerText
    })!

    const yOffset = -90
    const y = neededElement.getBoundingClientRect().top + window.scrollY + yOffset

    window.scrollTo({ top: y, behavior: 'smooth' })
}

function detectCurrentChapter(): void {
    const activeChapter = chapters.value.find(chapter => chapter.isActive)

    if (activeChapter) {
        currentChapter.value = activeChapter
        return
    }

    currentChapter.value = null

    if (screenSizeIs(['sm', 'md', 'lg'])) {
        showChapters.value = false
    }
}

function markVisibleChapters(elem: HTMLElement): void {
    const elemIsVisible = isElementInViewport(elem)

    if (!elemIsVisible) {
        return
    }

    chapters.value.map(chapter => {
        chapter.isActive =
            window.scrollY > 100 ? chapter.title === elem.innerText : false
        return chapter
    })

    detectCurrentChapter()
}

function saveCurrentChapter(): void {
    if (!currentChapter.value) {
        return
    }

    savedChapter.value = {
        chapter: currentChapter.value,
        top: window.scrollY,
    }

    localStorage.setItem(`chapter:${slug}`, JSON.stringify(savedChapter.value))

    showToast({
        text: "You've saved current position in the post. You'll be able to go back to it at any time",
        success: true,
        duration: 6000,
    })

    if (screenSizeIs('lg')) {
        showChapters.value = false
    }
}

function createChapterFromElement(elem: HTMLElement): void {
    const title = elem.innerText

    const chapter: ChapterType = {
        title,
        isActive: false,
        tag: elem.tagName as 'H2' | 'H3',
        slug: toSlug(title),
    }

    chapters.value.push(chapter)
}

function handleChapterClick(title: string): void {
    scrollToChapter(title)

    if (screenSizeIs('lg')) {
        showChapters.value = false
    }
}
</script>

<template>
    <button
        v-if="currentChapter"
        class="flex print:hidden lg:hidden items-center shadow-md gap-x-1 justify-center fixed border border-border top-[54px] bg-page w-full py-1 px-4 z-40 left-1/2 -translate-x-1/2"
        @click="showChapters = !showChapters"
    >
        <chevron-up-icon v-if="showChapters" class="w-6 h-6 inline-block" />
        <chevron-down-icon v-else class="size-6 inline-block" />
        <span class="tracking-wide">{{ currentChapter.title }}</span>
    </button>

    <slide-up-transition>
        <div
            v-if="showChapters"
            class="bg-page lg:bg-transparent print:hidden lg:h-[calc(100%-200px)] fixed w-full max-w-md left-1/2 lg:left-0 -translate-x-1/2 lg:translate-x-0 border border-border lg:border-none top-16 lg:top-0 bottom-[3%] p-4 lg:p-0 shadow-xl lg:shadow-none lg:relative z-30"
        >
            <nav
                class="sticky h-screen overflow-y-auto scrollbar pr-2 top-3 lg:top-24 mt-7 text-center lg:text-left pb-[70px]"
            >
                <div class="mb-4">
                    <div class="flex items-center gap-x-5 text-font-second">
                        <span class="text-2xl font-bold"> Chapters </span>

                        <div class="flex items-center gap-x-4 text-font-second">
                            <control-button
                                tipContent="Load saved position and get back to where you've left off"
                                :disabled="!savedChapter"
                                label="Scroll to a saved position"
                                @click="goToSavedChapter"
                            >
                                <LoadIcon class="w-4" />
                            </control-button>

                            <control-button
                                @click="saveCurrentChapter"
                                :disabled="false"
                                label="Save the current position"
                                tipContent="Save current position. You'll be able to load it and get back to where you've left off"
                            >
                                <SaveIcon class="w-4" />
                            </control-button>

                            <control-button
                                @click="deleteSave"
                                label="Delete saved position"
                                tipContent="Delete current saved position"
                                :disabled="!savedChapter"
                            >
                                <DeleteIcon class="w-4" />
                            </control-button>
                        </div>
                    </div>
                </div>

                <div v-if="headers">
                    <chapter
                        v-for="chapter in chapters"
                        :key="chapter.slug"
                        :chapter="chapter"
                        :headers="headers"
                        @clicked="handleChapterClick"
                    />
                </div>
            </nav>
        </div>
    </slide-up-transition>
</template>
