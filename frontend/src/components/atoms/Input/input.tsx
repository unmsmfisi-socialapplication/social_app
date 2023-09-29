import React from 'react';
import TextField from '@mui/material/TextField';
import InputAdornment from '@mui/material/InputAdornment';

interface WInputProps {
  value?: string;
  typeColor?: 'primary' | 'secondary';
  name?: string;
  icon?: React.ReactElement;
  placeholder?: string;
  size?: 'small' | 'medium';
  variant?: 'standard' | 'filled' | 'outlined';
  fullWidth?: boolean;
  type?: 'text' | 'password';
  error?: boolean;
  errorMessage?: string;
  onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void;
  onBlur?: (event: React.FocusEvent<HTMLInputElement>) => void;
}

const WInput: React.FC<WInputProps> = ({
  typeColor = 'primary',
  value,
  name,
  icon,
  placeholder = 'Placeholder',
  size = 'medium',
  variant = 'outlined',
  fullWidth = false,
  type = 'text',
  error = false,
  errorMessage = '',
  onChange = () => {},
  onBlur = () => {},
}) => {
  return (
    <TextField
      value={value}
      name={name}
      onChange={onChange}
      onBlur={onBlur}
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
