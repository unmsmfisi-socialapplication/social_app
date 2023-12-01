import React from 'react';
import Typography from '@mui/material/Typography';
import Box from '@mui/material/Box';
import { Button } from '@mui/material';

interface MainChatActionsProps {
    title?: string;
}

const styles = {
    backgroundColor: '#fff',
    color: 'black',
    alignItems: 'center',
    justifyContent: 'center',
    witdh: '320px',
    height: '166px',
    title: {
        fontSize: 20,
        fontWeight: 700,
    },
    buttonDelete: {
        backgroundColor: '#007AFF',
        color: 'white',
        width: 320,
        height: 47,
        fontSize: 14,
        fontWeight: 700,
        borderRadius: 10,
    },
    buttonReport: {
        backgroundColor: '#C91111',
        color: 'white',
        width: 320,
        height: 47,
        fontSize: 14,
        fontWeight: 600,
        borderRadius: 10,
    },
};


const MainChatActions: React.FC<MainChatActionsProps> = (
    {
        title = defaultValues.title,
    }
) => {
    const mergedProps = { ...defaultValues, title };
    return (
        <Box
            p={2}
            borderBottom={1}
            borderColor="white"
            width={styles.witdh}
            height={styles.height}
            sx={styles}
        >
            <Typography variant="subtitle1" fontWeight={styles.title.fontWeight} fontSize={styles.title.fontSize} sx={{marginBottom:'20px'}}>
                {mergedProps.title}
            </Typography>
            <Box
                display="flex"
                justifyContent="center"
                alignItems="center"
                margin={'10px'}
            >
                {/* Boton azul de eliminar chat */}
                <Button variant="contained" color="primary" style={styles.buttonDelete}>
                    Eliminar Chat
                </Button>
            </Box>
            <Box
                display="flex"
                justifyContent="center"
                alignItems="center"
                margin={'10px'}
            >
                {/* Boton rojo de reportar chat */}
                <Button variant="contained" color="primary" style={styles.buttonReport}>
                    Reportar Chat
                </Button>
            </Box>
        </Box>
    );
};


const defaultValues: MainChatActionsProps = {
    title: 'Configuración de chat',
};
export default MainChatActions;