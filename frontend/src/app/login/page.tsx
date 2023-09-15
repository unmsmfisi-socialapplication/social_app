"use client";

import React, { useState } from "react";
import Layout from "./layout";
import LoginPage from "../../components/LoginPage";

const LoginPageHome: React.FC = () => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleLogin = async () => {
    // Call an authentication API or perform login logic
    // Handle successful login
    // Handle login error
  };
  return (
    <Layout>
      <LoginPage
        email={email}
        password={password}
        setEmail={setEmail}
        setPassword={setPassword}
        onLogin={handleLogin}
      />
    </Layout>
  );
};

export default LoginPageHome;
