import React from "react";
import PersonOutlineIcon from "@mui/icons-material/PersonOutline";
import LockIcon from "@mui/icons-material/Lock";
import "./LoginForm.css";

const LoginForm: React.FC = () => {
  return (
    <div className="login-form">
      <div className="input-container">
        <label htmlFor="email">
          <input
            type="text"
            id="email"
            placeholder="Correo Electrónico"
            className="email-input"
          />
        </label>
        <PersonOutlineIcon />
      </div>
      <div className="input-container">
        <label htmlFor="password">
          <input
            type="password"
            id="password"
            placeholder="Contraseña"
            className="email-input"
          />
        </label>
        <LockIcon />
      </div>
    </div>
  );
};

export default LoginForm;
