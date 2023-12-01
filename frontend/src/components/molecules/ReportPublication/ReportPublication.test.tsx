import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import ReportPublication from '.'

describe('ReportPublication', () => {
    it('renders the ReportPublication molecule component', () => {
        const { getByTestId } = render(
            <ReportPublication userHandle="Will0603" open={true} dataTestid="reportPublicationTest" />,
        )
        expect(getByTestId('reportPublicationTest')).toBeInTheDocument()
        expect(screen.getByTestId('reportPublicationTest')).toHaveTextContent(
            '¿Qué tipo de problema quieres denunciar?',
        )
    })
})
