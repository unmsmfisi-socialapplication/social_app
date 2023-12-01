import React from 'react'
import Typography from '@mui/material/Typography'
import Box from '@mui/material/Box'
import Switch from '@mui/material/Switch'

interface FooterChatActionsProps {
    title?: string
    subtitle?: string
}

const styles = {
    backgroundColor: '#fff',
    color: 'black',
    alignItems: 'center',
    justifyContent: 'center',
    witdh: '320px',
    height: '120px',
    title: {
        fontSize: 20,
        fontWeight: 700,
    },
    subtitle: {
        fontSize: 14,
        fontWeight: 400,
    },
}

const FooterChatActions: React.FC<FooterChatActionsProps> = ({ title = defaultValues.title }) => {
    const mergedProps = { ...defaultValues, title }
    return (
        <Box p={2} borderBottom={1} borderColor="white" width={styles.witdh} height={styles.height} sx={styles}>
            <Typography
                variant="subtitle1"
                fontWeight={styles.title.fontWeight}
                fontSize={styles.title.fontSize}
                sx={{ marginBottom: '30px' }}
            >
                {mergedProps.title}
            </Typography>
            <Box display="flex" justifyContent="center" alignItems="center" margin={'10px'}>
                <Typography
                    variant="subtitle1"
                    fontWeight={styles.subtitle.fontWeight}
                    fontSize={styles.subtitle.fontSize}
                >
                    {mergedProps.subtitle}
                </Typography>
                <Switch />
            </Box>
        </Box>
    )
}

const defaultValues: FooterChatActionsProps = {
    title: 'Configuración de mensajes',
    subtitle: 'Silenciar Notificaciones de Mensajes',
}
export default FooterChatActions
