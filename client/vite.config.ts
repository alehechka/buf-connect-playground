import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [react()],
	server: {
		port: 8080,
		host: true,
		watch: {
			usePolling: true,
		},
		proxy: {
			'/api': {
				target: 'http://server:3000',
				changeOrigin: true,
			},
		},
	},
});
