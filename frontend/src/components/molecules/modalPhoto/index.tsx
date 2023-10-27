import * as React from 'react'
import Button from '@mui/material/Button'
import { Theme, styled } from '@mui/material/styles'
import Dialog from '@mui/material/Dialog'
import DialogTitle from '@mui/material/DialogTitle'
import DialogContent from '@mui/material/DialogContent'
import DialogActions from '@mui/material/DialogActions'
import IconButton from '@mui/material/IconButton'
import CloseIcon from '@mui/icons-material/Close'
import Typography from '@mui/material/Typography'
import { WButton } from '@/components'

interface CustomDialogProps {
    warning?: boolean
    title?: CustomText
    content?: string
    subtitle?: string
    size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl'
    btnText?: string
}

interface CustomText {
    textContent: string
    typeColor?: string
}

const BootstrapDialog = styled(Dialog)(() => ({
    // add border radius to the dialog's paper
    '& .MuiPaper-root': {
        borderRadius: '20px',
    },
}))

const sxDialogTitle = {
    p: 2,
    marginTop: 2,
    textAlign: 'center',
    fontSize: 20,
    fontWeight: 700,
}

const sxCloseButton = {
    position: 'absolute',
    right: 8,
    top: 8,
    color: (theme: Theme) => theme.palette.grey[500],
}

const CustomDialog: React.FC<CustomDialogProps> = ({ warning, content, title, subtitle, size, btnText }) => {
    const [open, setOpen] = React.useState(false)

    const handleClickOpen = () => {
        setOpen(true)
    }
    const handleClose = () => {
        setOpen(false)
    }

    return (
        <div>
            <Button variant="outlined" onClick={handleClickOpen}>
                Open dialog
            </Button>
            <BootstrapDialog onClose={handleClose} open={open} maxWidth={size} fullWidth={true}>
                <DialogTitle sx={sxDialogTitle} color={title?.typeColor}>
                    {warning && title?.textContent}
                </DialogTitle>

                <IconButton onClick={handleClose} sx={sxCloseButton}>
                    <CloseIcon />
                </IconButton>
                {subtitle && (
                    <Typography
                        sx={{
                            textAlign: 'center',
                            fontWeight: 700,
                            fontSize: 18,
                            marginTop: 2,
                            marginBottom: 3,
                        }}
                        gutterBottom
                    >
                        {subtitle}
                    </Typography>
                )}
                <DialogContent dividers>
                    <Typography
                        sx={{
                            whiteSpace: 'pre-line',
                            textAlign: 'center',
                            paddingX: 3,
                            fontWeight: 400,
                            fontSize: 14,
                        }}
                        gutterBottom
                    >
                        {content}
                    </Typography>
                </DialogContent>
                <DialogActions style={{ display: 'flex', justifyContent: 'center' }}>
                    <WButton text={btnText} size="large" />
                </DialogActions>
            </BootstrapDialog>
        </div>
    )
}

CustomDialog.defaultProps = {
    warning: false,
    title: {
        textContent: '!La foto no cumple con las especificaciones, subir otra foto¡',
        typeColor: 'error',
    },
    subtitle: 'Especificaciones de foto de portada',
    content:
        'Peso max (10 mb)\nTamaño max (851x315 px)\nLa plataforma acepta estas especificaciones para la foto de portada.',
    size: 'xs',
    btnText: 'Cargar Foto',
}

export default CustomDialog
