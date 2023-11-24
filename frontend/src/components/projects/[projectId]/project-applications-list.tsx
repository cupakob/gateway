import { useTranslations } from 'next-intl';
import { useRef } from 'react';
import { Plus } from 'react-bootstrap-icons';
import useSWR from 'swr';

import { handleRetrieveApplications, handleRetrievePromptConfigs } from '@/api';
import { CreateApplication } from '@/components/projects/[projectId]/applications/create-application';
import { ProjectApplicationsListTable } from '@/components/projects/[projectId]/project-applications-list-table';
import { ApiError } from '@/errors';
import {
	useApplications,
	usePromptConfigs,
	useSetProjectApplications,
	useSetPromptConfigs,
} from '@/stores/api-store';
import { useShowError } from '@/stores/toast-store';

export function ProjectApplicationsList({ projectId }: { projectId: string }) {
	const t = useTranslations('projectOverview');

	const applications = useApplications(projectId);
	const setProjectApplications = useSetProjectApplications();

	const promptConfigs = usePromptConfigs();
	const setPromptConfig = useSetPromptConfigs();

	const dialogRef = useRef<HTMLDialogElement>(null);

	const showError = useShowError();

	const { isLoading } = useSWR(projectId, handleRetrieveApplications, {
		onError({ message }: ApiError) {
			showError(message);
		},
		onSuccess(data) {
			setProjectApplications(projectId, data);
		},
	});

	useSWR(
		() => applications,
		(applications) =>
			Promise.all(
				applications.map((application) =>
					handleRetrievePromptConfigs({
						applicationId: application.id,
						projectId,
					}),
				),
			),
		{
			onSuccess(data) {
				data.forEach((promptConfig, index) => {
					setPromptConfig(applications![index].id, promptConfig);
				});
			},
		},
	);

	const openAppCreateFlow = () => {
		dialogRef.current?.showModal();
	};

	const closeAppCreateFlow = () => {
		dialogRef.current?.close();
	};

	return (
		<div data-testid="project-application-list-container" className="mt-9">
			<h2 className="card-header">{t('applications')}</h2>
			<div className="rounded-data-card flex flex-col">
				{applications?.length ? (
					<ProjectApplicationsListTable
						projectId={projectId}
						applications={applications}
						promptConfigs={promptConfigs}
					/>
				) : isLoading ? (
					<div className="w-full flex mb-8">
						<span className="loading loading-bars mx-auto" />
					</div>
				) : null}
				<button
					data-testid="new-application-btn"
					onClick={openAppCreateFlow}
					className="flex gap-2 items-center text-secondary hover:brightness-90"
				>
					<Plus className="text-secondary w-4 h-4 hover:brightness-90" />
					<span>{t('newApplication')}</span>
				</button>
			</div>
			<dialog ref={dialogRef} className="modal">
				<div className="dialog-box border-0 rounded-none">
					<CreateApplication
						onClose={closeAppCreateFlow}
						projectId={projectId}
					/>
				</div>
				<form method="dialog" className="modal-backdrop">
					<button />
				</form>
			</dialog>
		</div>
	);
}
