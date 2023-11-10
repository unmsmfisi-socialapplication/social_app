import React from 'react'
import { render } from '@testing-library/react'
import WDetailsImage from './index'

test('Renderización y funcionamiento de la imagen', () => {
    // Render the component with custom properties
    const { getByAltText, getByText } = render(
        <WDetailsImage
            accountName="Usuario de prueba"
            name="prueba123"
            icon={() => (
                <img
                    alt="Imagen de prueba"
                    src="https://images.unsplash.com/photo-1606425270259-998c37268501?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1758&q=80"
                />
            )}
        />,
    )

    // Verify the image details properties.
    const accountNameElement = getByText(/Usuario de prueba/i)
    expect(accountNameElement).toBeInTheDocument()

    const nameElement = getByText(/Posteado por prueba123/i)
    expect(nameElement).toBeInTheDocument()

    const imageElement = getByAltText(/Imagen de prueba/i)
    expect(imageElement).toBeInTheDocument()
})
