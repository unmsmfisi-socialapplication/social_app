import React from 'react'
import { Card, CardContent, Typography, Button } from '@mui/material'
import styles from './TopicFollow.module.scss'

interface WTopicFollowProps {
    name?: string
    topicHandle?: string
}

const WTopicFollow: React.FC<WTopicFollowProps> = ({ name, topicHandle }) => {
    return (
        <Card className={styles.cardFollow}>
            <CardContent className={styles.cardContent}>
                <div className={styles.topicInfo}>
                    <Typography variant="h6" component="div" className={styles.userName}>
                        {name}
                    </Typography>
                    <Typography color="textSecondary" className={styles.topicHandle}>
                        {topicHandle}
                    </Typography>
                </div>
            </CardContent>
            <Button variant="contained" color="primary" className={styles.followButton}>
                Seguir
            </Button>
            <Button className={styles.removeButton}>X</Button>
        </Card>
    )
}

export default WTopicFollow

WTopicFollow.defaultProps = {
    name: 'Nombre del Tópico',
    topicHandle: 'handle',
}
