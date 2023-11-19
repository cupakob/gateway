import { faker } from '@faker-js/faker';
import { fireEvent, waitFor } from '@testing-library/react';
import { ProjectUserAccountFactory } from 'tests/factories';
import { render, screen } from 'tests/test-utils';
import { beforeEach } from 'vitest';

import * as ProjectUsersAPI from '@/api/project-users-api';
import { InviteMember } from '@/components/projects/[projectId]/invite-member';
import { ApiError } from '@/errors';
import { ToastType } from '@/stores/toast-store';
import { AccessPermission } from '@/types';

describe('InviteMember', () => {
	const projectId = '1';
	const handleAddUserToProjectSpy = vi.spyOn(
		ProjectUsersAPI,
		'handleAddUsersToProject',
	);
	const handleRetrieveProjectUsersSpy = vi.spyOn(
		ProjectUsersAPI,
		'handleRetrieveProjectUsers',
	);

	beforeEach(() => {
		vi.resetAllMocks();
	});

	it('renders invite member', async () => {
		render(<InviteMember projectId={projectId} />);

		const emailInput = screen.getByTestId('invite-email-input');
		expect(emailInput).toBeInTheDocument();
		const permissionSelect = screen.getByTestId('permission-select');
		expect(permissionSelect).toBeInTheDocument();

		const sendInviteButton =
			screen.getByTestId<HTMLButtonElement>('send-invite-btn');
		expect(sendInviteButton.disabled).toBe(true);
	});

	test.each([
		[['plainaddress'], false],
		[['#@%^%#$@#$@#.com'], false],
		[['email.example.com'], false],
		[['.email@example.com'], false],
		[['a', 'b'], false],
		[['valid@gm.com', 'invalid@gm'], false],
		[['valid@gm.com', 'valid_again@gm.com'], true],
		[['valid@gm.com'], true],
		[['"email"@example.com'], true],
		[['_______@example.com'], true],
	])(
		'enables send invite only on valid email %p valid: %p',
		(emails, valid) => {
			render(<InviteMember projectId={projectId} />);

			const emailInput = screen.getByTestId('invite-email-input');
			emails.forEach((email) => {
				fireEvent.change(emailInput, {
					target: { value: email },
				});
				fireEvent.blur(emailInput);
			});

			const sendInviteButton =
				screen.getByTestId<HTMLButtonElement>('send-invite-btn');
			expect(sendInviteButton.disabled).toBe(!valid);
		},
	);

	it('changes role when a different role is chosen from dropdown', async () => {
		render(<InviteMember projectId={projectId} />);

		const permissionSelect =
			screen.getByTestId<HTMLSelectElement>('permission-select');
		expect(permissionSelect.value).toBe(AccessPermission.MEMBER);

		fireEvent.change(permissionSelect, {
			target: { value: AccessPermission.ADMIN },
		});

		await waitFor(() => {
			expect(permissionSelect.value).toBe(AccessPermission.ADMIN);
		});
	});

	it('sends invite to email and promptly clears the email input', async () => {
		render(<InviteMember projectId={projectId} />);

		const validEmail = faker.internet.email();
		const emailInput =
			screen.getByTestId<HTMLInputElement>('invite-email-input');
		fireEvent.change(emailInput, {
			target: { value: validEmail },
		});
		fireEvent.blur(emailInput);

		handleAddUserToProjectSpy.mockResolvedValueOnce(null);
		handleRetrieveProjectUsersSpy.mockResolvedValueOnce(
			ProjectUserAccountFactory.batchSync(2),
		);

		const sendInviteButton =
			screen.getByTestId<HTMLButtonElement>('send-invite-btn');
		fireEvent.click(sendInviteButton);

		expect(handleAddUserToProjectSpy).toHaveBeenCalledWith({
			data: [
				{
					email: validEmail,
					permission: AccessPermission.MEMBER,
				},
			],
			projectId,
		});
		await waitFor(() => {
			expect(emailInput.value).toBe('');
		});
	});

	it('debounce invite button when loading', async () => {
		render(<InviteMember projectId={projectId} />);

		const validEmail = faker.internet.email();
		const emailInput =
			screen.getByTestId<HTMLInputElement>('invite-email-input');
		fireEvent.change(emailInput, {
			target: { value: validEmail },
		});
		fireEvent.blur(emailInput);

		handleAddUserToProjectSpy.mockResolvedValueOnce(null);
		handleRetrieveProjectUsersSpy.mockResolvedValueOnce(
			ProjectUserAccountFactory.batchSync(2),
		);

		const sendInviteButton =
			screen.getByTestId<HTMLButtonElement>('send-invite-btn');
		fireEvent.click(sendInviteButton);
		fireEvent.click(sendInviteButton);

		expect(handleAddUserToProjectSpy).toHaveBeenCalledOnce();
	});

	it('throws error when user does not exist', async () => {
		render(<InviteMember projectId={projectId} />);

		const validEmail = faker.internet.email();
		const emailInput =
			screen.getByTestId<HTMLInputElement>('invite-email-input');
		fireEvent.change(emailInput, {
			target: { value: validEmail },
		});
		fireEvent.blur(emailInput);

		handleAddUserToProjectSpy.mockImplementationOnce(() => {
			throw new ApiError('user does not exist', {
				statusCode: 401,
				statusText: 'Bad Request',
			});
		});
		const sendInviteButton =
			screen.getByTestId<HTMLButtonElement>('send-invite-btn');
		fireEvent.click(sendInviteButton);

		const errorToast = screen.getByText('user does not exist');
		expect(errorToast.className).toContain(ToastType.ERROR);
	});

	it('removes duplicate emails on user input', async () => {
		render(<InviteMember projectId={projectId} />);

		const validEmail = faker.internet.email();
		const emailInput =
			screen.getByTestId<HTMLInputElement>('invite-email-input');

		fireEvent.change(emailInput, {
			target: { value: validEmail },
		});
		fireEvent.blur(emailInput);

		// add duplicate email
		fireEvent.change(emailInput, {
			target: { value: validEmail },
		});
		fireEvent.blur(emailInput);

		handleAddUserToProjectSpy.mockResolvedValueOnce(null);
		handleRetrieveProjectUsersSpy.mockResolvedValueOnce(
			ProjectUserAccountFactory.batchSync(2),
		);

		const sendInviteButton =
			screen.getByTestId<HTMLButtonElement>('send-invite-btn');
		fireEvent.click(sendInviteButton);

		expect(handleAddUserToProjectSpy).toHaveBeenCalledWith({
			data: [
				{
					email: validEmail,
					permission: AccessPermission.MEMBER,
				},
			],
			projectId,
		});
		await waitFor(() => {
			expect(emailInput.value).toBe('');
		});
	});

	it('disables submit when blank email is entered', async () => {
		render(<InviteMember projectId={projectId} />);

		const emailInput =
			screen.getByTestId<HTMLInputElement>('invite-email-input');

		fireEvent.change(emailInput, {
			target: { value: '' },
		});
		fireEvent.keyDown(emailInput, { code: 13, key: 'Enter' });

		const sendInviteButton =
			screen.getByTestId<HTMLButtonElement>('send-invite-btn');
		expect(sendInviteButton.disabled).toBe(true);
	});

	it('removes an email when clicked on cross button', async () => {
		render(<InviteMember projectId={projectId} />);

		const emailInput =
			screen.getByTestId<HTMLInputElement>('invite-email-input');

		const sendInviteButton =
			screen.getByTestId<HTMLButtonElement>('send-invite-btn');

		fireEvent.change(emailInput, {
			target: { value: faker.internet.email() },
		});
		fireEvent.blur(emailInput);

		expect(sendInviteButton.disabled).toBe(false);

		const removeEmailButton = screen.getByTestId('remove-email-btn');
		fireEvent.click(removeEmailButton);
		expect(sendInviteButton.disabled).toBe(true);
	});
});
