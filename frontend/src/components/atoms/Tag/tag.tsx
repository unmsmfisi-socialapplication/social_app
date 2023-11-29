'use client'
import React from 'react'
import { Box, SvgIcon, SvgIconProps } from '@mui/material'
import AllInclusive from '@mui/icons-material/AllInclusive'
import './index.scss'
import Link from 'next/link'

interface WTagProps {
    icon: React.ComponentType<SvgIconProps>
    text?: string
    isActive?: boolean
    path?: string
}

const WTag: React.FC<WTagProps> = ({ icon, text, isActive, path }) => {
    return (
        <Link href={`intranet/${path}`} style={{ textDecoration: 'none', color: 'black' }}>
            <Box className={isActive ? 'tagLink tagLink--active' : 'tagLink'}>
                <SvgIcon component={icon}></SvgIcon>
                {text}
            </Box>
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
