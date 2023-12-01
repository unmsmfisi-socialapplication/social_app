import React from 'react'
import { render } from '@testing-library/react'
import WChatPreview from './index'

describe('WChatPreview Component Tests', () => {
    // Prueba de renderizado básico
    test('renderiza el componente WChatPreview', () => {
        render(<WChatPreview avatar="test-avatar" name="John Doe" messagePreview="Mensaje de prueba" time="12:34" />)
    })

    // Prueba de renderizado personalizado con propiedades específicas y clases CSS
    test('renderiza WChatPreview con propiedades y clases personalizadas', () => {
        const propiedadesPersonalizadas = {
            avatar: 'custom-avatar-url',
            name: 'John Doe',
            messagePreview: '¡Hola, mundo!',
            time: '12:34',
        }

        const { container, getByText } = render(<WChatPreview {...propiedadesPersonalizadas} />)

        // Verifica si el contenedor principal tiene la clase CSS esperada
        expect(container.firstChild).toHaveClass('chatPreview')

        // Verifica si el avatar tiene la clase CSS esperada
        expect(container.querySelector('.avatar')).toHaveClass('avatar')

        // Verifica si la información tiene la clase CSS esperada
        expect(container.querySelector('.info')).toHaveClass('info')

        // Verifica si el nombre tiene la clase CSS esperada
        expect(container.querySelector('.name')).toHaveClass('name')

        // Verifica si el mensaje tiene la clase CSS esperada
        expect(container.querySelector('.message')).toHaveClass('message')

        // Verifica si el tiempo tiene la clase CSS esperada
        expect(container.querySelector('.time')).toHaveClass('time')

        // Verifica si las propiedades están renderizadas correctamente
        expect(getByText(propiedadesPersonalizadas.name)).toBeInTheDocument()
        expect(getByText(propiedadesPersonalizadas.messagePreview)).toBeInTheDocument()
        expect(getByText(propiedadesPersonalizadas.time)).toBeInTheDocument()
    })
})
