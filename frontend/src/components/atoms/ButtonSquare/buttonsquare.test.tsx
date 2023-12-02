import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import SquareButton from './ButtonSquare'

describe('SquareButton Component', () => {
    test('renders SquareButton correctly', () => {
        render(<SquareButton />)
        const outerButton = screen.getByTestId('outer-button')
        expect(outerButton).toBeInTheDocument()
    })

    test('executes onClick handler when clicked', () => {
        const onClickMock = jest.fn()
        render(<SquareButton onClick={onClickMock} />)
        const outerButton = screen.getByTestId('outer-button')

        fireEvent.click(outerButton)

        expect(onClickMock).toHaveBeenCalled()
    })
})
