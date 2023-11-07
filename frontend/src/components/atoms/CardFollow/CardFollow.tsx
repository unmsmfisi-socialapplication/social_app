import React from 'react'
import { Card, CardContent, Typography, Button, Avatar } from '@mui/material'
import styles from './CardFollow.module.scss'

interface WCardFollowProps {
    avatar?: string
    name?: string
    userhandle?: string
}

const WCardFollow: React.FC<WCardFollowProps> = ({ avatar, name, userhandle }) => {
    return (
        <Card className={styles.cardFollow}>
            <Avatar src={avatar} className={styles.avatar} />
            <CardContent className={styles.cardContent}>
                <div className={styles.userInfo}>
                    <Typography variant="h6" component="div" className={styles.userName}>
                        {name}
                    </Typography>
                    <Typography color="textSecondary" className={styles.userHandle}>
                        @{userhandle}
                    </Typography>
                </div>
            </CardContent>
            <Button variant="contained" color="primary" className={styles.followButton}>
                Seguir
            </Button>
        </Card>
    )
}

export default WCardFollow

WCardFollow.defaultProps = {
    avatar: 'https://www.pngkey.com/png/full/114-1149878_setting-user-avatar-in-specific-size-without-breaking.png',
    name: 'Nombre de usuario',
    userhandle: 'handle',
}
