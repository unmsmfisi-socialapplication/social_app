"use client";
import React from "react";
import { Box, Card } from "@mui/material";

interface WCardProps {
  typeColor?: "primary" | "secondary";
  title?: string;
  size?: "large";
  variant?: "outlined" | "elevation";
  children?: React.ReactNode;
}
const cardStyle: React.CSSProperties = {
    border: "none",
  };

export default function CardAuth({ title, children, variant }: WCardProps) {
  return (
    <Card style={cardStyle} variant={variant}>
      <span style={{ fontSize: "75px"}}>{title}</span>
      <Box style={{marginTop : "10px",display:"flex" , flexDirection:"column", gap : "0.8rem"}} >{children}</Box>
    </Card>
  );
}
CardAuth.defaultProps = {
  typeColor: "primary",
  title: "title",
  variant: "outlined",
};
