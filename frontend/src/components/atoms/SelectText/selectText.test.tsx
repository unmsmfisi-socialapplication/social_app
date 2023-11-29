import React from 'react'
import { render, screen } from '@testing-library/react'
import WSelectedText from './selectText'

describe('WSelectedText Component', () => {
    const socialNetworkOptions = [
        { value: 'Twitter' },
        { value: 'Facebook' },
        { value: 'Instagram' },
        { value: 'Tik Tok' },
    ]

    it('renders with default props', () => {
        render(<WSelectedText options={[]} />)
        expect(screen.getByLabelText('Select')).toBeInTheDocument()
    })

    it('renders with a custom icon', () => {
        render(<WSelectedText options={socialNetworkOptions} iconComponent={() => <div data-testid="custom-icon" />} />)
        const customIcon = screen.getByTestId('custom-icon')
        expect(customIcon).toBeInTheDocument()
    })
})
