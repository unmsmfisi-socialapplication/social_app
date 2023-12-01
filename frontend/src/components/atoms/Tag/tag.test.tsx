import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import AllInclusive from '@mui/icons-material/AllInclusive'
import { act } from 'react-dom/test-utils'
import WTag from '../../../components/atoms/Tag/tag'

// Mock of the next/router module
jest.mock('next/router', () => ({
    useRouter: jest.fn(),
}))
describe('WTag Component', () => {
    it('should navigate to the specified path on click', () => {
        const pushMock = jest.fn()
        // Configure useRouter to return the required object
        const useRouterMock = jest.spyOn(require('next/router'), 'useRouter')
        useRouterMock.mockReturnValue({ push: pushMock })

        render(<WTag icon={AllInclusive} text="TagLink" path="/some-path" />)

        act(() => {
            fireEvent.click(screen.getByText('TagLink'))
        })

        expect(pushMock).toHaveBeenCalledWith('/some-path')
    })

    it('should have active class when isActive is true', () => {
        render(<WTag icon={AllInclusive} text="TagLink" isActive={true} />)

        expect(screen.getByText('TagLink')).toHaveClass('tagLink--active')
    })

    it('should have default values when not provided', () => {
        render(<WTag />)

        expect(screen.getByText('TagLink')).toBeInTheDocument()
        expect(screen.getByText('TagLink')).not.toHaveClass('tagLink--active')
    })
})
