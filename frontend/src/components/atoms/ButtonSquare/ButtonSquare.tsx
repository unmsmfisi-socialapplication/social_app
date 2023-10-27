'use client'
import React from 'react'
import Button from '@mui/material/Button'
import IconButton from '@mui/material/IconButton'
import SettingsIcon from '@mui/icons-material/Settings'

const SquareButton: React.FC = () => {
    const buttonStyle: React.CSSProperties = {
        width: '50px',
        height: '50px',
        borderRadius: '2px',
        background: 'white',
        border: '1px solid #CCCCCC',
    }

    const iconStyle: React.CSSProperties = {
        color: '#CCCCCC',
    }

    return (
        <Button variant="contained" style={buttonStyle}>
            <IconButton>
                <SettingsIcon style={iconStyle} />
            </IconButton>
        </Button>
    )
}

export default SquareButton
