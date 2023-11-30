'use client'
import React from 'react'
import { Box, SvgIcon, SvgIconProps } from '@mui/material'
import AllInclusive from '@mui/icons-material/AllInclusive'
import { useRouter } from 'next/navigation'

import './index.scss'

interface WTagProps {
    icon: React.ComponentType<SvgIconProps>
    text?: string
    isActive?: boolean
    path?: string
    disabled?: boolean
}

const WTag: React.FC<WTagProps> = ({ icon, text, isActive, path = '/' }) => {
    const router = useRouter()

    return (
        <Box onClick={() => router.push(path)} className={isActive ? 'tagLink tagLink--active' : 'tagLink'}>
            <SvgIcon component={icon}></SvgIcon>
            {text}
        </Box>
    )
}

export default WTag

WTag.defaultProps = {
    icon: AllInclusive,
    isActive: false,
    text: 'TagLink',
    path: '/',
    disabled: false,
}
