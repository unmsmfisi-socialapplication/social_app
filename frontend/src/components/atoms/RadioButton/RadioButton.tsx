import * as React from 'react'
import Radio from '@mui/material/Radio'

interface WRadioButtonProps {
    selectedValue: string
    handleChange: (event: React.ChangeEvent<HTMLInputElement>) => void
    value: string
}

const WRadioButton: React.FC<WRadioButtonProps> = ({ selectedValue, handleChange, value }) => {
    return (
        <Radio
            checked={selectedValue === value}
            onChange={handleChange}
            value={value}
            name="color-radio-button-demo"
            inputProps={{ 'aria-label': value }}
        />
    )
}

export default WRadioButton

WRadioButton.defaultProps = {
    selectedValue: '',
    handleChange: () => {},
    value: 'a',
}
