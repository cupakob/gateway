import { useTranslations } from 'next-intl';
import { useState } from 'react';
import { Front, KeyFill } from 'react-bootstrap-icons';

import { handleCreateAPIKey } from '@/api';
import { MIN_NAME_LENGTH } from '@/constants';
import { ApiError } from '@/errors';
import { useShowError, useShowSuccess } from '@/stores/toast-store';
import { copyToClipboard, handleChange } from '@/utils/helpers';

export function CreateApiKey({
	projectId,
	applicationId,
	onSubmit,
	onCancel,
	initialAPIKeyHash = '',
}: {
	applicationId: string;
	initialAPIKeyHash?: string;
	onCancel: () => void;
	onSubmit: () => void;
	projectId: string;
}) {
	const t = useTranslations('application');
	const showError = useShowError();
	const showSuccess = useShowSuccess();

	const [apiKeyName, setAPIKeyName] = useState('');
	const [apiKeyHash, setAPIKeyHash] = useState(initialAPIKeyHash);
	const [loading, setLoading] = useState(false);

	const apiKeyNameValid = apiKeyName.trim().length >= MIN_NAME_LENGTH;

	async function createAPIKey() {
		if (loading) {
			return;
		}

		try {
			setLoading(true);
			const apiKey = await handleCreateAPIKey({
				applicationId,
				data: { name: apiKeyName },
				projectId,
			});
			setAPIKeyHash(apiKey.hash);
		} catch (e) {
			showError((e as ApiError).message);
		} finally {
			setLoading(false);
		}
	}

	function close() {
		if (apiKeyHash) {
			onSubmit();
		} else {
			onCancel();
		}
		setAPIKeyName('');
		setAPIKeyHash('');
	}

	return (
		<div className="bg-base-300">
			<div className="p-10 flex flex-col items-center border-b border-neutral">
				<h1
					data-testid="create-api-key-title"
					className="text-base-content font-bold text-xl"
				>
					{t('createApiKey')}
				</h1>
				<p className="mt-2.5 font-medium ">
					{t('createApiKeyDescription')}
				</p>
				{!apiKeyHash && (
					<div className="mt-8 self-start w-full">
						<label
							htmlFor="create-api-key-input"
							className="text-sm font-semibold text-neutral-content"
						>
							{t('name')}
						</label>
						<input
							type="text"
							id="create-api-key-input"
							data-testid="create-api-key-input"
							className="input mt-2.5 bg-neutral w-full text-neutral-content font-medium"
							placeholder={t('createApiKeyPlaceholder')}
							value={apiKeyName}
							onChange={handleChange(setAPIKeyName)}
						/>
					</div>
				)}
				{apiKeyHash && (
					<div className="mt-8 self-start w-full">
						<label
							htmlFor="create-api-key-input"
							className="text-sm font-semibold text-neutral-content"
						>
							{t('apiKey')}
						</label>
						<div className="flex relative items-center gap-4 mt-2.5">
							<KeyFill className="w-4 h-4 text-neutral-content" />
							<input
								data-testid="create-api-key-hash-input"
								className="font-medium text-success bg-transparent w-full focus:border-none"
								value={apiKeyHash}
								onChange={() => {
									return;
								}}
							/>
							<button
								data-testid="api-key-copy-btn"
								onClick={() => {
									copyToClipboard(apiKeyHash);
									showSuccess(t('apiKeyCopied'));
								}}
							>
								<Front className="w-3.5 h-3.5 text-secondary" />
							</button>
						</div>
					</div>
				)}
			</div>
			<div className="flex items-center justify-end py-8 px-5 gap-4">
				<button
					data-testid="create-api-key-close-btn"
					onClick={close}
					className="btn btn-neutral capitalize font-semibold text-neutral-content"
				>
					{apiKeyHash ? t('close') : t('cancel')}
				</button>
				{!apiKeyHash && (
					<button
						data-testid="create-api-key-submit-btn"
						onClick={() => void createAPIKey()}
						disabled={!apiKeyNameValid}
						className={`btn btn-primary capitalize font-semibold ${
							apiKeyNameValid ? '' : 'opacity-60'
						}`}
					>
						{loading ? (
							<span className="loading loading-spinner loading-xs mx-1.5" />
						) : (
							t('create')
						)}
					</button>
				)}
			</div>
		</div>
	);
}
