import React from 'react'
import { SvgIcon, SvgIconProps } from '@mui/material'
import AllInclusiveIcon from '@mui/icons-material/AllInclusive'
import './index.scss'

interface WCircleIconProps {
    typeColor?: 'primary' | 'secondary' | 'comment'
    icon?: React.ComponentType<SvgIconProps>
    iconSize?: number
    dataTestid?: string
    style?: React.CSSProperties
}

const WCircleIcon: React.FC<WCircleIconProps> = ({
    typeColor = 'primary',
    icon: IconComponent = AllInclusiveIcon,
    iconSize = 40,
    dataTestid,
    style,
}) => {
    return (
        <SvgIcon
            data-testid={dataTestid}
            component={IconComponent}
            sx={{ fontSize: iconSize, ...style }}
            className={`circleIcon circleIcon--${typeColor}`}
        />
    )
}

export default WCircleIcon
