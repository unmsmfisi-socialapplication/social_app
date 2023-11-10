import '@testing-library/jest-dom'
import React from 'react'
import { render, screen } from '@testing-library/react'
import WPostTypes from './index'
import AddPhotoAlternateIcon from '@mui/icons-material/AddPhotoAlternate'
import SubscriptionsIcon from '@mui/icons-material/Subscriptions'
import GifBoxIcon from '@mui/icons-material/GifBox'
import LocationOnIcon from '@mui/icons-material/LocationOn'

interface WPostTypesProps {
    iconComponent: React.ReactNode
    typeName: string
}

const renderWPostTypes = (props?: Partial<WPostTypesProps>) => {
    const defaultProps: WPostTypesProps = {
        iconComponent: <AddPhotoAlternateIcon />,
        typeName: 'FOTO',
    }

    return render(<WPostTypes {...defaultProps} {...props} />)
}

describe('WPostTypes', () => {
    it('renders with default props', () => {
        renderWPostTypes()
        // Test con los valores predeterminados
        expect(screen.getByTestId('w-post-types')).toBeInTheDocument()
        expect(screen.getByTestId('w-post-types')).toHaveTextContent('Publicar FOTO')
    })

    it('renders with custom props (VIDEO)', () => {
        const customProps: Partial<WPostTypesProps> = {
            iconComponent: <SubscriptionsIcon />,
            typeName: 'VIDEO',
        }

        renderWPostTypes(customProps)
        // Test con valores sobre tipo VIDEO
        expect(screen.getByTestId('w-post-types')).toBeInTheDocument()
        expect(screen.getByTestId('w-post-types')).toHaveTextContent('Publicar VIDEO')
    })

    it('renders with custom props (GIF)', () => {
        const customProps: Partial<WPostTypesProps> = {
            iconComponent: <GifBoxIcon />,
            typeName: 'GIF',
        }

        renderWPostTypes(customProps)
        // Test con valores sobre tipo GIF
        expect(screen.getByTestId('w-post-types')).toBeInTheDocument()
        expect(screen.getByTestId('w-post-types')).toHaveTextContent('Publicar GIF')
    })

    it('renders with custom props (UBICACIÓN)', () => {
        const customProps: Partial<WPostTypesProps> = {
            iconComponent: <LocationOnIcon />,
            typeName: 'UBICACIÓN',
        }

        renderWPostTypes(customProps)
        // Test con valores sobre tipo VIDEO
        expect(screen.getByTestId('w-post-types')).toBeInTheDocument()
        expect(screen.getByTestId('w-post-types')).toHaveTextContent('Publicar UBICACIÓN')
    })

    // Add more test cases for different scenarios, if needed
})
