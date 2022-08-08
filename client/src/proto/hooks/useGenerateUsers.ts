import { userClient } from '../client';
import { useCallback, useState } from 'react';
import useError from './useError';

const useGenerateUsers = () => {
	const [loading, setLoading] = useState(false);
	const [numCreated, setNumCreated] = useState<number>();
	const [error, setError] = useError();

	const generateUsers = useCallback((numUsers?: number) => {
		setLoading(true);
		setNumCreated(undefined);
		setError(undefined);
		return userClient
			.generateUsers({ numUsers })
			.then((res) => setNumCreated(res.numUsers))
			.catch((err) => setError(err))
			.finally(() => setLoading(false));
	}, []);

	return [numCreated, loading, error, generateUsers] as const;
};

export default useGenerateUsers;
