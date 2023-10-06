import * as React from 'react';
import Button from '@mui/material/Button';
import { Theme, styled } from '@mui/material/styles';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';
import Typography from '@mui/material/Typography';
import WButton, { WButtonProps } from "@/components/atoms/Button/button";


export interface CustomDialogProps {
  title?: CustomText,
  content?: CustomText,
  subtitle?: CustomText,
  size?: "xs" | "sm" | "md" | "lg" | "xl",
  buttonProps?: WButtonProps
}

export interface CustomText {
  textContent: string;
  typeColor?: string;
}

const BootstrapDialog = styled(Dialog)(() => ({
  // agregamos mas border radius para que se vea mas redondo
  '& .MuiPaper-root': {
    borderRadius: '20px',
  },
}));

const sxDialogTitle = {
  m: 0,
  p: 2,
  textAlign: 'center',
  fontSize: 20,
  fontWeight: 700,
  paddingX: '25px'
}

const sxCloseButton = {
  position: 'absolute',
  right: 8,
  top: 8,
  color: (theme: Theme) => theme.palette.grey[500],
}

const CustomDialog: React.FC<CustomDialogProps> = ({ content, title, subtitle, size, buttonProps }) => {
  const [open, setOpen] = React.useState(false);

  const handleClickOpen = () => {
    setOpen(true);
  };
  const handleClose = () => {
    setOpen(false);
  };

  return (
    <div>
      <Button variant="outlined" onClick={handleClickOpen}>
        Open dialog
      </Button>
      <BootstrapDialog
        onClose={handleClose}
        open={open}
        maxWidth={size}
        fullWidth={true}
      >
        <DialogTitle
          sx={sxDialogTitle}
          color={title?.typeColor}
        >
          {title?.textContent}
        </DialogTitle>

        <IconButton
          onClick={handleClose}
          sx={sxCloseButton}
        >
          <CloseIcon />
        </IconButton>

        <DialogContent dividers>
          {subtitle && (
            <Typography sx={{textAlign: 'center', fontWeight: 700, fontSize: 16, color: subtitle.typeColor }} gutterBottom>
              {subtitle.textContent}
            </Typography>
          )}
          <Typography sx={{ whiteSpace: 'pre-line', textAlign: 'center', paddingX: 3, fontWeight: 400, fontSize: 14, color: content?.typeColor }} gutterBottom>
            {content?.textContent}
          </Typography>
        </DialogContent>
        {buttonProps && (
          <DialogActions>
            <WButton text={buttonProps.text} size={buttonProps.size} typeColor={buttonProps.typeColor}  />
          </DialogActions>
        )}

      </BootstrapDialog>
    </div>
  );
}

CustomDialog.defaultProps = {
    title: {
      textContent: "!La foto no cumple con las especificaciones, subir otra foto¡",
      typeColor: "error",
    },
    content: {
      textContent: "Peso max (10 mb)\nTamaño max (851x315 px)\nLa plataforma acepta estas especificaciones para la foto de portada.",
    },
    subtitle: {
      textContent: "Especificaciones de foto de portada",
    },
    size: "xs",
    buttonProps: {
      text: "Cargar Foto",
      typeColor: "primary",
      size: "large",
    }
}

export default CustomDialog;