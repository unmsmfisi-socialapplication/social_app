import { render } from '@testing-library/react'
import Button from './button'
import './index.scss'
describe('Button', () => {
    it('renders the button', () => {
        const { getByTestId } = render(<Button dataTestid="button" text="Hola mundo" />)

        expect(getByTestId('button')).toHaveTextContent('Hola mundo')
    })
    it('renders the disabled button', () => {
        const { getByTestId } = render(<Button dataTestid="button" text="Hola mundo" disabled />)

        expect(getByTestId('button')).toBeDisabled()
    })
})
