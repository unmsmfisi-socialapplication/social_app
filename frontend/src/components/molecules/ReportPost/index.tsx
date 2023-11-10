'use client'
import React from 'react'
import styles from './index.module.scss'
import WComplaint from '../Complaint'
import { Button, Card, Typography } from '@mui/material'

interface WReportPostProps {
    header?: string
    title?: string
}

const WReportPost: React.FC<WReportPostProps> = ({ title, header }) => {
    return (
        <Card className={styles.cardFollow}>
            <div className={styles.header}>
                <Button className={styles.closeButton}>X</Button>
                <Typography variant="h6" component="div" className={styles.headerText}>
                    {header}
                </Typography>
            </div>
            <div className={styles.headerLine}></div>
            <Typography variant="h4" component="div" className={styles.title}>
                {title}
            </Typography>
            <WComplaint title="Odio" />
            <WComplaint title="Abuso y Acoso" />
            <WComplaint title="Privacidad" />
            <WComplaint title="Contenido inapropiado" />
        </Card>
    )
}

export default WReportPost

WReportPost.defaultProps = {
    header: 'Reunir Información',
    title: '¿Qué tipo de problemas quieres reportar?',
}
