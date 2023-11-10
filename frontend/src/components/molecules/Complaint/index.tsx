import React, { useState } from 'react'
import styles from './complaint.module.scss'
import { Card, CardContent, Typography } from '@mui/material'
import WRadioButton from '@/components/atoms/RadioButton/RadioButton'

interface WComplaintProps {
    title?: string
    description?: string
}

const WComplaint: React.FC<WComplaintProps> = ({ title, description }) => {
    const [selectedValue, setSelectedValue] = useState('')

    const handleChange = (event: React.ChangeEvent<HTMLInputElement>) => {
        setSelectedValue(event.target.value)
    }

    return (
        <Card className={styles.cardFollow}>
            <CardContent className={styles.cardContent}>
                <div className={styles.topicInfo}>
                    <Typography variant="h6" component="div" className={styles.title}>
                        {title}
                    </Typography>
                    <Typography color="textSecondary" className={styles.description}>
                        {description}
                    </Typography>
                </div>
            </CardContent>
            <WRadioButton selectedValue={selectedValue} handleChange={handleChange} value="a" />
        </Card>
    )
}

export default WComplaint

WComplaint.defaultProps = {
    title: 'Odio',
    description:
        'Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.',
}
