"use client";
import React from "react";
import AllInclusiveIcon from "@mui/icons-material/AllInclusive";
import { SvgIcon, SvgIconProps } from "@mui/material";
import "./index.scss";

interface WCircleIconProps {
  typeColor?: "primary" | "secondary" | "comment";
  icon: React.ComponentType<SvgIconProps>;
  iconSize?: number;
}

const WCircleIcon: React.FC<WCircleIconProps> = ({
  typeColor,
  icon,
  iconSize,
}) => {
  const borderStyle = {
    border: "2mm solid white",
    borderRadius: "50%", //To make sure the edge is circular
  };

  return (
    <SvgIcon
      component={icon}
      sx={{ fontSize: iconSize, ...borderStyle }} // Add the border to the style
      className={`circleIcon circleIcon--${typeColor}`}
    ></SvgIcon>
  );
};

export default WCircleIcon;

WCircleIcon.defaultProps = {
  typeColor: "primary",
  icon: AllInclusiveIcon,
  iconSize: 40,
};
