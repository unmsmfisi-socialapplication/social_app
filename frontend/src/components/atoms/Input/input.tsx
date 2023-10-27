import React, { useState } from 'react'
import TextField from '@mui/material/TextField'
import InputAdornment from '@mui/material/InputAdornment'
import IconButton from '@mui/material/IconButton'
import Visibility from '@mui/icons-material/Visibility'
import VisibilityOff from '@mui/icons-material/VisibilityOff'

interface WInputProps {
    value?: string
    typeColor?: 'primary' | 'secondary'
    dataTestid?: string
    name?: string
    icon?: React.ReactElement
    placeholder?: string
    size?: 'small' | 'medium'
    variant?: 'standard' | 'filled' | 'outlined'
    fullWidth?: boolean
    type?: 'text' | 'password'
    error?: boolean
    errorMessage?: string
    onChange?: (event: React.ChangeEvent<HTMLInputElement>) => void
    onBlur?: (event: React.FocusEvent<HTMLInputElement>) => void
}

const WInput: React.FC<WInputProps> = ({
    typeColor = 'primary',
    dataTestid,
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
    const [showPassword, setShowPassword] = useState(false)

    const togglePasswordVisibility = () => {
        setShowPassword((prevShowPassword) => !prevShowPassword)
    }

    return (
        <TextField
            value={value}
            name={name}
            data-testid={dataTestid}
            onChange={onChange}
            onBlur={onBlur}
            fullWidth={fullWidth}
            variant={variant}
            size={size}
            color={typeColor}
            placeholder={placeholder}
            type={type === 'password' && !showPassword ? 'password' : 'text'}
            error={error}
            InputProps={{
                endAdornment: (
                    <InputAdornment position="end">
                        {type === 'password' ? (
                            <IconButton onClick={togglePasswordVisibility} edge="end">
                                {showPassword ? <VisibilityOff /> : <Visibility />}
                            </IconButton>
                        ) : (
                            icon && <InputAdornment position="end">{icon}</InputAdornment>
                        )}
                    </InputAdornment>
                ),
            }}
            helperText={error && errorMessage}
        />
    )
}

export default WInput
