import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import WPostLocation from '.'

describe('WPostLocation', () => {
    it('should render a container with all the required elements', () => {
        render(<WPostLocation />)
        expect(screen.getByText('publicar UBICACIÓN')).toBeInTheDocument()
        expect(screen.getByText('X')).toBeInTheDocument()
        expect(screen.getByTitle('Google Maps')).toBeInTheDocument()
        expect(screen.getByText('Cancelar')).toBeInTheDocument()
        expect(screen.getByText('Publicar')).toBeInTheDocument()
    })

    it('should display the header text "publicar UBICACIÓN"', () => {
        render(<WPostLocation />)
        expect(screen.getByText('publicar UBICACIÓN')).toBeInTheDocument()
    })

    it('should render the Google Maps iframe with an invalid or broken URL', () => {
        render(<WPostLocation />);
        const iframe = screen.getByTitle('Google Maps');
        expect(iframe).toBeInTheDocument();
        expect(iframe.src).toBe('https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d387273.5004385689!2d-74.25908989999999!3d40.6976701!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x89c24fa5d33f083b%3A0xc80b8f06e177fe62!2sNew%20York%2C%20NY%2C%20USA!5e0!3m2!1sen!2sin!4v1638529735378!5m2!1sen!2sin');
      });
})
