import { ListUsersRequest, User } from '../gen/users/v1/user_pb';
import { userClient } from '../client';
import { useCallback, useState } from 'react';
import useError, { ProtoError } from './userError';

const useUsers = (append?: boolean) => {
	const [loading, setLoading] = useState(false);
	const [users, setUsers] = useState<User[]>([]);
	const [error, setError] = useError();

	const fetchUsers = useCallback(async (numUsers: number) => {
		setLoading(true);
		if (!append) setUsers([]);
		setError(undefined);

		const request = new ListUsersRequest({ numUsers });
		try {
			for await (const response of userClient.listUsers(request)) {
				if (response.user) {
					setUsers((prev) => [...prev, response.user!]);
				}
			}
		} catch (error) {
			setError(error as ProtoError);
		}

		setLoading(false);
	}, []);

	return [users, loading, error, fetchUsers] as const;
};

export default useUsers;
