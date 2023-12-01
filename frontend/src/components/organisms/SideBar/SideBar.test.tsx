import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import '@testing-library/jest-dom'
import SideBar from './SideBar'

jest.mock('next/navigation', () => ({
    useRouter: () => ({
        route: '/',
        pathname: '',
        query: '',
        asPath: '',
        push: jest.fn(),
    }),
}))

test('renders Sidebar component', async () => {
    render(<SideBar />)

    // Verify that the component has been rendered correctly
    expect(screen.getByAltText('FrameStudentNET')).toBeInTheDocument()
    expect(screen.getByText('Home')).toBeInTheDocument()
    expect(screen.getByText('Explorer')).toBeInTheDocument()
    expect(screen.getByText('Notifications')).toBeInTheDocument()
    expect(screen.getByText('Messages')).toBeInTheDocument()
    expect(screen.getByText('Lists')).toBeInTheDocument()
    expect(screen.getByTestId('BookmarkBorderIcon')).toBeInTheDocument()
    expect(screen.getByText('Communities')).toBeInTheDocument()
    expect(screen.getByText('Post')).toBeInTheDocument()

    // Simulate a click on another link
    fireEvent.click(screen.getByText('Explorer'))
})
