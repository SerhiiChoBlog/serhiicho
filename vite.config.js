import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import tailwindcss from '@tailwindcss/vite'

export default defineConfig({
    plugins: [
        vue({
            template: {
                transformAssetUrls: {
                    base: null,
                    includeAbsolute: false,
                },
            },
        }),
        tailwindcss(),
    ],

    build: {
        sourcemap: true,
        manifest: true,
        outDir: 'public/build',
        rollupOptions: {
            input: {
                // main
                main_js: 'web/ts/main/app.ts',
                main_css: 'web/css/main/app.css',

                // admin
                admin_js: 'web/ts/admin/app.ts',
                admin_css: 'web/css/admin/app.css',
                admin_fa: 'web/css/admin/font-awesome.css',

                // user
                user_js: 'web/ts/user/app.ts',
            },
        },
    },

    server: {
        host: '0.0.0.0',
        port: 5173,
        hmr: {
            host: 'localhost',
        },
        // proxy: {
        //     '/storage': {
        //         target: 'http://localhost',
        //         changeOrigin: true,
        //     },
        // },
    },

    resolve: {
        alias: {
            '@': '/web/ts/main',
            '@admin': '/web/ts/admin',
            '@user': '/web/ts/user',
            '@shared': '/web/ts/shared',
        },
    },
})
