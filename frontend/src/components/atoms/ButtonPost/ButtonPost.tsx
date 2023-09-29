"use client";
import React from "react";
import Button from "@mui/material/Button";
import "./index.scss";

interface WButtonPostProps {
    typeColor?: "primary" | "secondary" | "disabled";
    text?: string;
    size?: "large";
    disabled?: boolean;
    onChange?: (event: React.ChangeEvent<HTMLButtonElement>) => void;
}

const WButtonPost: React.FC<WButtonPostProps> = ({ disabled, typeColor, text, size, onChange }) => {
    const buttonClass = `button typeButton--${disabled ? "disabled" : typeColor}`;
  
    const handleChange = (event: React.ChangeEvent<HTMLButtonElement>) => {
      if (onChange) {
        onChange(event);
      }
    };
  
    return (
      <Button
        style={{ minWidth: size === "large" ? "100%" : "auto" }}
        className={buttonClass}
        size={size}
        disabled={disabled}
        onChange={handleChange}
      >
        {text}
      </Button>
    );
  };
  

export default WButtonPost;

WButtonPost.defaultProps = {
  typeColor: "primary",
  text: "Button",
  disabled: false,
};
