import { render, screen } from '@testing-library/react';
import WVerificationCode from './verificationCode';

it('renders WVerificationCode component with correct test IDs', () => {
  // Assuming dataTestid is 'defaultTestId'
  render(<WVerificationCode dataTestid="defaultTestId" />);
  
  // Ensure that the component renders without errors
  expect(screen.getByTestId('defaultTestId')).toBeInTheDocument();
  
});
