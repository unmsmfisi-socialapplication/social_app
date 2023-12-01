import React from 'react'
import { Box, Avatar, Typography } from '@mui/material'
import styles from './index.module.scss'

// Definición de las propiedades que espera el componente WChatPreview
interface WChatPreviewProps {
    avatar: string
    name: string
    messagePreview: string
    time: string
}

// Definición del componente funcional WChatPreview
const WChatPreview: React.FC<WChatPreviewProps> = ({ avatar, name, messagePreview, time }) => {
    return (
        // Contenedor principal del componente con la clase CSS chatPreview
        <Box className={styles.chatPreview}>
            {/* Contenedor de información con la clase CSS info */}
            <Box className={styles.info}>
                {/* Componente Avatar de Material-UI con la clase CSS avatar */}
                <Avatar src={avatar} className={styles.avatar} sizes="450px 450px" />
                {/* Contenedor de datos con la clase CSS data */}
                <Box className={styles.data}>
                    {/* Componente Typography de Material-UI con la clase CSS name */}
                    <Typography className={styles.name}>{name}</Typography>
                    {/* Componente Typography de Material-UI con la clase CSS message */}
                    <Typography className={styles.message}>{messagePreview}</Typography>
                </Box>
            </Box>
            {/* Componente Typography de Material-UI con la clase CSS time */}
            <Typography className={styles.time}>{time}</Typography>
        </Box>
    )
}

// Propiedades por defecto del componente en caso de no recibir valores
WChatPreview.defaultProps = {
    avatar: 'https://scontent.flim15-1.fna.fbcdn.net/v/t39.30808-6/327176494_836675897571806_7271269400185669869_n.png?_nc_cat=110&ccb=1-7&_nc_sid=efb6e6&_nc_ohc=yLWwkUErvAAAX9If-sL&_nc_ht=scontent.flim15-1.fna&oh=00_AfDjIBf_tedT-iuEaRQzUMgUpu_CoxBl_f9Wx2XTWinDiw&oe=656D608F',
    name: 'Samuel de Luque Batuecas',
    messagePreview: 'Hey muy buenas a todos guapísimos',
    time: '20:17',
}

// Exportación del componente WChatPreview como componente predeterminado del módulo
export default WChatPreview
