// WButtonPost.tsx
import React from 'react'
import Button from '@mui/material/Button'
import './index.scss'

interface WButtonPostProps {
    typeColor?: 'primary' | 'secondary' | 'disabled'
    text?: string
    size?: 'large'
    disabled?: boolean
    borderRadius?: string
    onClick?: () => void // Nueva prop onClick
}

const WButtonPost: React.FC<WButtonPostProps> = ({ disabled, typeColor, text, size, onClick }) => {
    const buttonClass = `button typeButton--${disabled ? 'disabled' : typeColor}`
    return (
        <Button
            style={{ minWidth: size === 'large' ? '100%' : 'auto', borderRadius: '18px' }}
            className={buttonClass}
            size={size}
            disabled={disabled}
            onClick={onClick} // Usa la prop onClick aquí
        >
            {text}
        </Button>
    )
}

export default WButtonPost

WButtonPost.defaultProps = {
    typeColor: 'primary',
    text: 'Button',
    disabled: false,
}
