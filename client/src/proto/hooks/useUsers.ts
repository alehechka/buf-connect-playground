import { ListUsersRequest, User } from '../gen/users/v1/user_pb';
import { userClient } from '../client';
import { useCallback, useState } from 'react';
import useError, { ProtoError } from './useError';

const useUsers = (append?: boolean) => {
	const [loading, setLoading] = useState(false);
	const [users, setUsers] = useState<User[]>([]);
	const [error, setError] = useError();
	const [page, setPage] = useState<bigint>(0n);
	const [prevCount, setPrevCount] = useState<number>(0);

	const fetchUsers = useCallback(
		async (numUsers: bigint) => {
			setLoading(true);
			if (!append) {
				setUsers([]);
			} else {
				setUsers((prev) => prev.slice(0, prev.length - prevCount));
			}
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
			if (count >= numUsers && numUsers > 0n) {
				setPage((prev) => prev + 1n);
				setPrevCount(0);
			} else {
				setPrevCount(count);
			}
		},
		[page, prevCount]
	);

	const clearUsers = useCallback(() => {
		setUsers([]);
		setPage(0n);
		setPrevCount(0);
	}, []);

	return [users, loading, error, fetchUsers, clearUsers] as const;
};

export default useUsers;
