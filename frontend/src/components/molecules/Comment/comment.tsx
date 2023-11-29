import React from 'react'
import WCircleIcon from './../../atoms/CircleIcon/circleIcon'
import FavoriteIcon from '@mui/icons-material/Favorite'
import ModeCommentIcon from '@mui/icons-material/ModeComment'
import CachedIcon from '@mui/icons-material/Cached'
import BarChartIcon from '@mui/icons-material/BarChart'
import StarIcon from '@mui/icons-material/Star'
import ShareIcon from '@mui/icons-material/Share'
import AccountCircleIcon from '@mui/icons-material/AccountCircle'
import './index.scss'

interface WCommentProps {
    user?: {
        usuario?: string
        name?: string
    }
    sendUser?: string
    typeColor?: 'primary' | 'secondary'
    comment?: string
    fullWidth?: boolean
    countHeart?: number
    countComment?: number
    countShareds?: number
    countStatistics?: number
    countStars?: number
}

const WComment: React.FC<WCommentProps> = ({
    user,
    sendUser,
    comment,
    countHeart,
    countComment,
    countShareds,
    countStatistics,
    countStars,
}) => {
    return (
        <div className="comment-main-container">
            <div className="comment-avatar-main">
                <WCircleIcon typeColor="comment" iconSize={50} icon={AccountCircleIcon} />
                <div>
                    <span>
                        <strong>{user?.name}</strong> @{user?.usuario}
                    </span>
                    <span>
                        Respondiendo a <a href="#">@{sendUser}</a>
                    </span>
                </div>
            </div>
            <div>{comment}</div>
            <div className="comment-icons-main">
                <div className="comment-icon">
                    <WCircleIcon typeColor="comment" iconSize={18} icon={FavoriteIcon} />
                    <span>{countHeart}</span>
                </div>
                <div className="comment-icon">
                    <WCircleIcon typeColor="comment" iconSize={18} icon={ModeCommentIcon} />
                    <span>{countComment}</span>
                </div>
                <div className="comment-icon">
                    <WCircleIcon typeColor="comment" iconSize={18} icon={BarChartIcon} />
                    <span>{countStatistics}</span>
                </div>
                <div className="comment-icon">
                    <WCircleIcon typeColor="comment" iconSize={18} icon={StarIcon} />
                    <span>{countStars}</span>
                </div>
                <div className="comment-icon">
                    <WCircleIcon typeColor="comment" iconSize={18} icon={ShareIcon} />
                </div>
            </div>
        </div>
    )
}

export default WComment

WComment.defaultProps = {
    user: {
        usuario: 'millys123',
        name: 'Milly',
    },
    sendUser: 'janedoel123',
    comment: 'Lorem ipsum dolor sit amet consectetur. Congue et nunc sed nascetur malesuada',
    countHeart: 12,
    countComment: 5,
    countShareds: 6,
    countStatistics: 126,
    countStars: 23,
}
