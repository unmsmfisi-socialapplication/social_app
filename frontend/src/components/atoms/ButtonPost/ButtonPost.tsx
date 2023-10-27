'use client'
import React from 'react'
import Button from '@mui/material/Button'
import './index.scss'

interface WButtonPost {
    typeColor?: 'primary' | 'secondary' | 'disabled'
    text?: string
    size?: 'large'
    disabled?: boolean
    borderRadius?: string
}

const WButtonPost: React.FC<WButtonPost> = ({ disabled, typeColor, text, size }) => {
    const buttonClass = `button typeButton--${disabled ? 'disabled' : typeColor}`
    return (
        <Button
            style={{ minWidth: size === 'large' ? '100%' : 'auto', borderRadius: '18px' }}
            className={buttonClass}
            size={size}
            disabled={disabled}
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
