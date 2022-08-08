import { useState } from 'react';

export type ProtoError = {
	name: string;
	rawMessage: string;
	code: number;
	metadata: Record<string, unknown>;
	details: string[];
};

const useError = (defaultError?: ProtoError) => useState<ProtoError | undefined>(defaultError);

export default useError;
