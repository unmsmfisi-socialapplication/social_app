'use client'
import React from 'react'
import AllInclusiveIcon from '@mui/icons-material/AllInclusive'
import PersonIcon from '@mui/icons-material/Person'
import './index.scss'
import { Avatar } from '@mui/material'

interface WCircleImageProps {
    avatarDefaultURL?: string
    typeColor?: 'primary' | 'secondary' | 'third'
    alt?: string
    size?: number
}

const WCircleImage: React.FC<WCircleImageProps> = ({ avatarDefaultURL, alt, typeColor, size }) => {
    return (
        <div className={`circleImage circleImage--${typeColor}`} data-testid="circle-image-container">
            <Avatar alt={alt} src={avatarDefaultURL || ''} sx={{ width: size, height: size }}>
                {!avatarDefaultURL && <PersonIcon style={{ width: '100%', height: '100%' }} />}
            </Avatar>
        </div>
    )
}

export default WCircleImage

WCircleImage.defaultProps = {
    typeColor: 'primary',
    size: 60,
}
