import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import CommentThink from './../../../components/molecules/CommentThink'

describe('CommentThink', () => {
    it('renders the CommentThink molecule component', () => {
        const { getByTestId } = render(<CommentThink publicTag="Público" placeholder="Escribe..." dataTestid="commentId" />)
        
        expect(getByTestId('commentId')).toBeInTheDocument()
        expect(screen.getByTestId('commentId')).toHaveTextContent("Público")
    })
})
