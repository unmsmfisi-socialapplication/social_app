import React, { useState } from 'react'
import { render } from '@testing-library/react'
import WTextArea from './textArea'

describe('TextArea', () => {
   
    it('Renderiza el componente atomo área de texto', () => {
        
        const { getByPlaceholderText } = render(
            <WTextArea placeholder="Escribe..."/>,
        )
        const textArea = getByPlaceholderText('Escribe...')
        expect(textArea.getAttribute('placeholder')).toBe('Escribe...')
    })

    it('Verifica el valor del componente atomo área de texto', () => {
        const value = "Hola"
        const { queryByText } = render(
            <WTextArea textAreaValue={value}/>,
        )
        const textArea = queryByText(value)
        expect(textArea).toBeInTheDocument()
    })
})