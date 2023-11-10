import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import AvatarInput from '.'

describe('AvatarInput', () => {
    global.URL.createObjectURL = jest.fn()
    it('should update the selected image container on image change', () => {
        render(<AvatarInput testId="avatar-input" />)

        // Encuentra el input de archivo por el atributo `data-testid`
        const fileInput = screen.getByTestId('avatar-input')

        console.log('rarara', fileInput)

        // Crea un archivo de prueba
        const mockFile = new File(['test file content'], 'example.jpg', { type: 'image/jpeg' })

        // Simula un evento de cambio con el archivo de prueba
        fireEvent.change(fileInput, { target: { files: [mockFile] } })

        // Verifica que el contenedor de la imagen seleccionada se actualiza
        const updatedImageContainer = screen.getByAltText('Selected Image')
        expect(updatedImageContainer).toBeInTheDocument()
    })
})
