import React, { useState } from 'react'
import { Avatar } from '@mui/material'
import './index.scss'

interface WUserCHATCTAProps {
    avatarURL: string
    userName: string
    userHandle: string
}

const WUserChatCTA = ({ avatarURL = '', userName = 'User', userHandle = 'user' }: WUserCHATCTAProps) => {
    const [showCTAOptions, setShowCTAOptions] = useState<boolean>(false)
    return (
        <div className="chat_user_cta_container">
            <div className="chat_user_cta_user_info">
                <Avatar className="chat_user_avatar" src={avatarURL} />
                <p>
                    <strong>{userName}</strong>
                    <span>@{userHandle}</span>
                </p>
            </div>
            <div className="chat_user_cta_user_options" onClick={() => setShowCTAOptions(!showCTAOptions)}>
                <span className="cta_dot"></span>
                <span className="cta_dot"></span>
                <span className="cta_dot"></span>
                {showCTAOptions && (
                    <ul>
                        <li>Cerrar Sesión</li>
                        <li>Ver Perfil</li>
                    </ul>
                )}
            </div>
        </div>
    )
}

export default WUserChatCTA