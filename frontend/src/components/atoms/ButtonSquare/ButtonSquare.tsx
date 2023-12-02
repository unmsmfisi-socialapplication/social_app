import React from 'react'
import Button from '@mui/material/Button'
import SettingsIcon from '@mui/icons-material/Settings'

interface SquareButtonProps {
    onClick?: () => void
    children?: React.ReactNode
}

const SquareButton: React.FC<SquareButtonProps> = ({ onClick, children }) => {
    const buttonStyle: React.CSSProperties = {
        width: '50px',
        height: '50px',
        borderRadius: '2px',
        background: 'white',
        border: '1px solid #CCCCCC',
        display: 'flex',
        alignItems: 'center',
        justifyContent: 'center',
    }

    const iconStyle: React.CSSProperties = {
        color: '#CCCCCC',
    }

    return (
        <Button variant="contained" style={buttonStyle} data-testid="outer-button" onClick={onClick}>
            <SettingsIcon style={iconStyle} />
            {children}
        </Button>
    )
}

export default SquareButton
