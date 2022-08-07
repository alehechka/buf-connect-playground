import { createPromiseClient, createConnectTransport } from '@bufbuild/connect-web';
import { UsersService } from './gen/users/v1/user_connectweb';

export const userClient = createPromiseClient(
	UsersService,
	createConnectTransport({
		baseUrl: import.meta.env.VITE_GRPC_BACKEND_HOST || 'http://localhost:3000',
	})
);
