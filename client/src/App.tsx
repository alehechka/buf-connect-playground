import { useState } from 'react';
import useGenerateUsers from './proto/hooks/useGenerateUsers';
import useUser from './proto/hooks/useUser';
import useUsers from './proto/hooks/useUsers';

function App() {
	const [userId, setUserID] = useState('');
	const [user, loading, error, fetchUser] = useUser();
	const [users, usersLoading, usersError, fetchUsers] = useUsers();
	const [numCreated, genUsersLoading, genUsersError, generateUsers] = useGenerateUsers();

	return (
		<div>
			<section style={{ padding: '10px', borderStyle: 'solid', marginBottom: '10px' }}>
				<input value={userId} onChange={(event) => setUserID(event.target.value)} />
				<button onClick={() => fetchUser(userId)}>submit</button>
				<br />
				{loading && <div>loading...</div>}
				{user && (
					<div>
						<h3>
							{user.firstName} {user.lastName}
						</h3>
						{user.gender > 0 && <p>Gender: {user.gender === 1 ? 'male' : 'female'}</p>}
						{user.birthday && (
							<p>
								Birthday: {user.birthday?.month}/{user.birthday?.day}/{user.birthday?.year}
							</p>
						)}
					</div>
				)}
				{error && <code style={{ color: 'red' }}>{JSON.stringify(error)}</code>}
			</section>
			<section style={{ padding: '10px', borderStyle: 'solid', marginBottom: '10px' }}>
				<button onClick={() => fetchUsers(10)}>retrieve users</button>
				{usersLoading && <div>loading...</div>}
				{users.length > 0 && (
					<ol>
						{users.map((user) => (
							<li key={user.userId}>
								{user.firstName} {user.lastName}
							</li>
						))}
					</ol>
				)}
				{usersError && <code style={{ color: 'red' }}>{JSON.stringify(usersError)}</code>}
			</section>
			<section style={{ padding: '10px', borderStyle: 'solid', marginBottom: '10px' }}>
				<button onClick={() => generateUsers(10)}>generate users</button>
				{genUsersLoading && <div>loading...</div>}
				{numCreated !== undefined && <p>generated {numCreated} users</p>}
				{genUsersError && <code style={{ color: 'red' }}>{JSON.stringify(genUsersError)}</code>}
			</section>
		</div>
	);
}

export default App;
