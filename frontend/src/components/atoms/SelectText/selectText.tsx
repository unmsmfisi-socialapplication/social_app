import React from 'react'
import TextField from '@mui/material/TextField'
import MenuItem from '@mui/material/MenuItem'

interface WSelectedTextProps {
    disabled?: boolean
    label?: string
    size?: 'small' | 'medium'
    variant?: 'standard' | 'filled' | 'outlined'
    options: { value: string }[]
    fullWidth?: boolean
    error?: boolean
    errorMessage?: string
    iconComponent?: React.ElementType
}

const WSelectedText: React.FC<WSelectedTextProps> = ({
    label,
    disabled = false,
    size = 'medium',
    variant = 'outlined',
    options = [],
    fullWidth = false,
    error = false,
    errorMessage = '',
    iconComponent: IconComponent = undefined,
}) => {
    return (
        <TextField
            select
            defaultValue=""
            label={label}
            disabled={disabled}
            fullWidth={fullWidth}
            variant={variant}
            size={size}
            error={error}
            helperText={error && errorMessage}
            SelectProps={{
                IconComponent: IconComponent,
            }}
        >
            {options.map((option) => (
                <MenuItem key={option.value} value={option.value}>
                    {option.value}
                </MenuItem>
            ))}
        </TextField>
    )
}

export default WSelectedText

WSelectedText.defaultProps = {
    label: 'Select',
    disabled: false,
    variant: 'outlined',
    fullWidth: true,
    size: 'medium',
}
