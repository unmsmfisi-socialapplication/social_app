import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import WPhotoInformation from '.'

describe('WPhotoInformation', () => {
    it('renders the close button', () => {
        render(<WPhotoInformation title="Title" />)
        expect(screen.getByText('X')).toBeInTheDocument()
    })

    it('renders the title', () => {
        render(<WPhotoInformation title="Title" />)
        expect(screen.getByText('Title')).toBeInTheDocument()
    })

    it('renders the subtitle for maximum weight', () => {
        render(<WPhotoInformation title="Title" />)
        expect(screen.getByText('Peso max(10 MB)')).toBeInTheDocument()
    })
})
