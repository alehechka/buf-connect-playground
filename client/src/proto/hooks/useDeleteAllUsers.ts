import { userClient } from '../client';
import { useCallback, useState } from 'react';
import useError from './useError';

const useDeleteAllUsers = () => {
	const [loading, setLoading] = useState(false);
	const [deleted, setDeleted] = useState<number>();
	const [error, setError] = useError();

	const deleteAllUsers = useCallback(() => {
		setLoading(true);
		setDeleted(undefined);
		setError(undefined);
		return userClient
			.deleteAllUsers({})
			.then((res) => setDeleted(res.numUsers))
			.catch((err) => setError(err))
			.finally(() => setLoading(false));
	}, []);

	return [deleted, loading, error, deleteAllUsers] as const;
};

export default useDeleteAllUsers;
