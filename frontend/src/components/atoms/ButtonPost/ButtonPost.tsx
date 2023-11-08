'use client'
import React from 'react'
import Button from '@mui/material/Button'
import './index.scss'

interface WButtonPost {
    id?: string
    typeColor?: 'primary' | 'secondary' | 'disabled'
    dataTestid?: string
    type?: 'submit'
    text?: string
    size?: 'large'
    disabled?: boolean
    borderRadius?: string
    variant?: 'outlined' | 'contained'
}

const WButtonPost: React.FC<WButtonPost> = ({ dataTestid, id, disabled, typeColor, text, type, size, variant }) => {
    const buttonClass = `button typeButton--${disabled ? 'disabled' : typeColor}`
    return (
        <Button
            id={id}
            data-testid={dataTestid}
            style={{ minWidth: size === 'large' ? '100%' : 'auto' }}
            type={type}
            className={buttonClass}
            size={size}
            disabled={disabled}
            variant={variant}
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
    variant: 'contained',
    dataTestid: 'button',
}
