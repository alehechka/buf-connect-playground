import { User } from '../gen/users/v1/user_pb';
import { userClient } from '../client';
import { useCallback, useState } from 'react';
import useError from './userError';

const useUser = () => {
	const [loading, setLoading] = useState(false);
	const [user, setUser] = useState<User>();
	const [error, setError] = useError();

	const fetchUser = useCallback((userId?: string) => {
		setLoading(true);
		setUser(undefined);
		setError(undefined);
		return userClient
			.getUser({ userId })
			.then((res) => setUser(res.user))
			.catch((err) => setError(err))
			.finally(() => setLoading(false));
	}, []);

	return [user, loading, error, fetchUser] as const;
};

export default useUser;
