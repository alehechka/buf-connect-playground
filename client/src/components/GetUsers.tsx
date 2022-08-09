import { useState } from 'react';
import useUsers from '../proto/hooks/useUsers';
import Box from './Box';

type Props = {};

const GetUsers = (props: Props) => {
	const [numUsers, setNumUsers] = useState<number>();
	const [users, usersLoading, usersError, fetchUsers, clearUsers] = useUsers(true);

	return (
		<Box>
			<input type='number' value={numUsers} onChange={(event) => setNumUsers(parseInt(event.target.value))}></input>
			<button onClick={() => fetchUsers(numUsers ? BigInt(numUsers) : 0n)}>retrieve users</button>
			<button onClick={clearUsers}>clear users</button>
			{usersLoading && <div>loading...</div>}
			{users.length > 0 && (
				<ol>
					{users.map((user) => (
						<li key={user.userId}>
							{user.firstName} {user.lastName}{' '}
							<ul>
								<li style={{ color: 'grey' }}>({user.userId})</li>
							</ul>
						</li>
					))}
				</ol>
			)}
			{usersError && (
				<div>
					<code style={{ color: 'red' }}>{JSON.stringify(usersError)}</code>
				</div>
			)}
		</Box>
	);
};

export default GetUsers;
