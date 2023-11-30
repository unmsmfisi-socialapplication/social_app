import { render, screen } from '@testing-library/react'
import WCircleImage from './circleImage'

describe('WCircleImage Component', () => {
    test('renders the component with default props', () => {
        render(<WCircleImage />)
        const circleImageContainer = screen.getByTestId('circle-image-container')

        expect(circleImageContainer).toBeInTheDocument()

        const avatar = circleImageContainer.querySelector('.MuiAvatar-root')

        expect(avatar).toBeInTheDocument()
    })

    test('renders the component with custom props', () => {
        render(
            <WCircleImage
                typeColor="secondary"
                size={80}
                avatarDefaultURL="https://example.com/avatar.jpg"
                alt="User"
            />,
        )
        const circleImageContainer = screen.getByTestId('circle-image-container')

        expect(circleImageContainer).toBeInTheDocument()

        const avatar = circleImageContainer.querySelector('.MuiAvatar-root')

        expect(avatar).toBeInTheDocument()
    })
})
