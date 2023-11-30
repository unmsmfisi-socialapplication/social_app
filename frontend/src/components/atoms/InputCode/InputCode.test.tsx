import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import WInputCode from './inputCode'

describe('WInputCode', () => {
    it('renders with default props', () => {
        render(<WInputCode />)

        expect(screen.getByRole('textbox')).toBeInTheDocument()
    })

    it('handles user input', () => {
        render(<WInputCode />)

        fireEvent.change(screen.getByRole('textbox'), { target: { value: '1' } })

        expect(screen.getByRole('textbox')).toHaveValue('1')
    })

    it('applies custom props', () => {
        render(<WInputCode typeColor="secondary" size="small" variant="filled" fullWidth={true} type="text" />)

        expect(screen.getByRole('textbox')).toBeInTheDocument()
    })
})
