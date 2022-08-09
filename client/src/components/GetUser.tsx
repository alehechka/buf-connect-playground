import React, { useState } from 'react';
import useUser from '../proto/hooks/useUser';
import Box from './Box';

type Props = {};

const GetUser = (props: Props) => {
	const [userId, setUserID] = useState('');
	const [user, loading, error, fetchUser] = useUser();

	return (
		<Box>
			<input value={userId} onChange={(event) => setUserID(event.target.value)} />
			<button onClick={() => fetchUser(userId)}>get user</button>
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
			{error && (
				<div>
					<code style={{ color: 'red' }}>{JSON.stringify(error)}</code>
				</div>
			)}
		</Box>
	);
};

export default GetUser;
