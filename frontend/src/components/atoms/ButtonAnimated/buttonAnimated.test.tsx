import  WButtonMotion  from './buttonAnimated'
import { render } from '@testing-library/react'

describe('ButtonAnimated', () => {
    it('renders the ButtonAnimated', () => {
        const { getByTestId } = render(<WButtonMotion dataTestid="buttonAnimated" text="Go" />)

        expect(getByTestId('buttonAnimated')).toHaveTextContent('Go')
    })
    it('renders the disabled ButtonAnimated', () => {
        const { getByTestId } = render(<WButtonMotion dataTestid="buttonAnimated" text="Go" disabled={true} />)

        expect(getByTestId('buttonAnimated')).toBeDisabled()
    })
})
