import React from 'react'
import { render, screen } from '@testing-library/react'
import WSpamProps from './Icon'
import PersonAddAlt1Icon from '@mui/icons-material/PersonAddAlt1'

describe('WSpamProps', () => {
    it('renders with default props', () => {
        render(<WSpamProps icon={PersonAddAlt1Icon} text="Default Text" />)

        expect(screen.getByText('Default Text')).toBeInTheDocument()
        expect(screen.getByTestId('PersonAddAlt1Icon')).toBeInTheDocument()
    })

    it('renders with custom props', () => {
        const CustomIcon = () => <div data-testid="custom-icon" />
        render(<WSpamProps typeColor="secondary" icon={CustomIcon} iconSize={30} text="Custom Text" />)

        expect(screen.getByTestId('custom-icon')).toBeInTheDocument()

        expect(screen.getByText('Custom Text')).toBeInTheDocument()
    })
})
