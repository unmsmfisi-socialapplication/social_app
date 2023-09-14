"use client";
import React from "react";
import TextField from '@mui/material/TextField';
import InputAdornment from '@mui/material/InputAdornment';
import "./index.scss";
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { IconDefinition } from '@fortawesome/fontawesome-svg-core';


interface WInputProps {
    typeColor?: "primary" | "secondary";
    icon?: IconDefinition;
    placeholder?: string;
    size?: 'small' | 'medium';
    
  }

const WInput: React.FC<WInputProps> = ({icon , placeholder, size}) => {
  return (
    <TextField
        variant="outlined"
        fullWidth
        size={size}
        InputProps={{
          startAdornment: icon ? (
            <InputAdornment position="start">
              <FontAwesomeIcon icon={icon} />
            </InputAdornment>
          ) : undefined,
        }}
        placeholder={placeholder}
    />
  );
};

export default WInput;