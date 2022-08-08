import { ListUsersRequest, User } from '../gen/users/v1/user_pb';
import { userClient } from '../client';
import { useCallback, useState } from 'react';
import useError, { ProtoError } from './useError';

const useUsers = (append?: boolean) => {
	const [loading, setLoading] = useState(false);
	const [users, setUsers] = useState<User[]>([]);
	const [error, setError] = useError();
	const [page, setPage] = useState<bigint>(0n);
	const [disabled, setDisabled] = useState<boolean>(false);

	const fetchUsers = useCallback(
		async (numUsers: bigint) => {
			setLoading(true);
			if (!append) setUsers([]);
			setError(undefined);

			const request = new ListUsersRequest({ numUsers: BigInt(numUsers), page });
			let count = 0;
			try {
				for await (const response of userClient.listUsers(request)) {
					if (response.user) {
						count++;
						setUsers((prev) => [...prev, response.user!]);
					}
				}
			} catch (error) {
				setError(error as ProtoError);
			}

			setLoading(false);
			if (count >= numUsers) {
				setPage((prev) => prev + 1n);
			} else {
				setDisabled(true);
			}
		},
		[page]
	);

	return [users, loading, error, fetchUsers, disabled] as const;
};

export default useUsers;
