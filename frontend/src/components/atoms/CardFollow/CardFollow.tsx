// WCardFollow.tsx
import React from 'react';
import { Card, CardContent, Typography, Button, Avatar } from '@mui/material';
import styles from './CardFollow.module.scss';

interface WCardFollowProps {
  userAvatar: string;
  userName: string;
  userHandle: string;
}

const WCardFollow: React.FC<WCardFollowProps> = ({ userAvatar, userName, userHandle }) => {
  return (
    <Card className={styles.cardFollow}>
    <Avatar src={userAvatar} className={styles.avatar} />
      <CardContent className={styles.cardContent}>   
        <div className={styles.userInfo}>
          <Typography variant="h6" component="div" className={styles.userName}>
            {userName}
          </Typography>
          <Typography color="textSecondary" className={styles.userHandle}>
            @{userHandle}
          </Typography>
        </div> 
      </CardContent>
        <Button variant="contained" color="primary" className={styles.followButton}>
          Seguir
        </Button>
    </Card>
  );
};

export default WCardFollow;
