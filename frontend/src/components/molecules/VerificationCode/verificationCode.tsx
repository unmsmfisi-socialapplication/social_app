import React from "react";
import WInputCode from "@/components/atoms/InputCode/inputCode";

interface WVerificationCodeProps {
  typeColor?: "primary" | "secondary";
  placeholder?: string;
  size?: "small" | "medium";
  variant?: "standard" | "filled" | "outlined";
  fullWidth?: boolean;
  type?: "text" | "password";
}

const WVerificationCode: React.FC<WVerificationCodeProps> = ({
  typeColor = "primary",
  size = "medium",
  variant = "outlined",
  fullWidth = false,
  type = "text",
}) => {
  const divStyle = {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    gap: "30px",
  };

  return (
    <div style={divStyle}>
      <WInputCode />
      <WInputCode />
      <WInputCode />
      <WInputCode />
    </div>
  );
};

export default WVerificationCode;