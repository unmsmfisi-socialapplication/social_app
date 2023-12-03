import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import WButtonPost from './ButtonPost'

describe('Button', () => {
    it('renders the button', () => {
        const { getByTestId } = render(<WButtonPost dataTestid="button" text="ButtonPost" />)
        expect(getByTestId('button')).toHaveTextContent('ButtonPost')
    })

    it('renders the disabled button', () => {
        const { getByTestId } = render(<WButtonPost dataTestid="button" text="ButtonPost" disabled />)

        expect(getByTestId('button')).toBeDisabled()
    })
    it('renders a large button', () => {
        const { getByTestId } = render(<WButtonPost dataTestid="button" text="Large Button" size="large" />)
        const button = getByTestId('button')
        expect(button).toHaveTextContent('Large Button')
        expect(button).toHaveStyle('min-width: 100%')
    })

    it('renders an outlined button', () => {
        const { getByTestId } = render(<WButtonPost dataTestid="button" text="Outlined Button" variant="outlined" />)
        const button = getByTestId('button')
        expect(button).toHaveTextContent('Outlined Button')
        expect(button).toHaveClass('MuiButton-outlined')
    })

    it('handles click event', () => {
        const handleClick = jest.fn()
        render(<WButtonPost onClick={handleClick} />)

        fireEvent.click(screen.getByRole('button'))
        expect(handleClick).toHaveBeenCalledTimes(1)
    })
})
