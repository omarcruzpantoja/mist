import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

const APP_PORT = parseInt(process.env.APP_PORT || "3050", 10);
// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: APP_PORT,
  },
});
