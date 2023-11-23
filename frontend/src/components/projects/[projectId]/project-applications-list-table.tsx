import { useRouter } from 'next/navigation';
import { useTranslations } from 'next-intl';
import { PencilFill } from 'react-bootstrap-icons';

import { applicationPageTabNames } from '@/app/[locale]/projects/[projectId]/applications/[applicationId]/page';
import { Navigation } from '@/constants';
import { Application, PromptConfig } from '@/types';
import { setRouteParams } from '@/utils/navigation';

export function ProjectApplicationsListTable({
	applications,
	projectId,
	promptConfigs,
}: {
	applications: Application[];
	projectId: string;
	promptConfigs: Record<string, PromptConfig<any>[] | undefined>;
}) {
	const t = useTranslations('projectOverview');
	const router = useRouter();

	return (
		<table className="custom-table mb-16">
			<thead>
				<tr>
					<th>{t('name')}</th>
					<th>{t('description')}</th>
					<th>{t('configs')}</th>
					<th>{t('edit')}</th>
				</tr>
			</thead>
			<tbody>
				{applications.map(
					({ name, description, id: applicationId }) => {
						return (
							<tr key={applicationId}>
								<td>
									<button
										data-testid="application-name-anchor"
										onClick={() => {
											router.push(
												setRouteParams(
													Navigation.ApplicationDetail,
													{
														applicationId,
														projectId,
													},
												),
											);
										}}
										className="btn-link"
									>
										<span>{name}</span>
									</button>
								</td>
								<td className="max-w-md">
									<span className="text-xs line-clamp-1 hover:line-clamp-none text-info">
										{description}
									</span>
								</td>
								<td
									data-testid={`application-prompt-config-count-${applicationId}`}
								>
									<span className="text-info">
										{promptConfigs[applicationId]?.length}
									</span>
								</td>
								<td>
									<button
										data-testid="application-edit-anchor"
										onClick={() => {
											router.push(
												setRouteParams(
													Navigation.ApplicationDetail,
													{
														applicationId,
														projectId,
													},
													applicationPageTabNames.SETTINGS,
												),
											);
										}}
									>
										<PencilFill className="w-3.5 h-3.5 text-secondary" />
									</button>
								</td>
							</tr>
						);
					},
				)}
			</tbody>
		</table>
	);
}
