import React from 'react'
import { render } from '@testing-library/react'
import WLogo from './titleSearch'

describe('WLogo Component', () => {
    it('renders WLogo component with specified properties', () => {
        // Define las propiedades que deseas probar
        const testProps = {
            alt: 'CustomAltText',
            size: 40,
        }

        // Renderiza el componente con las propiedades de prueba
        const { getByAltText } = render(<WLogo {...testProps} />)

        // Assertions para verificar que el componente se ha renderizado correctamente con las propiedades proporcionadas
        const imgElement = getByAltText('CustomAltText')
        expect(imgElement).toBeInTheDocument()
        expect(imgElement).toHaveAttribute('src', '/images/FrameStudentNET.png')
        expect(imgElement).toHaveStyle('width: 40px; height: 40px;')
    })

    it('renders WLogo component with default properties', () => {
        // Renderiza el componente sin proporcionar propiedades para probar los valores predeterminados
        const { getByAltText } = render(<WLogo />)

        // Assertions para verificar que el componente se ha renderizado correctamente con las propiedades predeterminadas
        const imgElement = getByAltText('FrameStudentNET')
        expect(imgElement).toBeInTheDocument()
        expect(imgElement).toHaveAttribute('src', '/images/FrameStudentNET.png')
        expect(imgElement).toHaveStyle('width: 30px; height: 30px;')
    })
})
