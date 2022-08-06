import { useCallback, useState } from 'react';
import { UsersService } from './proto/gen/users/v1/user_connectweb';
import { User } from './proto/gen/users/v1/user_pb';
import { createPromiseClient, createConnectTransport } from '@bufbuild/connect-web';

function App() {
	const [userId, setUserID] = useState('');

	const [user, setUser] = useState<User>();

	const client = createPromiseClient(
		UsersService,
		createConnectTransport({
			baseUrl: 'http://localhost:8080',
		})
	);

	const getUser = useCallback(() => {
		client
			.getUser({ userId })
			.then((res) => setUser(res.user))
			.catch(() => setUser(undefined));
	}, [userId]);

	return (
		<div>
			<input value={userId} onChange={(event) => setUserID(event.target.value)} />
			<button onClick={getUser}>submit</button>
			{user && <code>{JSON.stringify(user)}</code>}
		</div>
	);
}

export default App;
