import { defineConfig, type DefaultTheme } from "vitepress";

export const enConfig = defineConfig({
    lang: "en-US",
    themeConfig: {
        // https://vitepress.dev/reference/default-theme-config
        nav: nav(),

        sidebar: {
            "/api/": [
                {
                    text: "About",
                    items: [{ text: "Introduction", link: "/api/" }],
                },
                {
                    text: "Api Docs",
                    items: [
                        {
                            text: "API v1",
                            base: "/api/v1",
                            collapsed: true,
                            items: [
                                { text: "Introduction", link: "/" },
                                { text: "Auth", link: "/auth" },
                            ],
                        },
                    ],
                },
            ],
        },

        socialLinks: [
            { icon: "github", link: "https://github.com/Lands-Horizon-Corp" },
        ],

        footer: {
            message: "Horizon Corps Dev Team",
            copyright: "Copyright © 2024-present Horizon Corps",
        },
    },
});

function nav(): DefaultTheme.NavItem[] {
    return [
        {
            text: "About",
            link: "/about",
            activeMatch: "/about/",
        },
        {
            text: "Api Docs",
            link: "/api",
            activeMatch: "/api/",
        },
    ];
}
