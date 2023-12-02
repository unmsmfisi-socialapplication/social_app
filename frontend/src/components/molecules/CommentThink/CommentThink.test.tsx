import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import CommentThink from '.'

describe('CommentThink', () => {
    it('should render a WCircleImage component with props avatarDefaultURL when CommentThink is rendered', () => {
        const avatarDefaultURL = 'https://example.com/avatar.png'
        const size = 80
        const typeColor = 'third'
        render(<CommentThink avatarDefaultURL={avatarDefaultURL} />)
        const circleImage = screen.getByTestId('circle-image-container')
        expect(circleImage).toBeInTheDocument()
        expect(circleImage).toHaveClass(`circleImage--${typeColor}`)
    })

    it('should render a WButton component with default props when publicTag is not provided when CommentThink is rendered', () => {
        render(<CommentThink />)
        const button = screen.getByRole('button')
        expect(button).toBeInTheDocument()
        expect(button).toHaveClass('button')
        expect(button).toHaveClass('typeButton--white')
        expect(button).toHaveTextContent('Público')
    })

    it("should render a Textarea component with style height set to 'opx' when CommentThink is rendered", () => {
        const placeholder = 'Enter your comment...'
        render(<CommentThink placeholder={placeholder} />)
        const textarea = screen.getByRole('textbox')
        expect(textarea).toBeInTheDocument()
        expect(textarea).toHaveStyle({ height: '0px' })
    })
})
