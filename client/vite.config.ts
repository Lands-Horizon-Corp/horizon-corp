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
    visualizer({
      filename: 'bundle-analysis.html',
      template: 'treemap',
      open: true,
      gzipSize: true,
      brotliSize: true,
    }) as PluginOption,
  ],
  build: {
    sourcemap: true,
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ['react', 'react-dom'],
        },
      },
    },
  },
  define: {
    'process.env': process.env
  }
});
