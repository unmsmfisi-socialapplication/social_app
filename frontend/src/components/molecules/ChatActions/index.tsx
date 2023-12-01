import React from 'react'
import HeaderChatActions from './components/header'
import { Box } from '@mui/material'
import MainChatActions from './components/main'
import FooterChatActions from './components/footer'

interface ChatActionsProps {
    username?: string
    avatarSrc?: string
}

const styles = {
    backgroundColor: '#fff',
    color: 'black',
    alignItems: 'center',
    witdh: '389px',
    height: '934px',
}

const ChatActions: React.FC<ChatActionsProps> = ({
    username = defaultValues.username,
    avatarSrc = defaultValues.avatarSrc,
}) => {
    const mergedProps = { ...defaultValues, username, avatarSrc }
    return (
        <Box p={2} borderBottom={1} borderColor="divider" width={styles.witdh} height={styles.height} sx={styles}>
            <Box display={'flex'} justifyContent={'center'}>
                <HeaderChatActions username={mergedProps.username} avatarSrc={mergedProps.avatarSrc} />
            </Box>
            <Box display={'flex'} justifyContent={'center'}>
                <MainChatActions />
            </Box>
            <Box display={'flex'} justifyContent={'center'}>
                <FooterChatActions />
            </Box>
        </Box>
    )
}

const defaultValues: ChatActionsProps = {
    username: 'Samuel De Luque',
    avatarSrc: 'https://i.pinimg.com/originals/2a/3a/6b/2a3a6b3e4a1b5b3d7e3b8e4c9a4c0f8d.png',
}

export default ChatActions
