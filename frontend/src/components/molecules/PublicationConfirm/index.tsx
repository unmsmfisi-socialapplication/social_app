import { Avatar, Box, Modal, Typography } from '@mui/material'
import CloseIcon from '@mui/icons-material/Close'
import React from 'react'
import './index.scss'
import WButton from './../../atoms/Button/button'

interface PublicationConfirmProps {
    userName?: string
    userHandle?: string
    avatarURL?: string
    publicationContent?: string
    open?: boolean
    dataTestid?: string
    onClose?: () => void
    onConfirm?: () => void
}

const PublicationConfirm: React.FC<PublicationConfirmProps> = ({
    userName,
    userHandle,
    avatarURL,
    publicationContent,
    open,
    dataTestid,
    onClose,
    onConfirm,
}) => {
    const handleConfirmPublication = () => {
        // aca se debe reemplazar por una funcion asincrona que debe esperar a que se
        // haga la peticion post de la creacion de la publicacion y luego si el resultado es exitoso
        // cerrar el modal
        if (onConfirm && onClose) {
            onConfirm()
            onClose()
        }
    }
    return (
        <>
            <Modal
                open={open || false}
                onClose={onClose}
                className="modal--main--container"
                aria-labelledby="modal-modal-title"
                aria-describedby="modal-modal-description"
                data-testid={dataTestid}
            >
                <Box className="modal--publication--confirm">
                    <Box className="modal--publication--header">
                        <CloseIcon onClick={onClose} />
                        <Typography variant="h6" fontWeight={800} component="h2">
                            Confirmar publicación
                        </Typography>
                    </Box>
                    <Box className="modal--publication--body">
                        <Box className="modal--publication--user--info">
                            <Avatar className="chat_user_avatar" src={avatarURL} />
                            <Box>
                                <Typography style={{ fontWeight: '700' }}>{userName}</Typography>
                                <Typography style={{ color: 'rgba(0, 0, 0, 0.6)' }}>@{userHandle}</Typography>
                            </Box>
                        </Box>
                        <Typography>{publicationContent}</Typography>
                        <Box className="modal--publication-buttons">
                            <WButton text="Cancelar" onClick={onClose} />
                            <WButton text="Confirmar" onClick={handleConfirmPublication} />
                        </Box>
                    </Box>
                </Box>
            </Modal>
        </>
    )
}

export default PublicationConfirm

PublicationConfirm.defaultProps = {
    avatarURL:
        'https://images.ecestaticos.com/FjaDMYL1rpd8bqAVvR91YL-gZbY=/0x0:2252x1336/1200x1200/filters:fill(white):format(jpg)/f.elconfidencial.com%2Foriginal%2Fae2%2F47e%2F66d%2Fae247e66d9b8d8928d41a592b61690ca.jpg',
    userName: 'XokasXD',
    userHandle: 'XokasXD',
    publicationContent:
        'Lorem, ipsum dolor sit amet consectetur adipisicing elit. Fugiat repellat quaerat totam,aliquam omnis dolorem doloribus molestiae veritatis, quasi fugit odio eveniet nemo illumeligendi aut.',
    open: false,
}
