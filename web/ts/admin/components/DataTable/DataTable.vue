<template>
    <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
        <table class="w-full">
            <thead class="bg-gray-700 text-left">
                <tr>
                    <slot name="head" />
                </tr>
            </thead>
            <tbody>
                <slot />
            </tbody>
        </table>
    </div>
</template>
