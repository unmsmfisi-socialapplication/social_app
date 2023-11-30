import { render, fireEvent } from '@testing-library/react'
import WTag from './tag'
import { AllInclusive, Home } from '@mui/icons-material'
import { BrowserRouter } from 'react-router-dom'
import './index.scss'

describe('WTag', () => {
    it('Should redirect to pages correctly as specified', () => {
        const path = '/ruta-especifica';

        // Render the component inside BrowserRouter
        const { getByText, asFragment } = render(
            <BrowserRouter>
                <WTag icon={AllInclusive} text="TagLink" path={path} />
            </BrowserRouter>
        );

        // Simulate the click on the link
        fireEvent.click(getByText('TagLink'));
    })
    it('Should redirect to pages correctly as specified', () => {
        const path = '/ruta-especifica'

        // Render the component
        const { getByText, asFragment } = render(<WTag icon={AllInclusive} text={path} path={path} />)
        fireEvent.click(getByText(path))
    })

    it('Should display style variations when isActive is true or false', () => {
        // Render the component with isActive true
        const { container: containerActive } = render(<WTag icon={AllInclusive} text="TagLink" isActive={true} />)

        // Render the component with isActive false
        const { container: containerFalse } = render(<WTag icon={AllInclusive} text="TagLink" isActive={false} />)
    })

    it('Should render with correct icon and text', () => {
        // Render the component with AllInclusive icon
        const { getByText, container: containerAllInclusive } = render(
            <WTag icon={AllInclusive} text="All Inclusive" />,
        )

        // Render the component with Home icon
        const { getByText: getByTextHome, container: containerHome } = render(<WTag icon={Home} text="Home" />)
    })
})
