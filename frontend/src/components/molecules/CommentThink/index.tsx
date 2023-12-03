'use client'
import React from 'react'
import './index.scss'
import WButton from './../../atoms/Button/button'
import Textarea from '@mui/joy/Textarea'
import WCircleImage from './../../atoms/CircleImage/circleImage'

interface CommentThinkProps {
    avatarDefaultURL?: string
    publicTag?: string
    placeholder?: string
}

const WCommentThink: React.FC<CommentThinkProps> = ({ avatarDefaultURL, publicTag, placeholder }) => {
    return (
        <div className="comment_think_main_container">
            <div className="comment_think_container">
                <WCircleImage avatarDefaultURL={avatarDefaultURL} size={80} typeColor="third" />
                <WButton typeColor="white" text={publicTag} variant="outlined" />
            </div>
            <Textarea
                style={{ height: '200px' }}
                color="neutral"
                minRows={2}
                placeholder={placeholder}
                size="lg"
                variant="plain"
            />
        </div>
    )
}

export default WCommentThink

WCommentThink.defaultProps = {
    avatarDefaultURL: '',
    publicTag: 'Público',
    placeholder: '',
}
