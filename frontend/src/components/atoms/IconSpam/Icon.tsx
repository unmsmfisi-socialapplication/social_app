'use client'
import React from 'react'
import PersonAddAlt1Icon from '@mui/icons-material/PersonAddAlt1'
import { SvgIcon, SvgIconProps } from '@mui/material'
import './index.scss'

interface WSpamProps {
    typeColor?: 'primary' | 'secondary' | 'comment'
    icon: React.ComponentType<SvgIconProps>
    iconSize?: number
    testId?: string
    text: string
}

const WSpamProps: React.FC<WSpamProps> = ({ typeColor, icon, iconSize, testId, text }) => {
    return (
        <div data-testid={testId} className="content">
            <SvgIcon
                component={icon}
                sx={{ fontSize: iconSize }}
                className={`SpamProps SpamProps--${typeColor}`}
            ></SvgIcon>
            <span className="text">{text}</span>
        </div>
    )
}

export default WSpamProps
WSpamProps.defaultProps = {
    typeColor: 'primary',
    icon: PersonAddAlt1Icon,
    iconSize: 40,
    text: '',
}
