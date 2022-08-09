import Box from './Box';
import useDeleteAllUsers from '../proto/hooks/useDeleteAllUsers';

type Props = {};

const DeleteAllUsers = (props: Props) => {
	const [deleted, loading, deleteAllError, deleteAllUsers] = useDeleteAllUsers();

	return (
		<Box>
			<button onClick={deleteAllUsers}>delete all users</button>
			{loading && <div>loading...</div>}
			{deleted !== undefined && <p>deleted {deleted} users</p>}
			{deleteAllError && (
				<div>
					<code style={{ color: 'red' }}>{JSON.stringify(deleteAllError)}</code>
				</div>
			)}
		</Box>
	);
};

export default DeleteAllUsers;
