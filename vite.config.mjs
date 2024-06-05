import { defineConfig } from "vite";

export default defineConfig({
  build: {
    outDir: "./cmd/web/assets/runtime",
    assetsDir: "output",
    rollupOptions: {
      input: {
        app: "./internal/assets/app.js",
        trends: "./internal/assets/pokemon-trends-carousel.js",
      },
      output: {
        manualChunks(id, { getModuleInfo }) {
          if (id.includes("htmx.org") || id.includes("alpinejs")) {
            return "vendor";
          }
        },
      },
    },
  },
});
