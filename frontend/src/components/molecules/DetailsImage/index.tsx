import { Avatar, Box, Typography } from '@mui/material'
import React from 'react'
import './index.scss'

interface WDetailsImageProps {
    userName?: string
    userHandle?: string
    avatarURL?: string
    avatarSize?: number
    textSize?: number
}

const WDetailsImage: React.FC<WDetailsImageProps> = ({ userName, userHandle, avatarURL, avatarSize, textSize }) => {
    return (
        <Box className="details--user--info">
            <Avatar
                className="details--user--avatar"
                style={{ width: `${avatarSize}px`, height: `${avatarSize}px` }}
                src={avatarURL}
                alt="User Avatar"
            />
            <Box>
                <Typography style={{ fontWeight: '700', fontSize: `${textSize}px` }}>{userName}</Typography>
                <Typography style={{ color: 'rgba(0, 0, 0, 0.6)', fontSize: textSize ? `${textSize - 1}px` : 15 }}>
                    @{userHandle}
                </Typography>
            </Box>
        </Box>
    )
}

export default WDetailsImage

WDetailsImage.defaultProps = {
    avatarURL:
        'https://images.ecestaticos.com/FjaDMYL1rpd8bqAVvR91YL-gZbY=/0x0:2252x1336/1200x1200/filters:fill(white):format(jpg)/f.elconfidencial.com%2Foriginal%2Fae2%2F47e%2F66d%2Fae247e66d9b8d8928d41a592b61690ca.jpg',
    userName: 'XokasXD',
    userHandle: 'XokasXD',
    avatarSize: 60,
    textSize: 16,
}
