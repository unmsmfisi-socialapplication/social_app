import React from 'react'
import { render } from '@testing-library/react'
import WInput from './input'

describe('Input', () => {
    it('Renderiza el tipo de entrada password con un botón de visualización', () => {
        const { getByPlaceholderText } = render(
            <WInput typeColor="primary" fullWidth placeholder="Contraseña" type="password" />,
        )
        const input = getByPlaceholderText('Contraseña')
        expect(input.getAttribute('type')).toBe('password')
    })
    it('Verifica que el atributo error y el mensaje de error funcionen correctamente', () => {
        const errorMessage = 'Este es un mensaje de error'
        const { queryByText } = render(
            <WInput typeColor="primary" fullWidth error={true} type="password" errorMessage={errorMessage} />,
        )
        const errorText = queryByText(errorMessage)
        expect(errorText).toBeInTheDocument()
    })
})
