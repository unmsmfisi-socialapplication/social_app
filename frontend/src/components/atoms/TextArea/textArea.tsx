import React from 'react'
import { styled } from '@mui/material';
import { TextareaAutosize as BaseTextareaAutosize } from '@mui/base/TextareaAutosize';

const blue = {
    100: '#DAECFF',
    200: '#b6daff',
    400: '#3399FF',
    500: '#007FFF',
    600: '#0072E5',
    900: '#003A75',
  };

  const Textarea = styled(BaseTextareaAutosize)(
    ({ theme }) => `
    font-family: IBM Plex Sans, sans-serif;
    color: #0F1419;
    font-size: 20px;
    font-weight: 400;
    line-height: 1.5;
    padding: 8px 12px;
    border-radius: 8px;
    background: #fff;
    border: none;
    box-shadow: none;
    width: 95%;

    &:hover {
      border-color: #fff;
    }

    &:focus {
      border-color: #fff;
      box-shadow: 0 0 0 1px ${theme.palette.mode === 'dark' ? blue[600] : blue[200]};
    }

    // firefox
    &:focus-visible {
      outline: 0;
    }
  `,
  );

interface TextAreaProps {
    textAreaValue: string;
    handleChangeTextArea: any;
    placeholder: string
}

const WTextArea: React.FC<TextAreaProps> = ({ textAreaValue, handleChangeTextArea, placeholder}) => {
    return (
        <Textarea id="textArea" name="textArea" aria-label="minimum height" value={textAreaValue} placeholder={placeholder} onChange={handleChangeTextArea}/>
    )
}

export default WTextArea;

WTextArea.defaultProps = {
    textAreaValue: "",
    placeholder: ""
}