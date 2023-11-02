'use client';

import { useTranslations } from 'next-intl';
import { useState } from 'react';
import { Gear, KeyFill, Speedometer2 } from 'react-bootstrap-icons';

import { ApiKeys } from '@/components/projects/[projectId]/applications/[applicationId]/api-keys';
import { ApplicationAnalyticsPage } from '@/components/projects/[projectId]/applications/[applicationId]/application-analytics-page';
import { ApplicationDeletion } from '@/components/projects/[projectId]/applications/[applicationId]/application-deletion';
import { ApplicationGeneralSettings } from '@/components/projects/[projectId]/applications/[applicationId]/application-general-settings';
import { ApplicationPromptConfigs } from '@/components/projects/[projectId]/applications/[applicationId]/application-prompt-configs';
import { TabData, TabNavigation } from '@/components/tab-navigation';
import { useAuthenticatedUser } from '@/hooks/use-authenticated-user';
import { useProjectBootstrap } from '@/hooks/use-project-bootstrap';
import { useApplication } from '@/stores/project-store';

enum TAB_NAMES {
	OVERVIEW,
	SETTINGS,
	API_KEYS,
}

export default function Application({
	params: { projectId, applicationId },
}: {
	params: { projectId: string; applicationId: string };
}) {
	useAuthenticatedUser();
	useProjectBootstrap(false);
	const t = useTranslations('application');
	const application = useApplication(projectId, applicationId);

	const tabs: TabData<TAB_NAMES>[] = [
		{
			id: TAB_NAMES.OVERVIEW,
			text: t('overview'),
			icon: <Speedometer2 className="w-3.5 h-3.5" />,
		},
		{
			id: TAB_NAMES.SETTINGS,
			text: t('settings'),
			icon: <Gear className="w-3.5 h-3.5" />,
		},
		{
			id: TAB_NAMES.API_KEYS,
			text: t('apiKeys'),
			icon: <KeyFill className="w-3.5 h-3.5" />,
		},
	];
	const [selectedTab, setSelectedTab] = useState(TAB_NAMES.OVERVIEW);

	if (!application) {
		return null;
	}

	return (
		<div data-testid="application-page" className="my-8 mx-32">
			<h1
				data-testid="application-page-title"
				className="text-2xl font-semibold text-base-content"
			>
				{t('application')} / {application.name}
			</h1>
			<div className="mt-3.5 w-full mb-8">
				<TabNavigation<TAB_NAMES>
					tabs={tabs}
					selectedTab={selectedTab}
					onTabChange={setSelectedTab}
					trailingLine={true}
				/>
			</div>
			{selectedTab === TAB_NAMES.OVERVIEW && (
				<>
					<ApplicationAnalyticsPage
						applicationId={applicationId}
						projectId={projectId}
					/>
					<ApplicationPromptConfigs
						applicationId={applicationId}
						projectId={projectId}
					/>
				</>
			)}
			{selectedTab === TAB_NAMES.SETTINGS && (
				<>
					<ApplicationGeneralSettings
						applicationId={applicationId}
						projectId={projectId}
					/>
					<ApplicationDeletion
						applicationId={applicationId}
						projectId={projectId}
					/>
				</>
			)}
			{selectedTab === TAB_NAMES.API_KEYS && (
				<ApiKeys applicationId={applicationId} projectId={projectId} />
			)}
		</div>
	);
}
