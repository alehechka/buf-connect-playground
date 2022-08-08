import { createPromiseClient, createConnectTransport } from '@bufbuild/connect-web';
import { UsersService } from './gen/users/v1/user_connectweb';

export const userClient = createPromiseClient(UsersService, createConnectTransport({ baseUrl: '/api' }));
