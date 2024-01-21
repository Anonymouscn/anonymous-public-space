import {defineConfig, loadEnv} from 'vite'
import vue from '@vitejs/plugin-vue'
import { createSvgIconsPlugin } from 'vite-plugin-svg-icons'
import path from "path"
import viteCompression from 'vite-plugin-compression'
import viteImagemin from 'vite-plugin-imagemin'
import { visualizer } from "rollup-plugin-visualizer";

// https://vitejs.dev/config/
export default ({mode}) => defineConfig({
    plugins: [
        visualizer({
                open: false,  // 打包后自动打开页面
                gzipSize: true,  // 查看 gzip 压缩大小
                brotliSize: true, // 查看 brotli 压缩大小
                filename: "report.html", //分析图生成的文件名
        }),
        vue(),
        viteCompression(), // 打包压缩
        createSvgIconsPlugin({
            // 指定需要缓存的图标文件夹(路径为存放所有svg图标的文件夹不单个svg图标)
            iconDirs: [path.resolve(process.cwd(), 'src/assets/icon')],
            // 指定symbolId格式
            symbolId: 'icon-[dir]-[name]'
        }),
        viteImagemin({
            gifsicle: {
                optimizationLevel: 7,
                interlaced: false
            },
            optipng: {
                optimizationLevel: 7
            },
            mozjpeg: {
                quality: 20
            },
            pngquant: {
                quality: [0.8, 0.9],
                speed: 4
            },
            svgo: {
                plugins: [
                    {
                        name: 'removeViewBox'
                    },
                    {
                        name: 'removeEmptyAttrs',
                        active: false
                    }
                ]
            }
        }),
    ],
    base: loadEnv(mode, process.cwd()).VITE_APP_BASE_URL || '/',
    build: {
        outDir: 'dist',
        assetsDir: 'static',
        cssCodeSplit: true, // 如果设置为false，整个项目中的所有 CSS 将被提取到一个 CSS 文件中
    }
})