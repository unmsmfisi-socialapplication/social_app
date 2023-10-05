"use client";
import React from "react";
import Button from "@mui/material/Button";
import "./index.scss";

export interface WButtonProps {
  id?: string;
  typeColor?: "primary" | "secondary" | "disabled";
  type?: "submit" ;
  text?: string;
  size?: "large";
  disabled?: boolean;
}


const WButton: React.FC<WButtonProps> = ({ id , disabled, typeColor, text, type, size }) => {
  const buttonClass = `button typeButton--${disabled ? "disabled" : typeColor}`;
  return (
    <Button
      id={id}
      style={{ minWidth: size === "large" ? "100%" : "auto" }}
      type = {type}
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
