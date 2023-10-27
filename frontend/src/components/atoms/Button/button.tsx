'use client'
import React from 'react'
import Button from '@mui/material/Button'
import './index.scss'

interface WButtonProps {
    id?: string
    dataTestid?: string
    typeColor?: 'primary' | 'secondary' | 'terciary' | 'quaternary' | 'disabled' | 'white'
    type?: 'submit'
    text?: string
    size?: 'large'
    disabled?: boolean
    variant?: 'outlined' | 'contained'
}

const WButton: React.FC<WButtonProps> = ({ dataTestid, id, disabled, typeColor, text, type, size, variant }) => {
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

export default WButton

WButton.defaultProps = {
    typeColor: 'primary',
    text: 'Button',
    disabled: false,
    variant: 'contained',
    dataTestid: 'button',
}
