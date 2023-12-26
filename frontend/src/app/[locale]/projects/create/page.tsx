'use client';
import { useRouter } from 'next/navigation';
import { useTranslations } from 'next-intl';
import { useCallback, useEffect, useRef } from 'react';

import { Logo } from '@/components/logo';
import { CreateProjectForm } from '@/components/projects/create/create-project-form';
import { Navigation } from '@/constants';
import { useAnalytics } from '@/hooks/use-analytics';
import { useAuthenticatedUser } from '@/hooks/use-authenticated-user';
import { useProjects } from '@/stores/api-store';

export default function CreateProjectPage() {
	useAuthenticatedUser();

	const t = useTranslations('createProject');

	const router = useRouter();
	const { initialized, page } = useAnalytics();

	const projects = useProjects();

	const handleCancel = useCallback(() => {
		router.replace(`${Navigation.Projects}/${projects[0].id}`);
	}, [projects]);

	useEffect(() => {
		if (initialized) {
			page('create_project');
		}
	}, [initialized]);

	const isInitialRef = useRef(projects.length > 0);

	return (
		<div
			className="flex flex-col min-h-screen w-full bg-base-100"
			data-testid="create-projects-container"
		>
			<div className="page-content-container">
				<div
					data-testid="navbar-header"
					className="navbar flex-grow gap-4 content-baseline"
				>
					<Logo />
				</div>
				<div
					className="fixed top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 12/12  md:w-6/12 lg:w-5/12 2xl:w-4/12 bg-base-300 flex-col rounded"
					data-testid="create-project-view-flex-container"
				>
					<div
						className="pt-10 pb-6"
						data-testid="create-project-view-header"
					>
						<h1
							className="text-base-content text-center font-bold text-xl mb-2"
							data-testid="create-project-view-title"
						>
							{t('title')}
						</h1>
						<span
							className="text-center block text-sm text-neutral-content px-10"
							data-testid="create-project-view-sub-title"
						>
							{t('subTitle')}
						</span>
					</div>
					<CreateProjectForm
						allowCancel={!isInitialRef}
						handleCancel={handleCancel}
						validateApplicationName={(value) => !!value}
						validateProjectName={(value) => !!value}
					/>
				</div>
			</div>
		</div>
	);
}
