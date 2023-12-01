import React from 'react';
import Avatar from '@mui/material/Avatar';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { Button } from '@mui/material';

interface HeaderChatActionsProps {
    username?: string;
    avatarSrc?: string;
}

const styles = {
    backgroundColor: '#fff',
    color: 'black',
    alignItems: 'center',
    justifyContent: 'center',
    witdh: '250px',
    height: '354px',
    avatar: {
        width: 250,
        height: 250,
    },
    username: {
        fontWeight: 'bold',
        fontSize: 29,
    },
    profile: {
        fontSize: 21,
        fontWeight: 400,
    },
    buttonProfile: {
        backgroundColor: '#FFFFFF',
        color: 'black',
        width: 320,
        height: 47,
        fontSize: 21,
        fontWeight: 400,
        borderRadius: 10,
    },
};


const HeaderChatActions: React.FC<HeaderChatActionsProps> = (
    {
        username = defaultValues.username,
        avatarSrc = defaultValues.avatarSrc,
    }
) => {
    const mergedProps = { ...defaultValues, username, avatarSrc };
    return (
        <Box
            p={2}
            borderBottom={1}
            borderColor="white"
            width={styles.witdh}
            height={styles.height}
            sx={styles}
        >
            <Avatar src={mergedProps.avatarSrc} alt="User Avatar" sx={{ width: styles.avatar.width, height: styles.avatar.height }} />
            <Typography variant="subtitle1" fontWeight={styles.username.fontWeight} fontSize={styles.username.fontSize}>
                {mergedProps.username}
            </Typography>
            {/* Boton de "ver perfil con bg transparente" */}
            <Box
                display="flex"
                justifyContent="center"
                alignItems="center"
            >
                {/* <Typography variant="subtitle1" fontWeight={styles.profile.fontWeight} fontSize={styles.profile.fontSize}>
                    Ver perfil
                </Typography> */}
                <Button variant="text" style={styles.buttonProfile} sx={{textTransform: 'none'}}>
                    Ver perfil
                </Button>
            </Box>
        </Box>
    );
};


const defaultValues: HeaderChatActionsProps = {
    username: 'Samuel De Luque',
};
export default HeaderChatActions;