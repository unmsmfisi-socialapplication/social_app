'use client'
import { Avatar, Box, Modal, Typography } from '@mui/material'
import CloseIcon from '@mui/icons-material/Close'
import React, { useState } from 'react'
import './index.scss'
import { WButton } from '@/components'

interface ReportPublicationProps {
    userHandle?: string
    open?: boolean
    onClose?: () => void
    onConfirm?: (reason: string) => void
}

const ReportPublication: React.FC<ReportPublicationProps> = ({ userHandle, open, onClose, onConfirm }) => {
    const [reasonReport, setReasonReport] = useState<string>('')
    const handleSubmitReport = () => {
        // aca se debe reemplazar por una funcion asincrona que debe esperar a que se
        // haga la peticion post de la creacion del reporte y luego si el resultado es exitoso
        // cerrar el modal
        if (onConfirm && onClose) {
            onConfirm(reasonReport)
            onClose()
        }
    }
    const handleChangeReason = (e: React.ChangeEvent<HTMLInputElement>) => {
        setReasonReport(e.target.value)
    }
    return (
        <Modal
            open={open || false}
            onClose={onClose}
            className="modal--main--container"
            aria-labelledby="modal-modal-title"
            aria-describedby="modal-modal-description"
        >
            <Box className="modal--report--container">
                <Box className="modal--report--header">
                    <CloseIcon onClick={onClose} />
                    <Typography variant="h6" fontWeight={800} component="h2">
                        Reunir información
                    </Typography>
                </Box>
                <Box className="modal--report--body">
                    <Typography variant="h5" fontWeight={800} textAlign="center">
                        ¿Qué tipo de problema quieres denunciar?
                    </Typography>
                    <Box className="modal--report--options">
                        <Box display="flex" flexDirection="row" gap={2} alignItems="center">
                            <Box>
                                <Typography variant="h6" fontWeight={800}>
                                    Odio
                                </Typography>
                                <Typography>
                                    Palabras ofensivas, esteriotipos racistas o sexistas, deshumanizacion, incitacion al
                                    miedo o la discriminacion, referencias a discursos de odio y logotipos relacionados
                                    al odio
                                </Typography>
                            </Box>
                            <input type="radio" name="reason" value="hate" onChange={handleChangeReason}></input>
                        </Box>
                        <Box display="flex" flexDirection="row" gap={2} alignItems="center">
                            <Box>
                                <Typography variant="h6" fontWeight={800}>
                                    Abuso y acoso
                                </Typography>
                                <Typography>
                                    Insultos, contenido no deseado de carácter sexual, contenido no apto para el
                                    ambiente laboral y contenido gráfico no deseado, acoso dirigido e incitación al
                                    acoso
                                </Typography>
                            </Box>
                            <input type="radio" name="reason" value="harrasment" onChange={handleChangeReason}></input>
                        </Box>
                        <Box display="flex" flexDirection="row" gap={2} alignItems="center">
                            <Box>
                                <Typography variant="h6" fontWeight={800}>
                                    Privacidad
                                </Typography>
                                <Typography>
                                    Distribución de información privada, distribución no consensuada de imágenes
                                    íntimas, distribución de imágenes de mí que no deseo que estén en la plataforma
                                </Typography>
                            </Box>
                            <input type="radio" name="reason" value="privacity" onChange={handleChangeReason}></input>
                        </Box>
                    </Box>
                    <Box className="modal--report--submit" display="flex" justifyContent="center">
                        <WButton text={`Reportar a @${userHandle}`} onClick={handleSubmitReport} />
                    </Box>
                </Box>
            </Box>
        </Modal>
    )
}

export default ReportPublication

ReportPublication.defaultProps = {
    userHandle: 'XokasXD',
    open: false,
}
