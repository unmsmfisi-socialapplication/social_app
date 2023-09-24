import React from 'react';
import TextField from '@mui/material/TextField';
import InputAdornment from '@mui/material/InputAdornment';

interface WInputProps {
  typeColor?: 'primary' | 'secondary';
  size?: 'small' | 'medium';
  variant?: 'standard' | 'filled' | 'outlined';
  fullWidth?: boolean;
  type?: string;
}

import { styled } from '@mui/material/styles';

const CustomizedTextField = styled(TextField)`
    
  .MuiInputBase-input{
    color: #000;
    font-size: 22px;
    text-align: center;
    font-weight: 600;

    width: 60px;
    display: flex;
    justifyContent: center;
    alignItems: center;
  }

  .css-1t8l2tu-MuiInputBase-input-MuiOutlinedInput-input{
    padding: 10px 7.5px
  }
  
`;


const WInputCode: React.FC<WInputProps> = ({
  typeColor = 'primary',
  size = 'medium',
  variant = 'outlined',
  fullWidth = false,
  type = 'text',
}) => {

  return (
    <CustomizedTextField 
    id="outlined-basic" 
    variant="outlined" 
    inputProps={{ maxLength: 1 }}
    />
  );
};

export default WInputCode;