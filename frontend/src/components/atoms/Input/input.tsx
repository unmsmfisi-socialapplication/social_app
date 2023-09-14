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
}

const WInput: React.FC<WInputProps> = ({
  typeColor = 'primary',
  icon,
  placeholder = 'Placeholder',
  size = 'medium',
  variant = 'filled',
  fullWidth = false,
}) => {
  return (
    <TextField
      fullWidth={fullWidth}
      variant={variant}
      size={size}
      color={typeColor}
      placeholder={placeholder}
      InputProps={{
        endAdornment:  icon && (
          <InputAdornment position="end">{icon}</InputAdornment>
        ) ,
      }}
    />
  );
};

export default WInput;
