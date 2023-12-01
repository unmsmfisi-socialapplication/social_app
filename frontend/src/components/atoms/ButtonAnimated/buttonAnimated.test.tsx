import WButtonMotion from './buttonAnimated'
import { render, fireEvent } from '@testing-library/react'

describe('ButtonAnimated', () => {
    it('renders the ButtonAnimated with the correct text', () => {
        const { getByTestId } = render(<WButtonMotion dataTestid="buttonAnimated" text="Go" />)

        expect(getByTestId('buttonAnimated')).toHaveTextContent('Go')
    })

    it('renders the disabled ButtonAnimated', () => {
        const { getByTestId } = render(<WButtonMotion dataTestid="buttonAnimated" text="Go" disabled={true} />)

        expect(getByTestId('buttonAnimated')).toBeDisabled()
    })

    it('executes the onClick handler when clicked', () => {
        const onClickMock = jest.fn()
        const { getByTestId } = render(<WButtonMotion dataTestid="buttonAnimated" text="Go" onClick={onClickMock} />)

        fireEvent.click(getByTestId('buttonAnimated'))

        expect(onClickMock).toHaveBeenCalled()
    })

    it('applies custom styles to the ButtonAnimated', () => {
        const customStyle = { color: 'red', fontSize: '24px' }
        const { getByTestId } = render(<WButtonMotion dataTestid="buttonAnimated" text="Go" style={customStyle} />)

        const buttonElement = getByTestId('buttonAnimated')

        expect(buttonElement).toHaveStyle('color: red')
        expect(buttonElement).toHaveStyle('font-size: 24px')
    })
})
