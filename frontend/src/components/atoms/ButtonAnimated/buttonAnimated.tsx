import React from "react";
import Button from "@mui/material/Button";
import "./ButtonAnimated.scss"; // Importa el archivo SCSS que contiene los estilos

interface WButtonMotionProps {
  id?: string;
  type?: "submit";
  text?: string;
  size?: "large";
  disabled?: boolean;
}

const WButtonMotion: React.FC<WButtonMotionProps> = ({ id, type, text, size }) => {
  return (
    <Button
      id={id}
      style={{ minWidth: size === "large" ? "100%" : "auto" }}
      type={type}
      className="button-animated" 
    >
      {text}
    </Button>
  );
};

WButtonMotion.defaultProps = {
  text: "Button",
};

export default WButtonMotion;
