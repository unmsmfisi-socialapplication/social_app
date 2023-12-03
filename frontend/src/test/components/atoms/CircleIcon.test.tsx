import React from 'react'
import { render } from '@testing-library/react'
import '@testing-library/jest-dom'
import WCircleIcon from './../../../components/atoms/CircleIcon/circleIcon'
import AllInclusiveIcon from '@mui/icons-material/AllInclusive'

describe('CircleIcon', () => {
    it('renders the CircleIcon', () => {
        const { getByTestId } = render(<WCircleIcon dataTestid="icon" icon={AllInclusiveIcon} />)

        expect(getByTestId('icon')).toBeInTheDocument()
    })

    it('applies custom styles to the WCircleIcon', () => {
        const customStyle = { color: 'red', fontSize: '24px' }
        const { getByTestId } = render(<WCircleIcon dataTestid="icon" icon={AllInclusiveIcon} style={customStyle} />)

        const iconElement = getByTestId('icon')
        expect(iconElement).toBeInTheDocument()
        expect(iconElement).toHaveStyle('color: red')
        expect(iconElement).toHaveStyle('font-size: 24px')
    })
})
