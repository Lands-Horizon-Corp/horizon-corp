// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2024-04-03",
  devtools: { enabled: true },
  devServer: {
    port: import.meta.env.CLIENT_PORT || '80'
  },
  modules: [
    "@nuxtjs/tailwindcss",
    "@nuxt/eslint",
    "shadcn-nuxt",
    "@vee-validate/nuxt",
    "@pinia/nuxt",
  ],
  css: ["~/assets/css/tailwind.css"],
  shadcn: {
    prefix: "",
    componentDir: "./components/ui",
  },
  postcss: {
    plugins: {
      tailwindcss: {},
      autoprefixer: {},
    },
  },
});
