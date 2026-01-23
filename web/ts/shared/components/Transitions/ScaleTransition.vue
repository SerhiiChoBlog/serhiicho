<template>
    <transition
        enter-active-class="duration-300 ease-out"
        enter-from-class="transform opacity-0 scale-95"
        enter-to-class="opacity-100 scale-100"
        leave-active-class="duration-300 ease-in"
        leave-from-class="opacity-100 scale-100"
        leave-to-class="transform opacity-0 scale-95"
        appear
    >
        <slot />
    </transition>
</template>
