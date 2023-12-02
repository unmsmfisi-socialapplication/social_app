import React from 'react'
import Button from '@mui/material/Button'
import './buttonAnimated.scss'

interface WButtonMotionProps {
    id?: string
    dataTestid?: string
    type?: 'submit'
    text?: string
    size?: 'large'
    disabled?: boolean
    onClick?: () => void
    style?: React.CSSProperties
}

const WButtonMotion: React.FC<WButtonMotionProps> = ({
    id,
    type,
    text,
    size,
    dataTestid,
    disabled,
    onClick,
    style,
}) => {
    return (
        <Button
            id={id}
            data-testid={dataTestid}
            style={{ minWidth: size === 'large' ? '100%' : 'auto', ...style }}
            type={type}
            className="button-animated"
            disabled={disabled}
            onClick={onClick}
        >
            {text}
        </Button>
    )
}

WButtonMotion.defaultProps = {
    text: 'Button',
}

export default WButtonMotion
