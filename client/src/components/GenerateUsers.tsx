import { useState } from 'react';
import useGenerateUsers from '../proto/hooks/useGenerateUsers';
import Box from './Box';

type Props = {};

const GenerateUsers = (props: Props) => {
	const [numUsers, setNumUsers] = useState<number>();
	const [numCreated, genUsersLoading, genUsersError, generateUsers] = useGenerateUsers();

	return (
		<Box>
			<input type='number' value={numUsers} onChange={(event) => setNumUsers(parseInt(event.target.value))}></input>
			<button onClick={() => generateUsers(numUsers)}>generate users</button>
			{genUsersLoading && <div>loading...</div>}
			{numCreated !== undefined && <p>generated {numCreated} users</p>}
			{genUsersError && <code style={{ color: 'red' }}>{JSON.stringify(genUsersError)}</code>}
		</Box>
	);
};

export default GenerateUsers;
