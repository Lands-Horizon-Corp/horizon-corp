import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { config } from 'dotenv';
// import { TanStackRouterVite } from '@tanstack/router-plugin/vite';

import tsconfigPaths from "vite-tsconfig-paths";

config();

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    // TanStackRouterVite(),
    react(),
    tsconfigPaths(),
  ],

  define: {
    'process.env': process.env
  }
});
