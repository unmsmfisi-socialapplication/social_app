import React from 'react';
import WInputCode from './../../../components/atoms/InputCode/inputCode';

interface WVerificationCodeProps {
  typeColor?: 'primary' | 'secondary';
  placeholder?: string;
  size?: 'small' | 'medium';
  variant?: 'standard' | 'filled' | 'outlined';
  fullWidth?: boolean;
  type?: 'text' | 'password';
  dataTestid?: string;
}

const WVerificationCode: React.FC<WVerificationCodeProps> = ({
  typeColor = 'primary',
  size = 'medium',
  variant = 'outlined',
  fullWidth = false,
  type = 'text',
  dataTestid,
}) => {
  const divStyle = {
    display: 'flex',
    alignItems: 'center',
    justifyContent: 'center',
    gap: '30px',
  };

  return (
    <div style={divStyle} data-testid={dataTestid}>
      <WInputCode  />
      <WInputCode  />
      <WInputCode  />
      <WInputCode  />
    </div>
  );
};

export default WVerificationCode;
