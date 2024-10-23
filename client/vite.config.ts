import { defineConfig, type PluginOption } from "vite";
import react from "@vitejs/plugin-react";
import { visualizer } from "rollup-plugin-visualizer";
// import { TanStackRouterVite } from '@tanstack/router-plugin/vite';

import tsconfigPaths from "vite-tsconfig-paths";


// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    // TanStackRouterVite(),
    react(),
    tsconfigPaths(),
    visualizer() as PluginOption,
  ],

  define: {
    'process.env': process.env
  }
});
