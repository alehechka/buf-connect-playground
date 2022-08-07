import { useState } from 'react';
import useUser from './proto/hooks/useUser';

function App() {
	const [userId, setUserID] = useState('');
	const [user, loading, error, fetchUser] = useUser();

	return (
		<div>
			<input value={userId} onChange={(event) => setUserID(event.target.value)} />
			<button onClick={() => fetchUser(userId)}>submit</button>
			<br />
			{loading && <div>loading...</div>}
			{user && <code style={{ color: 'blue' }}>{JSON.stringify(user)}</code>}
			{error && <code style={{ color: 'red' }}>{JSON.stringify(error)}</code>}
		</div>
	);
}

export default App;
