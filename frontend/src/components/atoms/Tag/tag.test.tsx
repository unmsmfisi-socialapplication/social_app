import { render, fireEvent, screen } from '@testing-library/react'
import WTag from './tag'
import { AllInclusive } from '@mui/icons-material'
import './index.scss'
import mockRouter from 'next-router-mock';
import { MemoryRouterProvider } from 'next-router-mock/MemoryRouterProvider';

describe('WTag', () => {
    it('Should redirect to pages correctly as specified', () => {
        const path = '/ruta-especifica'

        // Render the component
        const { getByText, asFragment } = render(<WTag icon={AllInclusive} text={path} path={path} />)
        fireEvent.click(getByText(path))

        const componentSnapshot = asFragment()
    })

    it('Should display style variations when isActive is true or false', () => {
        // Render the component with isActive true
        const { container: containerActive } = render(<WTag icon={AllInclusive} text="TagLink" isActive={true} />)

        // Render the component with isActive false
        const { container: containerFalse } = render(<WTag icon={AllInclusive} text="TagLink" isActive={false} />)
    })

    it('NextLink can be rendered', () => {
        render(<WTag icon={AllInclusive} dataTestId="link-test" path="/auth/login" text="Go to Login" />, {
            wrapper: MemoryRouterProvider,
        })
        const linkElement = screen.getByTestId('link-test')
        fireEvent.click(linkElement)
        expect(mockRouter.asPath).toEqual('/auth/login')
    })

})
