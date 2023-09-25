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
  error?: boolean;
  errorMessage?: string;
}

const WInput: React.FC<WInputProps> = ({
  typeColor = 'primary',
  icon,
  placeholder = 'Placeholder',
  size = 'medium',
  variant = 'outlined',
  fullWidth = false,
  type = 'text',
  error = false,
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
      error={error}
      InputProps={{
        endAdornment: icon && (
          <InputAdornment position="end">{icon}</InputAdornment>
        ),
      }}
      helperText={error && errorMessage}
    />
  );
};

export default WInput;
