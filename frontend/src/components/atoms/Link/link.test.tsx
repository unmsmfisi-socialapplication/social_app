import { render, fireEvent, screen } from '@testing-library/react'
import mockRouter from 'next-router-mock'
import { MemoryRouterProvider } from 'next-router-mock/MemoryRouterProvider'
import WBasicsLink from './link'

it('NextLink can be rendered', () => {
    render(<WBasicsLink dataTestid="link-test" href="/auth/login" text="Go to Login" />, {
        wrapper: MemoryRouterProvider,
    })
    const linkElement = screen.getByTestId('link-test')
    fireEvent.click(linkElement)
    expect(mockRouter.asPath).toEqual('/auth/login')
})