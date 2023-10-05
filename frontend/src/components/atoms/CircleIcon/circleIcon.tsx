"use client";
import React from "react";
import AllInclusiveIcon from "@mui/icons-material/AllInclusive";
import { SvgIcon, SvgIconProps } from "@mui/material";
import "./index.scss";

interface WCircleIconProps {
  typeColor?: "primary" | "secondary";
  icon: React.ComponentType<SvgIconProps>;
  iconSize?: number;
}

const WCircleIcon: React.FC<WCircleIconProps> = ({
  typeColor,
  icon,
  iconSize,
}) => {
  return (
    <SvgIcon
      component={icon}
      sx={{ fontSize: iconSize }}
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
