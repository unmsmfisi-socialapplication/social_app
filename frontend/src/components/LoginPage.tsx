import React from "react";
import LoginForm from "./LoginForm";
import SocialLogin from "./SocialLogin";

import "../../resources/styles/LoginPage.scss";

const LoginPage: React.FC = () => {
  return (
    <div className="login-page">
      <div className="login-container">
        <div className="welcome-text">
          <h1 className="welcome-title">Bienvenido</h1>
          <p className="welcome-paragraph">Inicia sesión para continuar</p>
        </div>
        <LoginForm />
        <div className="forgot-password">
          <a href="#">¿Olvidaste la contraseña?</a>
        </div>
        <button className="login-button login-btn">Inicio de sesión</button>
        <div className="separator">
          <div className="separator-line"></div>
          <div className="separator-circle"></div>
          <div className="separator-line"></div>
        </div>
        <button className="create-account-button">
          Crear una nueva cuenta
        </button>
        <p className="continuar">Continuar con</p>
        <SocialLogin />
      </div>
    </div>
  );
};

export default LoginPage;
