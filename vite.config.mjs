import { defineConfig } from "vite";

export default defineConfig({
  build: {
    outDir: "./cmd/web/assets/runtime",
    assetsDir: "output",
    rollupOptions: {
      input: {
        app: "./internal/assets/app.js",
      },
      output: {
        manualChunks(id, { getModuleInfo }) {
          if (id.includes("node_modules")) {
            return "vendor";
          }
        },
      },
    },
  },
});
