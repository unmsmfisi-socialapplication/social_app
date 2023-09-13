import React, { ChangeEvent } from 'react';
import "./index.scss";
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {faUser} from '@fortawesome/free-solid-svg-icons'
<link rel="stylesheet" href="https://fonts.googleapis.com/css?family=Poppins&display=swap" />

interface InputProps {
  value: string;
  onChange: (e: ChangeEvent<HTMLInputElement>) => void;
}

const Input: React.FC<InputProps> = ({ value, onChange }) => {
  return (
    <div className="input-container">
      <input
        type="input"
        name="input"
        placeholder="Correo Electrónico"
        value={value}
        onChange={onChange}
        className="input"
      />
       <div className="icon">
        <FontAwesomeIcon icon={faUser} />
      </div>
    </div>
  );
};

export default Input;