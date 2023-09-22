"use client";
import React from "react";
import Button from "@mui/material/Button";
import "./index.scss";

interface WButtonProps {
  typeColor?: "primary" | "secondary" | "disabled";
  text?: string;
  size?: "large";
}

const WButton: React.FC<WButtonProps> = ({ typeColor, text, size }) => {
  return (
    <Button
      style={{ minWidth: size === "large" ? "100%" : "auto" }}
      className={`button button--${typeColor}`}
      size={size}
    >
      {text}
    </Button>
  );
};

export default WButton;

WButton.defaultProps = {
  typeColor: "primary",
  text: "Button",
};
