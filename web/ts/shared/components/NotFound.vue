<template>
    <div class="container py-8">
        <div
            class="flex flex-col lg:flex-row gap-4 lg:gap-10 items-center justify-center"
        >
            <h1 class="text-8xl md:text-[10rem] font-main font-black leading-[120%]">
                404
            </h1>

            <h2 class="flex flex-col text-md md:text-xl lg:text-2xl font-post">
                <span>Page not found</span>
                <span>It might be that this URL is old</span>

                <div class="flex gap-4 mt-8 lg:mt-3 justify-center lg:justify-start">
                    <a
                        v-if="Auth"
                        :href="`/user/${Auth.slug}`"
                        class="border border-amber-400 hover:border-amber-500 shadow-md shadow-transparent duration-500 hover:shadow-amber-300 transition-shadow px-5 py-1.5 rounded-full text-center uppercase font-main text-sm min-w-[130px]"
                    >
                        Go to a profile
                    </a>
                </div>
            </h2>
        </div>
    </div>
</template>
