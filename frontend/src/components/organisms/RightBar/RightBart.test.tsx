import React from 'react'
import { render, fireEvent } from '@testing-library/react'
import RightBar from './RightBar'

describe('RightBar Component', () => {
    it('renders RightBar component correctly', () => {
        const { getByText } = render(<RightBar dataTestid="confirmación" />)
        const salirButton = getByText('Salir')
        expect(salirButton).toBeInTheDocument()
    })

    it('calls handleLogout function when "Salir" button is clicked', () => {
        const consoleSpy = jest.spyOn(console, 'log')
        const { getByText } = render(<RightBar />)
        fireEvent.click(getByText('Salir'))
        expect(consoleSpy).toHaveBeenCalledWith('handleLogout called')
        consoleSpy.mockRestore()
    })
})
