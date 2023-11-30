import React from 'react'
import { render } from '@testing-library/react'
import '@testing-library/jest-dom'
import WCardFollow from './../../../components/atoms/CardFollow/CardFollow'

describe('CardFollow', () => {
    it('renders the CardFollow', () => {
        const { getByText } = render(
            <WCardFollow
                avatar={
                    'https://www.pngkey.com/png/full/114-1149878_setting-user-avatar-in-specific-size-without-breaking.png'
                }
                name="Wilfredo"
                userhandle="handle"
            />,
        )

        expect(getByText('Wilfredo')).toBeInTheDocument()
    })
})
