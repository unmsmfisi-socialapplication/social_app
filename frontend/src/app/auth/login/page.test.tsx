// page.test.tsx
import React from 'react'
import { render, fireEvent, waitFor } from '@testing-library/react'
import LoginPage from './page' // Ajusta la ruta según tu estructura de archivos

// Mock de las dependencias que se utilizan en el componente (por ejemplo, Redux)
jest.mock('react-redux', () => ({
    ...jest.requireActual('react-redux'),
    useSelector: jest.fn(),
    useDispatch: jest.fn(),
}))

describe('LoginPage Component', () => {
    test('renders login page correctly', () => {
        const { getByText } = render(<LoginPage />)
        const welcomeText = getByText(/Bienvenido/i)
        expect(welcomeText).toBeInTheDocument()
    })

    test('handles form submission correctly', async () => {
        const { getByPlaceholderText, getByText } = render(<LoginPage />)

        // Simula la entrada del usuario
        const usernameInput = getByPlaceholderText('Ingrese su usuario')
        const passwordInput = getByPlaceholderText('Ingrese su contraseña')

        fireEvent.change(usernameInput, { target: { value: 'usuario_de_prueba' } })
        fireEvent.change(passwordInput, { target: { value: 'contraseña_de_prueba' } })

        // Simula la presentación del formulario
        const submitButton = getByText('Iniciar Sesión')
        fireEvent.click(submitButton)

        // Espera a que se maneje la lógica de inicio de sesión (puedes ajustar esto según tu lógica)
        await waitFor(() => {
            // Verifica que la lógica de inicio de sesión se haya manejado correctamente (puedes ajustar esto según tu lógica)
            // Por ejemplo, puedes verificar que el estado se haya actualizado correctamente después del inicio de sesión
        })
    })
})
