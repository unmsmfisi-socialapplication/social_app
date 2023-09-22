"use client";
import React from "react";
import Button from "@mui/material/Button";
import "./index.scss";

interface WButtonProps {
  typeColor?: "primary" | "secondary" | "disabled";
  text?: string;
  size?: "large";
  disabled?: boolean;
}


const WButton: React.FC<WButtonProps> = ({ disabled, typeColor, text, size }) => {
  const buttonClass = `button typeButton--${disabled ? "disabled" : typeColor}`;
  return (
    <Button
      style={{ minWidth: size === "large" ? "100%" : "auto" }}
      className={buttonClass}
      size={size}
      disabled={disabled}
    >
      {text}
    </Button>
  );
};

export default WButton;

WButton.defaultProps = {
  typeColor: "primary",
  text: "Button",
  disabled: false,
};
