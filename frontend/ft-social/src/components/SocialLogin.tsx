import React from "react";
import GoogleIcon from "@mui/icons-material/Google";
import "./SocialLogin.css";

const SocialLogin: React.FC = () => {
  return (
    <div className="social-login">
      <button className="google-button">
        <GoogleIcon /> Continuar con Google
      </button>
    </div>
  );
};

export default SocialLogin;
