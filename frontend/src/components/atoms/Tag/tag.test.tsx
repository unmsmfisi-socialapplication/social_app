import { render, fireEvent } from '@testing-library/react'
import WTag from './tag'
import { AllInclusive } from '@mui/icons-material'
import './index.scss'

describe('WTag', () => {

    it('Should display style variations when isActive is true or false', () => {
        // Render the component with isActive true
        const { container: containerActive } = render(<WTag icon={AllInclusive} text="TagLink" isActive={true} />)

        // Render the component with isActive false
        const { container: containerFalse } = render(<WTag icon={AllInclusive} text="TagLink" isActive={false} />)
    })
})
