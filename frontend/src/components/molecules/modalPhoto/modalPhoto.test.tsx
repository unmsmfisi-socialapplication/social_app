import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import CustomDialog from './../../../components/molecules/modalPhoto/index';

describe('CustomDialog', () => {
    it('renders the CustomDialog component', async () => {
        render(
            <CustomDialog
                warning={true}
                title={{ textContent: 'Error', typeColor: 'error' }}
                subtitle="Subtitle Example"
                content="Content Example"
                size="xs"
                btnText="Confirm"
                dataTestid="customDialogTest"
            />,
        );

        expect(screen.getByTestId('customDialogTest')).toBeInTheDocument();

        expect(screen.queryByText('Error')).not.toBeInTheDocument();

        fireEvent.click(screen.getByTestId('customDialogTest'));

        await waitFor(() => {
            expect(screen.getByText('Error')).toBeInTheDocument();
        });

        expect(screen.getByText('Subtitle Example')).toBeInTheDocument();
        expect(screen.getByText('Content Example')).toBeInTheDocument();
        
    });
});
