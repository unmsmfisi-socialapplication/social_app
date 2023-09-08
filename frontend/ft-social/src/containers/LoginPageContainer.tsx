import React, { useState } from "react";
import LoginPage from "../components/LoginPage";

const LoginPageContainer: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async () => {
    // Call an authentication API or perform login logic
    // Handle successful login
    // Handle login error
  };

  return (
    <LoginPage
      email={email} //TODO: fix
      password={password}
      setEmail={setEmail}
      setPassword={setPassword}
      onLogin={handleLogin}
    />
  );
};

export default LoginPageContainer;
