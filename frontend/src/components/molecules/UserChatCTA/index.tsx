'use client'
import React, { useState } from 'react'
import { Avatar } from '@mui/material'
import './index.scss'

interface WUserCHATCTAProps {
    avatarURL?: string
    userName?: string
    userHandle?: string
}

const WUserChatCTA: React.FC<WUserCHATCTAProps> = ({ avatarURL, userName, userHandle }: WUserCHATCTAProps) => {
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

WUserChatCTA.defaultProps = {
    avatarURL:
        'https://images.ecestaticos.com/FjaDMYL1rpd8bqAVvR91YL-gZbY=/0x0:2252x1336/1200x1200/filters:fill(white):format(jpg)/f.elconfidencial.com%2Foriginal%2Fae2%2F47e%2F66d%2Fae247e66d9b8d8928d41a592b61690ca.jpg',
    userName: 'XokasXD',
    userHandle: 'XokasXD',
}
