import React from 'react'
import { render, screen } from '@testing-library/react'
import SquareButton from './ButtonSquare'

describe('SquareButton Component', () => {
    test('renders SquareButton correctly', () => {
        render(<SquareButton />)
        const outerButton = screen.getByTestId('outer-button')
        expect(outerButton).toBeInTheDocument()
    })
})
