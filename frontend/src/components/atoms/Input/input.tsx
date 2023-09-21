import React from 'react';
import TextField from '@mui/material/TextField';
import InputAdornment from '@mui/material/InputAdornment';

interface WInputProps {
  typeColor?: 'primary' | 'secondary';
  icon?: React.ReactElement;
  placeholder?: string;
  size?: 'small' | 'medium';
  variant?: 'standard' | 'filled' | 'outlined';
  fullWidth?: boolean;
  type?: 'text' | 'password';
  error?: boolean; // Propiedad para controlar el error
  errorMessage?: string; // Nuevo: Mensaje de error
}

const WInput: React.FC<WInputProps> = ({
  typeColor = 'primary',
  icon,
  placeholder = 'Placeholder',
  size = 'medium',
  variant = 'outlined',
  fullWidth = false,
  type = 'text',
  error = false, // Valor por defecto: no hay error
  errorMessage = '',
}) => {
  return (
    <TextField
      fullWidth={fullWidth}
      variant={variant}
      size={size}
      color={typeColor}
      placeholder={placeholder}
      type={type}
      error={error} // Pasar la propiedad de error a TextField
      InputProps={{
        endAdornment: icon && (
          <InputAdornment position="end">{icon}</InputAdornment>
        ),
      }}
      label={error ? errorMessage : ''} // Mostrar mensaje de error si hay error
      helperText={error ? '' : errorMessage} // Mostrar mensaje de error si hay error
    />
  );
};

export default WInput;
