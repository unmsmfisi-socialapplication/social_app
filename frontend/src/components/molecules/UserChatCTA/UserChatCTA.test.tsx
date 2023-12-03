import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import WUserCHATCTA from '.'

describe('WUserCHATCTA', () => {
    it('renders the UserCHATCTA molecule component', () => {
        const { getByTestId } = render(
            <WUserCHATCTA
                avatarURL="https://images.ecestaticos.com/FjaDMYL1rpd8bqAVvR91YL-gZbY=/0x0:2252x1336/1200x1200/filters:fill(white):format(jpg)/f.elconfidencial.com%2Foriginal%2Fae2%2F47e%2F66d%2Fae247e66d9b8d8928d41a592b61690ca.jpg"
                userName="will0603"
                userHandle="will0603"
                dataTestid="userCHATCTATest"
            />,
        )
        expect(getByTestId('userCHATCTATest')).toBeInTheDocument()
        expect(screen.getByTestId('userCHATCTATest')).toHaveTextContent('will0603')
    })
})
