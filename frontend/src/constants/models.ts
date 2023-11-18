import {
	CohereModelsRecord,
	CohereModelType,
	ModelConfig,
	ModelType,
	ModelVendor,
	OpenAIModelParameters,
	OpenAIModelsRecord,
	OpenAIModelType,
} from '@/types';

const openAIIcon = 'openAI-icon.svg';
const cohereIcon = 'cohere-icon.svg';

const tempatureTooltip =
	'Controls randomness: Lowering results in less random completions. As the temperature approaches zero, the model will become deterministic and repetitive. Higher temperature results in more random completions.';

const topPTooltip =
	'Controls diversity via nucleus sampling: 0.5 means half of all likelihood-weighted options are considered. As the top_p approaches 1, the model will become deterministic and repetitive.';

const frequencyPenaltyTooltip =
	"How much to penalize new tokens based on their existing frequency in the text so far. Decreases the model's likelihood to repeat the same line verbatim.";

const presencePenaltyTooltip =
	'How much to penalize new tokens based on whether they appear in the text so far. Increases the model’s likelihood to talk about new topics.';

const maxTokensTooltip =
	'The maximum number of tokens to generate. Requests can use up to 2048 tokens shared between prompt and response.';

export const OPEN_AI_MAX_TOKENS = 2048;

export const vendorImageSourceMap: Record<ModelVendor, string> = {
	[ModelVendor.OpenAI]: `/images/${openAIIcon}`,
	[ModelVendor.Cohere]: `/images/${cohereIcon}`,
};

export const modelTypeToNameMap: Record<ModelType<any>, string> = {
	[OpenAIModelType.Gpt35Turbo]: 'GPT-3.5 Turbo',
	[OpenAIModelType.Gpt3516K]: 'GPT-3.5 Turbo 16K',
	[OpenAIModelType.Gpt4]: 'GPT-4',
	[OpenAIModelType.Gpt432K]: 'GPT-4 32K',
	[CohereModelType.Command]: 'Command',
	[CohereModelType.CommandLight]: 'Command Light',
	[CohereModelType.CommandNightly]: 'Command Nightly',
	[CohereModelType.CommandLightNightly]: 'Command Light Nightly',
};

export const modelTypeToMaxTokensMap: Record<ModelType<any>, number> = {
	[OpenAIModelType.Gpt35Turbo]: 4096,
	[OpenAIModelType.Gpt3516K]: 16_384,
	[OpenAIModelType.Gpt4]: 8192,
	[OpenAIModelType.Gpt432K]: 32_768,
	[CohereModelType.Command]: 4096,
	[CohereModelType.CommandLight]: 4096,
	[CohereModelType.CommandNightly]: 4096,
	[CohereModelType.CommandLightNightly]: 4096,
};

const createOpenAIModelParameter = (
	maxTokens: number,
): ModelConfig<Required<OpenAIModelParameters>>['parameters'] => ({
	frequencyPenalty: { max: 2, min: 0, step: 0.1 },
	maxTokens: { max: maxTokens, min: 2, step: 10 },
	presencePenalty: { max: 2, min: 0, step: 0.1 },
	temperature: { max: 1, min: 0, step: 0.1 },
	topP: { max: 1, min: 0, step: 0.1 },
});

export const openAIModels: OpenAIModelsRecord = {
	[OpenAIModelType.Gpt35Turbo]: {
		icon: openAIIcon,
		name: 'GPT-3.5',
		parameters: createOpenAIModelParameter(4096),
	},
	[OpenAIModelType.Gpt3516K]: {
		icon: openAIIcon,
		name: 'GPT-3.5 16K',
		parameters: createOpenAIModelParameter(16_384),
	},
	[OpenAIModelType.Gpt4]: {
		icon: openAIIcon,
		name: 'GPT-4',
		parameters: createOpenAIModelParameter(8192),
	},
	[OpenAIModelType.Gpt432K]: {
		icon: openAIIcon,
		name: 'GPT-4 32K',
		parameters: createOpenAIModelParameter(32_768),
	},
};

export const cohereModels: CohereModelsRecord = {
	[CohereModelType.Command]: {
		icon: cohereIcon,
		name: 'Command',
		parameters: {
			temperature: { max: 1, min: 0, step: 0.1 },
		},
	},
	[CohereModelType.CommandLight]: {
		icon: cohereIcon,
		name: 'Command Light',
		parameters: {
			temperature: { max: 1, min: 0, step: 0.1 },
		},
	},
	[CohereModelType.CommandNightly]: {
		icon: cohereIcon,
		name: 'Command Nightly',
		parameters: {
			temperature: { max: 1, min: 0, step: 0.1 },
		},
	},
	[CohereModelType.CommandLightNightly]: {
		icon: cohereIcon,
		name: 'Command Light Nightly',
		parameters: {
			temperature: { max: 1, min: 0, step: 0.1 },
		},
	},
};

export const ToolTipText = {
	frequencyPenalty: frequencyPenaltyTooltip,
	maxTokens: maxTokensTooltip,
	presencePenalty: presencePenaltyTooltip,
	temperature: tempatureTooltip,
	topP: topPTooltip,
};
