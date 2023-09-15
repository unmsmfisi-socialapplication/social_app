import React from "react";
/* import GoogleIcon from "@mui/icons-material/Google"; */
import "../../resources/styles/SocialLogin.scss";

const SocialLogin: React.FC = () => {
  return (
    <div className="social-login">
      {/* <button className="google-button">
        <GoogleIcon /> Continuar con Google
      </button> */}
      <a className="google"></a>
      <a className="threads"></a>
      <a className="mastodon"></a>
    </div>
  );
};

export default SocialLogin;
