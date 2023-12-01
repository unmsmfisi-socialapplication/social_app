'use client'
import React from 'react'
import styles from './index.module.scss'
import { Box, Button, Card, Typography } from '@mui/material'
import WButton from './../../atoms/Button/button'

interface PhotoInformationProps {
    title?: string
    open?: boolean
}

const WPhotoInformation: React.FC<PhotoInformationProps> = ({ title, open }) => {
    return (
        <Card className={styles.cardFollow}>
            <div className={styles.header}>
                <Button className={styles.closeButton}>X</Button>
                <Typography variant="h6" component="div" className={styles.headerText}>
                    {title}
                </Typography>
            </div>
            <div className={styles.headerLine}></div>
            <Typography variant="h6" component="div" className={styles.subtitle}>
                Peso max(10 MB)
            </Typography>
            <Typography variant="h6" component="div" className={styles.subtitle}>
                Tamaño max(851 x 315 px)
            </Typography>
            <Typography variant="h6" component="div" className={styles.subtitle}>
                La plataforma acepta estas especificaciones para la foto de portada
            </Typography>
            <Box className="modal--report--submit" display="flex" justifyContent="center">
                <WButton text={`Cargar foto`} />
            </Box>
        </Card>
    )
}

export default WPhotoInformation

WPhotoInformation.defaultProps = {
    title: 'XokasXD',
    open: false,
}
