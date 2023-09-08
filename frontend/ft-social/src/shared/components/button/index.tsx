import React from 'react';
import Button from '@mui/material/Button';
import './index.less';

type WButtonProps = {
    typeColor: 'primary' | 'secondary' | 'tertiary';
    text: string; // Definir que text es de tipo string
    size: 'small' | 'medium' | 'large'  ;
};
  
  const WButton: React.FC<WButtonProps> = ({ typeColor , text ,size}) => {
    return (
      <Button
        className={`custom-button ${typeColor}`}
        size={size}
      >
        {text}
      </Button>
    );
  };
  
export default WButton;

WButton.defaultProps = {
    typeColor: 'primary',
    text: 'Default text',
    size: undefined,
};