import React from 'react'
import { render, fireEvent } from '@testing-library/react'
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
    it('handles click events if onClick prop is provided', () => {
        const mockOnClick = jest.fn()
        const { getByText } = render(
            <WCardFollow
                avatar="https://www.pngkey.com/png/full/114-1149878_setting-user-avatar-in-specific-size-without-breaking.png"
                name="Wilfredo"
                userhandle="handle"
                onClick={mockOnClick}
            />,
        )

        const cardFollow = getByText('Wilfredo')
        fireEvent.click(cardFollow)

        expect(mockOnClick).toHaveBeenCalled()
    })
})
