import { render, screen } from 'tests/test-utils';

import { Footer } from '@/components/static-site/footer';

describe('Footer', () => {
	it('should render a footer element', () => {
		render(<Footer />);

		const footerElement = screen.getByTestId('static-site-footer');
		expect(footerElement).toBeInTheDocument();
		const footerTextElement = screen.getByTestId('footer-text');
		const footerCopyrightElement = screen.getByTestId('footer-copyright');
		expect(footerTextElement).toBeInTheDocument();
		expect(footerCopyrightElement).toBeInTheDocument();
	});
});
