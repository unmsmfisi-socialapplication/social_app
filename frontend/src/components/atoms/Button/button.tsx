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
    loading?: boolean
    onClick?: () => void
}

const WButton: React.FC<WButtonProps> = ({
    dataTestid,
    id,
    disabled,
    typeColor,
    text,
    type,
    size,
    variant,
    loading,
    onClick,
}) => {
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
            onClick={onClick}
        >
            {loading ? 'Cargando...' : text}
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
    loading: false,
    onClick: () => {},
}
