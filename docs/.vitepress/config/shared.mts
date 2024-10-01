import { defineConfig, type DefaultTheme } from "vitepress";

// https://vitepress.dev/reference/site-config

export const shared = defineConfig({
    title: "Horizon Corps Docs",
    description: "A Horizon Corps Docs",

    rewrites: {
        "en/:rest*": ":rest*",
    },

    lastUpdated: true,
    cleanUrls: true,
    metaChunk: true,

    themeConfig: {
        search: {
            provider: "local",
        },
    },
});
