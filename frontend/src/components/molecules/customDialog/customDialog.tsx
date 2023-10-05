import * as React from 'react';
import Button from '@mui/material/Button';
import { styled } from '@mui/material/styles';
import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import IconButton from '@mui/material/IconButton';
import CloseIcon from '@mui/icons-material/Close';
import Typography from '@mui/material/Typography';
import WButton, { WButtonProps } from "@/components/atoms/Button/button";

const BootstrapDialog = styled(Dialog)(({ theme }) => ({
  '& .MuiDialogContent-root': {
    padding: theme.spacing(2),
  },
  '& .MuiDialogActions-root': {
    padding: theme.spacing(1),
  },
  // agregamos mas border radius para que se vea mas redondo
  '& .MuiPaper-root': {
    borderRadius: '20px',
  },
}));

export interface CustomDialogProps {
  title: CustomText,
  content: CustomText,
  subtitle?: CustomText,
  size: "xs" | "sm" | "md" | "lg" | "xl",
  buttonProps?: WButtonProps
}

export interface CustomText {
  textContent: string;
  typeColor?: string;
  fontSize?: number;
  fontWeight?: number;
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
        aria-labelledby="customized-dialog-title"
        open={open}
        maxWidth={size}
        fullWidth={true}
      >
        <DialogTitle
          sx={{
            m: 0,
            p: 2,
            textAlign: 'center',
            fontSize: title.fontSize,
            fontWeight: title.fontWeight,
            paddingX: '25px' // Espacio adicional para el botón de cerrar
          }}
          id="customized-dialog-title"
          color={title.typeColor}
        >
          {title.textContent}
        </DialogTitle>

        <IconButton
          aria-label="close"
          onClick={handleClose}
          sx={{
            position: 'absolute',
            right: 8,
            top: 8,
            color: (theme) => theme.palette.grey[500],
          }}
        >
          <CloseIcon />
        </IconButton>

        <DialogContent dividers>
          {subtitle && (
            <Typography sx={{ whiteSpace: 'pre-line', textAlign: 'center', paddingX: 3, fontWeight: subtitle.fontWeight, fontSize: subtitle.fontSize, color: subtitle.typeColor }} gutterBottom>
              {subtitle.textContent}
            </Typography>
          )}
          <Typography sx={{ whiteSpace: 'pre-line', textAlign: 'center', paddingX: 3, fontWeight: content.fontWeight, fontSize: content.fontSize, color: content.typeColor }} gutterBottom>
            {content.textContent}
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

export default CustomDialog;