import { mockPage, mockReady, routerPushMock } from 'tests/mocks';
import { render, screen, waitFor } from 'tests/test-utils';

import LandingPage from '@/app/[locale]/page';
import { Navigation } from '@/constants';

describe('Landing Page', () => {
	it('should replace the route to Sign-In when clicking the Sign-Up button', () => {
		render(<LandingPage />);
		const button = screen.getByTestId('intro-section-signup-button');
		button.click();
		expect(routerPushMock).toHaveBeenCalledWith(
			Navigation.SignIn,
			Navigation.SignIn,
			{
				locale: undefined,
				scroll: true,
				shallow: undefined,
			},
		);
	});

	it('should replace the route to Sign-In when clicking on the Sign-In button', () => {
		render(<LandingPage />);
		const button = screen.getByTestId('header-sign-in-button');
		button.click();
		expect(routerPushMock).toHaveBeenCalledWith(
			Navigation.SignIn,
			Navigation.SignIn,
			{
				locale: undefined,
				scroll: true,
				shallow: undefined,
			},
		);
	});

	it('calls page tracking hook', async () => {
		render(<LandingPage />);
		await waitFor(() => {
			expect(mockReady).toHaveBeenCalled();
		});
		await waitFor(() => {
			expect(mockPage).toHaveBeenCalledWith(
				'landing-page',
				expect.any(Object),
			);
		});
	});
});
