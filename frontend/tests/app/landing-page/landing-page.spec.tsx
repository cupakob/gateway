import { routerPushMock } from 'tests/mocks';
import { render, screen } from 'tests/test-utils';

import LandingPage from '@/app/[locale]/page';
import { Navigation } from '@/constants';

describe('Landing Page', () => {
	it('should render all sections', () => {
		render(<LandingPage />);

		const featureCards = screen.getAllByTestId(/-section$/);
		expect(featureCards).toHaveLength(3);
	});

	it('should replace the route to Sign-In when clicking the Sign-Up button', () => {
		render(<LandingPage />);
		const button = screen.getByTestId('intro-section-signup-button');
		button.click();
		expect(routerPushMock).toHaveBeenCalledWith(Navigation.SignIn);
	});

	it('should replace the route to Sign-In when clicking on the Sign-In button', () => {
		render(<LandingPage />);
		const button = screen.getByTestId('header-sign-in-button');
		button.click();
		expect(routerPushMock).toHaveBeenCalledWith(Navigation.SignIn);
	});
});
