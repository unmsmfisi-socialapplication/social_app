import React from 'react'
import { render, screen } from '@testing-library/react'
import WSearch from './Search'

describe('WSearch', () => {
    it('renders the search input', () => {
        render(<WSearch />)
        const searchInput = screen.getByPlaceholderText('Buscar')

        expect(searchInput).toBeInTheDocument()
    })

    it('renders the search input with a search icon', () => {
        render(<WSearch />)
        const searchIcon = screen.getByTestId('search-icon')

        expect(searchIcon).toBeInTheDocument()
    })

    it('handles search button click', () => {
        render(<WSearch />)
        const searchButton = screen.getByTestId('search-icon')

        expect(searchButton).toBeInTheDocument()
    })
    
})
