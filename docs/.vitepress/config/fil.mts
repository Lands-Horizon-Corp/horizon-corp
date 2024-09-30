import { defineConfig, type DefaultTheme } from "vitepress";

export const filConfig = defineConfig({
    lang: "fil-PH",
    themeConfig: {
        // https://vitepress.dev/reference/default-theme-config
        nav: nav(),

        sidebar: {
            "/fil/api/": [
                {
                    text: "Tungkol",
                    base: "/fil",
                    items: [{ text: "Panimula", link: "/api/" }],
                },
                {
                    text: "Dokumentasyon ng API",
                    base: "/fil/api",
                    items: [
                        {
                            text: "API v1",
                            base: "/fil/api/v1",
                            collapsed: true,
                            items: [
                                { text: "Panimula", link: "/" },
                                { text: "Pagpapatunay (Auth)", link: "/auth" },
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
            message: "Koponan ng mga Nag-develop ng Horizon Corps",
            copyright: "Karapatang-ari © 2024-kasalukuyan Horizon Corps",
        },
    },
});

function nav(): DefaultTheme.NavItem[] {
    return [
        {
            text: "Tungkol",
            link: "/fil/about",
            activeMatch: "/fil/about/",
        },
        {
            text: "Dokumentasyon ng API",
            link: "/fil/api",
            activeMatch: "/fil/api/",
        },
    ];
}
