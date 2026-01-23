<script setup lang="ts">
import MenuItems from '@shared/components/Dropdown/MenuItems.vue'
import MenuItem from '@shared/components/Dropdown/MenuItem.vue'
import { Menu, MenuButton } from '@headlessui/vue'
import logout from '@shared/modules/logout'

function redirectTo(url: string) {
    window.location.href = url
}

const profileUrl =
    Auth && Auth.slug === 'serhii-cho' ? '/admin/dashboard' : '/users/profile'
</script>

<template>
    <Menu>
        <menu-button>
            <slot />
        </menu-button>

        <menu-items>
            <menu-item @click="redirectTo(profileUrl)">Profile</menu-item>
            <menu-item @click="redirectTo('/users/settings')">Settings</menu-item>
            <menu-item @click="logout">Logout</menu-item>
        </menu-items>
    </Menu>
</template>
