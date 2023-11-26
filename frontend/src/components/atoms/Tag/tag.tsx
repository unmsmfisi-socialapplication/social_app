'use client'
import React from 'react'
import { SvgIcon, SvgIconProps } from '@mui/material'
import AllInclusive from '@mui/icons-material/AllInclusive'
import './index.scss'
import Link from 'next/link'

interface WTagProps {
    icon: React.ComponentType<SvgIconProps>
    text?: string
    isActive?: boolean
    path?: string
    dataTestId?: string
}


//aca se crea el link
const WTag: React.FC<WTagProps> = ({ icon, text, isActive, path, dataTestId }) => {
    return (
        <Link data-testid={dataTestId} className={isActive ? 'tagLink tagLink--active' : 'tagLink'} href={path || ''}>
            <SvgIcon component={icon}></SvgIcon>
            {text}
        </Link>
    )
} 

export default WTag

WTag.defaultProps = {
    icon: AllInclusive,
    isActive: false,
    text: 'TagLink',
    path: '/',
}