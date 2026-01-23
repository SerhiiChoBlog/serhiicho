<script setup lang="ts">
import links from '@admin/router/links'
</script>

<template>
    <nav
        :class="[
            'bg-gray-800 opacity-75 fixed top-0 pt-16 bottom-0 w-11 overflow-hidden',
            'transition-all translate-z-0 z-40 hover:overflow-visible',
            'hover:w-[180px] hover:h-full hover:overflow-y-auto hover:opacity-95',
        ]"
    >
        <ul>
            <li v-for="link in links" :key="link.title">
                <router-link
                    :to="{ name: link.routeName }"
                    :class="[
                        'relative w-[180px] m-auto transition-colors flex items-center border-collapse',
                        'text-white text-md leading-10 hover:bg-border py-0.5',
                    ]"
                >
                    <i
                        class="fas w-[27px] text-2xl text-center ml-[7px] mr-3"
                        :class="link.icon"
                    ></i>

                    <span>{{ link.title }}</span>
                </router-link>
            </li>
        </ul>
    </nav>
</template>
