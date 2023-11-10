// page.test.tsx
import React from 'react'
import { render, fireEvent, waitFor } from '@testing-library/react'
import LoginPage from './page' 

jest.mock('react-redux', () => ({
    ...jest.requireActual('react-redux'),
    useSelector: jest.fn(),
    useDispatch: jest.fn(),
}))

describe('LoginPage Component', () => {
    test('renders login page correctly', () => {
        const { getByText } = render(<LoginPage />)
        const welcomeText = getByText(/Bienvenido/i)
    })

    test('handles form submission correctly', async () => {
        const { getByPlaceholderText, getByText } = render(<LoginPage />)

        const usernameInput = getByPlaceholderText('Ingrese su usuario')
        const passwordInput = getByPlaceholderText('Ingrese su contraseña')

        fireEvent.change(usernameInput, { target: { value: 'usuario_de_prueba' } })
        fireEvent.change(passwordInput, { target: { value: 'contraseña_de_prueba' } })

        const submitButton = getByText('Iniciar Sesión')
        fireEvent.click(submitButton)
        await waitFor(() => {
        })
    })
})
