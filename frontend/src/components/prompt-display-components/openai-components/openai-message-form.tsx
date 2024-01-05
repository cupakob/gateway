import { useTranslations } from 'next-intl';
import React from 'react';
import { ArrowDown, ArrowUp, XCircle } from 'react-bootstrap-icons';

import { Dimensions } from '@/constants';
import { openAIRoleColorMap } from '@/constants/models';
import { OpenAIContentMessageRole, OpenAIPromptMessageRole } from '@/types';
import { handleChange } from '@/utils/events';

export function OpenAIMessageForm({
	id,
	role,
	setRole,
	name,
	setName,
	content,
	setContent,
	handleDeleteMessage,
	handleArrowUp,
	handleArrowDown,
	showArrowUp,
	showArrowDown,
}: {
	content: string;
	handleArrowDown: () => void;
	handleArrowUp: () => void;
	handleDeleteMessage: (id: string) => void;
	id: string;
	name?: string;
	role: OpenAIContentMessageRole;
	setContent: (content: string) => void;
	setName: (name: string) => void;
	setRole: (role: OpenAIContentMessageRole) => void;
	showArrowDown: boolean;
	showArrowUp: boolean;
}) {
	const t = useTranslations('openaiPromptTemplate');

	const roleColor = `text-${openAIRoleColorMap[role]}`;

	return (
		<div
			data-testid={`openai-message-container-${id}`}
			className="rounded-dark-card flex cursor-grab"
		>
			<div
				className={`flex flex-col justify-${
					showArrowUp ? 'between' : 'end'
				} items-center pr-2`}
			>
				{showArrowUp && (
					<button
						data-testid={`openai-message-arrow-up-${id}`}
						className="btn btn-ghost p-0 text-accent/70 hover:text-accent ${showArrowUp ?'' : 'hidden'}}"
						onClick={handleArrowUp}
					>
						<ArrowUp height={Dimensions.Five} />
					</button>
				)}
				{showArrowDown && (
					<button
						data-testid={`openai-message-arrow-down-${id}`}
						className="btn btn-ghost p-0 text-accent/70 hover:text-accent "
						onClick={handleArrowDown}
					>
						<ArrowDown height={Dimensions.Five} />
					</button>
				)}
			</div>
			<div
				className={`grow p-3 border-2 rounded border-neutral border-opacity-50 border-double hover:border-opacity-100 flex gap-4`}
			>
				<div className="flex flex-col gap-2">
					<div className="form-control" data-no-dnd="true">
						<label
							className="label"
							data-testid={`openai-message-role-select-label-${id}`}
						>
							<span className="label-text">
								{t('messageRole')}
							</span>
						</label>
						<select
							className={`select select-bordered select-sm rounded bg-neutral ${roleColor} text-xs`}
							value={role}
							onChange={handleChange(setRole)}
							data-testid={`openai-message-role-select-${id}`}
						>
							{Object.entries(OpenAIPromptMessageRole)
								.filter(
									(v) =>
										v[1] !==
										OpenAIPromptMessageRole.Function,
								)
								.map(([key, value]) => (
									<option
										key={value}
										id={value}
										value={value}
									>
										{`${key} ${t('message')}`}
									</option>
								))}
						</select>
					</div>
					<div className="form-control" data-no-dnd="true">
						<label
							className="label"
							data-testid={`openai-message-name-input-label-${id}`}
						>
							<span className="label-text">
								{t('messageName')}
							</span>
							<span className="label-text-alt">
								{t('optional')}
							</span>
						</label>
						<input
							type="text"
							placeholder={t('messageNameInputPlaceholder')}
							className="input input-bordered input-sm rounded bg-neutral text-neutral-content text-xs"
							data-testid={`openai-message-name-input-${id}`}
							value={name ?? ''}
							onChange={handleChange(setName)}
						/>
					</div>
				</div>
				<div
					className="form-control grow min-h-full"
					data-no-dnd="true"
				>
					<label
						className="label"
						data-testid={`openai-message-content-textarea-label-${id}`}
					>
						<span className="label-text">
							{t('messageContent')}
						</span>
						<span className="text-info label-text-alt text-sm">
							{t('wrapVariable')}
						</span>
					</label>
					<textarea
						className="textarea rounded bg-neutral text-neutral-content h-full min-h-fit"
						placeholder={t('messageContentPlaceholder')}
						value={content}
						onChange={handleChange(setContent)}
						data-testid={`openai-message-content-textarea-${id}`}
					/>
				</div>
			</div>
			<button
				className="btn btn-ghost self-center text-warning hover:text-error"
				data-testid={`openai-message-delete-button-${id}`}
				onClick={() => {
					handleDeleteMessage(id);
				}}
			>
				<XCircle height={Dimensions.Four} width={Dimensions.Four} />
			</button>
		</div>
	);
}
