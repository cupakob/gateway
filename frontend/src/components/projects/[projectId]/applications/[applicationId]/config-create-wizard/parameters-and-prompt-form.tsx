import { useTranslations } from 'next-intl';
import { useMemo, useState } from 'react';
import { ChevronDown, ChevronUp } from 'react-bootstrap-icons';

import {
	CohereModelParametersForm,
	CoherePromptTemplate,
} from '@/components/projects/[projectId]/applications/[applicationId]/config-create-wizard/cohere-form-components';
import {
	OpenAIModelParametersForm,
	OpenAIPromptTemplate,
} from '@/components/projects/[projectId]/applications/[applicationId]/config-create-wizard/openai-form-components';
import { Dimensions } from '@/constants';
import {
	CohereModelParameters,
	CohereModelType,
	CoherePromptMessage,
	ModelParameters,
	ModelType,
	ModelVendor,
	OpenAIContentMessage,
	OpenAIModelParameters,
	OpenAIModelType,
	ProviderMessageType,
} from '@/types';

export function PromptConfigParametersAndPromptForm<T extends ModelVendor>({
	modelVendor,
	modelType,
	messages,
	existingParameters,
	setParameters,
	setMessages,
}: {
	existingParameters?: ModelParameters<T>;
	messages?: ProviderMessageType<T>[];
	modelType: ModelType<T>;
	modelVendor: T;
	setMessages: (messages: ProviderMessageType<T>[]) => void;
	setParameters: (parameters: ModelParameters<T>) => void;
}) {
	const t = useTranslations('createConfigWizard');

	const [showAdvancedOptions, setShowAdvancedOptions] = useState(false);

	const parametersForm = useMemo(() => {
		if (modelVendor === ModelVendor.Cohere) {
			return (
				<CohereModelParametersForm
					modelType={modelType as CohereModelType}
					setParameters={setParameters}
					existingParameters={
						existingParameters as CohereModelParameters
					}
				/>
			);
		}
		return (
			<OpenAIModelParametersForm
				modelType={modelType as OpenAIModelType}
				setParameters={setParameters}
				existingParameters={existingParameters as OpenAIModelParameters}
			/>
		);
	}, [modelVendor]);

	const promptTemplateForm = useMemo(() => {
		if (modelVendor === ModelVendor.Cohere) {
			return (
				<CoherePromptTemplate
					messages={(messages ?? []) as CoherePromptMessage[]}
					setMessages={
						setMessages as (messages: CoherePromptMessage[]) => void
					}
				/>
			);
		}
		return (
			<OpenAIPromptTemplate
				messages={(messages ?? []) as OpenAIContentMessage[]}
				setMessages={
					setMessages as (messages: OpenAIContentMessage[]) => void
				}
			/>
		);
	}, [modelVendor, messages]);

	return (
		<div
			className="flex flex-col"
			data-testid="parameters-and-prompt-form-container"
		>
			<div>{promptTemplateForm}</div>
			<div className="card-divider" />
			<button
				className="flex w-fit justify-between items-center gap-2"
				onClick={() => {
					setShowAdvancedOptions(!showAdvancedOptions);
				}}
			>
				<span>{t('advancedOptions')}</span>
				{showAdvancedOptions ? (
					<ChevronUp
						height={Dimensions.Four}
						width={Dimensions.Four}
					/>
				) : (
					<ChevronDown
						height={Dimensions.Four}
						width={Dimensions.Four}
					/>
				)}
			</button>
			{showAdvancedOptions && <div>{parametersForm}</div>}
		</div>
	);
}
