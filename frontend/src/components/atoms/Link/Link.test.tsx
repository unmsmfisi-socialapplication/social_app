import React from 'react'
import { render, screen } from '@testing-library/react'
import WBasicsLink from './link'

describe('WBasicsLink', () => {
    it('renders the link with default props', () => {
        render(<WBasicsLink />)
        const linkElement = screen.getByText(/link/i)
        expect(linkElement).toBeInTheDocument()
        expect(linkElement).toHaveAttribute('href', '#')
        expect(linkElement).toHaveStyle({ textDecoration: 'underline' })
    })

    it('renders the link with custom props', () => {
        render(<WBasicsLink underline="none" text="Custom Link" href="https://example.com" displayType="inline-flex" />)

        const linkElement = screen.getByText(/custom link/i)
        expect(linkElement).toBeInTheDocument()
        expect(linkElement).toHaveAttribute('href', 'https://example.com')
        expect(linkElement).toHaveStyle({ textDecoration: 'none' })
    })
})
