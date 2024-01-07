import {
	CohereMessageFactory,
	OpenAIContentMessageFactory,
} from 'tests/factories';
import { getLocaleNamespace, render, screen } from 'tests/test-utils';
import { expect } from 'vitest';

import { PromptContentDisplay } from '@/components/prompt-display-components/prompt-content-display';
import { ModelVendor } from '@/types';

describe('PromptContentDisplay', () => {
	const namespace = getLocaleNamespace('createConfigWizard');

	it('should display an array of openAI messages correctly', () => {
		const messages = OpenAIContentMessageFactory.batchSync(10);

		render(
			<PromptContentDisplay
				messages={messages}
				modelVendor={ModelVendor.OpenAI}
			/>,
		);

		const messageContainer = screen.getByTestId(
			'prompt-content-display-messages',
		);
		expect(messageContainer).toBeInTheDocument();
		expect(messageContainer.children.length).toBe(10);

		for (const [i, message] of messages.entries()) {
			const content = `[${message.role} message]: ${message.content}`;
			expect(content).toBe(messageContainer.children[i].textContent);
		}
	});

	it('should display an array of cohere messages correctly', () => {
		const messages = CohereMessageFactory.batchSync(5);
		render(
			<PromptContentDisplay
				messages={messages}
				modelVendor={ModelVendor.Cohere}
			/>,
		);
		const messageContainer = screen.getByTestId(
			'prompt-content-display-messages',
		);
		expect(messageContainer).toBeInTheDocument();
		expect(messageContainer.children.length).toBe(5);

		for (const [i, message] of messages.entries()) {
			const content = `[${i}]: ${message.message}`;
			expect(content).toBe(messageContainer.children[i].textContent);
		}
	});

	it('should have the correct title', () => {
		const messages = CohereMessageFactory.batchSync(5);
		render(
			<PromptContentDisplay
				messages={messages}
				modelVendor={ModelVendor.Cohere}
			/>,
		);
		const title = screen.getByTestId('prompt-content-display-title');
		expect(title).toBeInTheDocument();
		expect(title.textContent).toBe(namespace.promptTemplate);
	});
});
