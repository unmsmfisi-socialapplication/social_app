import React from 'react'
import { render, screen, fireEvent } from '@testing-library/react'
import WTopicFollow from './TopicFollow'

describe('WTopicFollow Component', () => {
    it('renders correctly with default props', () => {
        render(<WTopicFollow />)

        expect(screen.getByText('Nombre del Tópico')).toBeInTheDocument()
        expect(screen.getByText('handle')).toBeInTheDocument()
        expect(screen.getByText('Seguir')).toBeInTheDocument()
        expect(screen.getByText('X')).toBeInTheDocument()
    })

    it('renders with provided props', () => {
        render(<WTopicFollow name="Custom Topic" topicHandle="customHandle" />)

        expect(screen.getByText('Custom Topic')).toBeInTheDocument()
        expect(screen.getByText('customHandle')).toBeInTheDocument()
        expect(screen.getByText('Seguir')).toBeInTheDocument()
        expect(screen.getByText('X')).toBeInTheDocument()
    })

    it('triggers follow button click', () => {
        render(<WTopicFollow />)

        const followButton = screen.getByText('Seguir')
        expect(followButton).toBeInTheDocument()

        fireEvent.click(followButton)
    })

    it('triggers remove button click', () => {
        render(<WTopicFollow />)

        const removeButton = screen.getByText('X')
        expect(removeButton).toBeInTheDocument()

        fireEvent.click(removeButton)
    })

    it('renders with different props after updating', () => {
        const { rerender } = render(<WTopicFollow name="Topic" topicHandle="handle" />)

        expect(screen.getByText('Topic')).toBeInTheDocument()
        expect(screen.getByText('handle')).toBeInTheDocument()

        rerender(<WTopicFollow name="Updated Topic" topicHandle="updatedHandle" />)

        expect(screen.getByText('Updated Topic')).toBeInTheDocument()
        expect(screen.getByText('updatedHandle')).toBeInTheDocument()
    })
})
