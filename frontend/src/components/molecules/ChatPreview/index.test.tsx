import React from 'react'
import { render } from '@testing-library/react'
import WChatPreview from './index'

describe('WChatPreview Component Tests', () => {
    // Prueba básica de renderizado
    test('renderiza el componente WChatPreview', () => {
        render(<WChatPreview avatar="test-avatar" name="John Doe" messagePreview="Mensaje de prueba" time="12:34" />)
    })

    // Prueba de renderizado personalizado con propiedades específicas
    test('renderiza WChatPreview con propiedades personalizadas', () => {
        const propiedadesPersonalizadas = {
            avatar: 'custom-avatar-url',
            name: 'John Doe',
            messagePreview: '¡Hola, mundo!',
            time: '12:34',
        }

        const { getByText } = render(<WChatPreview {...propiedadesPersonalizadas} />)

        // Verifica si el contenido de texto específico está presente en el componente
        expect(getByText(propiedadesPersonalizadas.name)).toBeInTheDocument()
        expect(getByText(propiedadesPersonalizadas.messagePreview)).toBeInTheDocument()
        expect(getByText(propiedadesPersonalizadas.time)).toBeInTheDocument()
    })

    // Prueba de snapshot
    test('coincide con el snapshot', () => {
        const { asFragment } = render(
            <WChatPreview avatar="test-avatar" name="John Doe" messagePreview="Mensaje de prueba" time="12:34" />,
        )
        expect(asFragment()).toMatchSnapshot()
    })

    // Prueba de clases CSS
    test('contiene las clases CSS esperadas', () => {
        const propiedadesPersonalizadas = {
            avatar: 'custom-avatar-url',
            name: 'John Doe',
            messagePreview: '¡Hola, mundo!',
            time: '12:34',
        }

        const { container } = render(<WChatPreview {...propiedadesPersonalizadas} />)

        // Verifica si el contenedor principal tiene la clase CSS esperada
        expect(container.firstChild).toHaveClass('chatPreview')

        // Verifica si el avatar tiene la clase CSS esperada
        expect(container.querySelector('.avatar')).toHaveClass('avatar')
    })
})
