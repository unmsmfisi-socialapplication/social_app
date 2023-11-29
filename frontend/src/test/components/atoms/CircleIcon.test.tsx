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
})
