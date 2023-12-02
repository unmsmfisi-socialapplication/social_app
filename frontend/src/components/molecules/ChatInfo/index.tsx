import React from 'react';
import Avatar from '@mui/material/Avatar';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';

interface ChatInfoProps {
    username?: string;
    avatarSrc?: string;
    message?: string;
    time?: string;
}

const styles = {
    backgroundColor: '#113E70',
    color: 'white',
    alignItems: 'center',
    witdh: '365px',
    height: '65px',
    avatar: {
        width: 65,
        height: 65,
    },
    username: {
        fontWeight: 'bold',
        fontSize: 14,
    },
    message: {
        fontSize: 11,
    },
    time: {
        fontSize: 11,
    },
};

const ChatInfo: React.FC<ChatInfoProps> = (
    { 
        username = defaultValues.username,
        avatarSrc= defaultValues.avatarSrc,
        message= defaultValues.message,
        time= defaultValues.time 
    }
) => {
    const MAX_MESSAGE_LENGTH = 30;
    const mergedProps = { ...defaultValues, username, avatarSrc, message, time };
    const displayMessage =
        message && message.length > MAX_MESSAGE_LENGTH ? `${message.slice(0, MAX_MESSAGE_LENGTH)}...` : message;
    return (
        <Box
            display="flex"
            p={2}
            borderBottom={1}
            borderColor="divider"
            width={styles.witdh}
            height={styles.height}
            sx={styles}
        >
            <Avatar src={mergedProps.avatarSrc} alt="User Avatar" sx={{ width: styles.avatar.width, height: styles.avatar.height }} />
            <Box ml={2} flexGrow={1}>
                <Typography variant="subtitle1" fontWeight={styles.username.fontWeight} fontSize={styles.username.fontSize}>
                    {mergedProps.username}
                </Typography>
                <Typography variant="body1" fontSize={styles.message.fontSize}>
                    {displayMessage}
                </Typography>
            </Box>
            <Typography variant="body2" fontSize={styles.time.fontSize}>
                {mergedProps.time}
            </Typography>
        </Box>
    );
};

const defaultValues: ChatInfoProps = {
    username: 'Nombre de Usuario',
    avatarSrc: 'https://png.pngtree.com/png-vector/20190710/ourlarge/pngtree-user-vector-avatar-png-image_1541962.jpg',
    message: 'Hola, este es un mensaje largo de prueba',
    time: '10:30 AM',
};
export default ChatInfo;




