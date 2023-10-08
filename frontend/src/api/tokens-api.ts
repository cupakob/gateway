import { fetcher } from '@/api/fetcher';
import { HttpMethod } from '@/constants';
import { Token, TokenCreateBody } from '@/types';

export async function handleCreateToken({
	applicationId,
	projectId,
	data,
}: {
	applicationId: string;
	projectId: string;
	data: TokenCreateBody;
}): Promise<Token> {
	return await fetcher<Token>({
		url: `projects/${projectId}/applications/${applicationId}/tokens/`,
		method: HttpMethod.Post,
		data,
	});
}

export async function handleDeleteToken({
	applicationId,
	projectId,
	tokenId,
}: {
	applicationId: string;
	projectId: string;
	tokenId: string;
}): Promise<void> {
	await fetcher<undefined>({
		url: `projects/${projectId}/applications/${applicationId}/tokens/${tokenId}/`,
		method: HttpMethod.Delete,
	});
}
