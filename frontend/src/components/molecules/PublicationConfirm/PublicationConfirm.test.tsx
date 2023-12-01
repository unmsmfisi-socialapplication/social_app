import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import PublicationConfirm from '.'

describe('PublicationConfirm', () => {
    it('renders the PublicationConfirm molecule component', () => {
        const { getByTestId } = render(
            <PublicationConfirm
                avatarURL="https://images.ecestaticos.com/FjaDMYL1rpd8bqAVvR91YL-gZbY=/0x0:2252x1336/1200x1200/filters:fill(white):format(jpg)/f.elconfidencial.com%2Foriginal%2Fae2%2F47e%2F66d%2Fae247e66d9b8d8928d41a592b61690ca.jpg"
                userName="will"
                userHandle="Wilfredo"
                publicationContent="¡Felicidades!"
                open={true}
                dataTestid="publicationConfirmTest"
            />,
        )
        expect(getByTestId('publicationConfirmTest')).toBeInTheDocument()
        expect(screen.getByTestId('publicationConfirmTest')).toHaveTextContent('¡Felicidades!')
    })
})
