import React from 'react'
import { render } from '@testing-library/react'
import WDetailsImage from './index'

test('Renderización y funcionamiento de la imagen', () => {
    // Render the component with custom properties
    const { getByAltText, getByText } = render(<WDetailsImage userName="Usuario de prueba" userHandle="prueba123" />)

    // Verify the image details properties.
    const accountNameElement = getByText(/Usuario de prueba/i)
    expect(accountNameElement).toBeInTheDocument()

    const nameElement = getByText(/@prueba123/i)
    expect(nameElement).toBeInTheDocument()

    const imageElement = getByAltText(/User Avatar/i)
    expect(imageElement).toBeInTheDocument()
})
