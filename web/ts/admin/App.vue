<script setup lang="ts">
import UserSidebar from './components/UserSidebar.vue'
</script>

<template>
    <user-sidebar />

    <div class="pl-9 lg:pl-14 py-10">
        <router-view />
    </div>
</template>
