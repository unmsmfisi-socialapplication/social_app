import React, { useState } from 'react'
import { render } from '@testing-library/react'
import WTextArea from './textArea'

describe('TextArea', () => {
   
    it('Renders the text area atom component', () => {
        
        const { getByPlaceholderText } = render(
            <WTextArea placeholder="Escribe..."/>,
        )
        const textArea = getByPlaceholderText('Escribe...')
        expect(textArea.getAttribute('placeholder')).toBe('Escribe...')
    })

    it('Check the value of the textarea atom component', () => {
        const value = "Hola"
        const { queryByText } = render(
            <WTextArea textAreaValue={value}/>,
        )
        const textArea = queryByText(value)
        expect(textArea).toBeInTheDocument()
    })
})