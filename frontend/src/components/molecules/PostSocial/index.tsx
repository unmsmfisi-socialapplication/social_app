import React from 'react'

import { Avatar, Box, CardMedia } from '@mui/material'
import { WCircleIcon } from '@/components'
import FavoriteIcon from '@mui/icons-material/Favorite'
import ModeCommentIcon from '@mui/icons-material/ModeComment'
import CachedIcon from '@mui/icons-material/Cached'
import ChatBubbleOutlineIcon from '@mui/icons-material/ChatBubbleOutline'
import IosShareIcon from '@mui/icons-material/IosShare'
import Link from 'next/link'
import './index.scss'
import { useHistory } from '@/utilities/Functions'
interface WPostSocialProps {
    published?: string
    postImg?: string
    postText?: string
    user?: {
        name?: string
        userHandler?: string
        alt?: string
        img?: string
    }
    stadistics?: {
        countHeart?: number
        countComment?: number
        countShareds?: boolean
        countCached?: boolean
    }
}

export default function WPostSocial({ published, postImg, user, postText, stadistics }: WPostSocialProps) {
    return (
        <Box className="post_social">
            <section className="post__social--avatar">
                <Avatar alt={user?.alt} src={user?.img} />
            </section>
            <div className="post__social--info">
                <div className="post__social--header">
                    <div className="post__social--title">
                        <h3>{user?.name}</h3>
                        <Link href="#" color="inherit">
                            @{user?.userHandler}
                        </Link>
                        <span>· {published}</span>
                    </div>
                    <div>...</div>
                </div>
                <div className="post__social--body">
                    {postText && <h4 className="post__social--body--text">{postText}</h4>}
                    {postImg && (
                        <CardMedia
                            component="img"
                            image="https://previews.123rf.com/images/grase/grase1410/grase141000110/32759482-lindo-gato-somal%C3%AD-miente-dentro-del-rect%C3%A1ngulo-de-pl%C3%A1stico.jpg"
                            alt="Random"
                            sx={{ width: '98%', height: 380, borderRadius: '10px' }}
                        />
                    )}
                </div>
                <div className="post__social--footer">
                    <Link href="/intranet/post" className="post-icon">
                        <WCircleIcon typeColor="comment" iconSize={18} icon={ChatBubbleOutlineIcon} />
                        <span>{stadistics?.countComment}</span>
                    </Link>
                    <div className="post-icon">
                        <WCircleIcon typeColor="comment" iconSize={18} icon={CachedIcon} />
                    </div>
                    <div className="post-icon">
                        <WCircleIcon typeColor="comment" iconSize={18} icon={FavoriteIcon} />
                        <span>{stadistics?.countHeart}</span>
                    </div>
                    <div className="post-icon">
                        <WCircleIcon typeColor="comment" iconSize={18} icon={IosShareIcon} />
                    </div>
                </div>
            </div>
        </Box>
    )
}

WPostSocial.defaultProps = {
    published: 'hace 1 hora',
    postText:
        'El diseño es un mayor trabajo para la comunidad por ello es importante tomar diferentes variantes , y ocupar trabajos de diferentes tipos como lo menciona san diego en su libro de diseño emocional , El diseño es un mayor trabajo para la comunidad por ello es importante tomar diferentes variantes , y ocupar trabajos de diferentes tipos como lo menciona san diego en su libro de diseño emocional',
    user: {
        name: 'Don Norman',
        userHandler: 'jndler',
    },
    postImg:
        'https://previews.123rf.com/images/grase/grase1410/grase141000110/32759482-lindo-gato-somal%C3%AD-miente-dentro-del-rect%C3%A1ngulo-de-pl%C3%A1stico.jpg',
    stadistics: {
        countHeart: 16,
        countComment: 3,
        countShareds: true,
        countCached: true,
    },
}
