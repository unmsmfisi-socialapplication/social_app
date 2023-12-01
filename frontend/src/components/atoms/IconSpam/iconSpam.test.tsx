import React from 'react'
import PersonAddAlt1Icon from '@mui/icons-material/PersonAddAlt1'
import { render } from '@testing-library/react'
import WSpamProps from './Icon'

test('renders WSpamProps component with specified properties', () => {
    // Define the properties you want to test
    // Render the component with the test properties
    const { getByText, getByTestId } = render(
        <WSpamProps
            testId="wspamprops-test"
            typeColor="primary"
            icon={PersonAddAlt1Icon}
            iconSize={20}
            text="IconSpamTest"
        />,
    )

    // Assertions to check if the component rendered with the specified properties
    const svgIconElement = getByTestId('wspamprops-test').querySelector('svg')
    expect(svgIconElement).toBeInTheDocument()
    expect(svgIconElement).toHaveClass('SpamProps--primary')
    const textElement = getByText('IconSpamTest')
    expect(textElement).toBeInTheDocument()
})
