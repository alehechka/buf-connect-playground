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
	build: { target: 'es2020' },
	optimizeDeps: {
		esbuildOptions: { target: 'es2020', supported: { bigint: true } },
	},
});
