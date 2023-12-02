import React from 'react'
import './index.scss'
import { Card, Typography } from '@mui/material'

interface WProfileStatiticsProps {
    posts?: number
    photos?: number
    followers?: number
    following?: number
}

const WProfileStatitics: React.FC<WProfileStatiticsProps> = ({ posts, photos, followers, following }) => {
    return (
        <Card className="user-profile">
            <div className="stat">
                <span className="number">{posts}</span>
                <Typography variant="body2" component="div" className="title">
                    posts
                </Typography>
            </div>
            <div className="stat">
                <span className="number">{photos}</span>
                <Typography variant="body2" component="div" className="title">
                    photos
                </Typography>
            </div>
            <div className="stat">
                <span className="number">{followers}</span>
                <Typography variant="body2" component="div" className="title">
                    followers
                </Typography>
            </div>
            <div className="stat">
                <span className="number">{following}</span>
                <Typography variant="body2" component="div" className="title">
                    following
                </Typography>
            </div>
        </Card>
    )
}

export default WProfileStatitics

WProfileStatitics.defaultProps = {
    posts: 0,
    photos: 0,
    followers: 0,
    following: 0,
}
