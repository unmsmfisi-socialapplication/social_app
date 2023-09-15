import React from "react";
import PersonOutlineIcon from "@mui/icons-material/PersonOutline";
import LockIcon from "@mui/icons-material/Lock";
import "../../resources/styles/LoginForm.scss";
import { TextField } from "@mui/material";

const LoginForm: React.FC = () => {
  return (
    <div className="login-form">
      {/* <TextField
        id="filled-basic"
        label="Filled"
        variant="filled"
        placeholder="Correo Electrónico"
        className="input-container email email-input"
      /> */}
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
