import { PropsWithChildren } from 'react';

const Box = ({ children }: PropsWithChildren) => {
	return <section style={{ padding: '10px', borderStyle: 'solid', marginBottom: '10px' }}>{children}</section>;
};

export default Box;
