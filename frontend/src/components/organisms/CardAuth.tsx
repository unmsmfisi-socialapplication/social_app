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
    display: "flex",
    flexDirection: "column",
  };

export default function CardAuth({ title, children, variant }: WCardProps) {
  return (
    <Card style={cardStyle} variant={variant}>
      <span style={{ fontSize: "60px" }}>{title}</span>
      <Box>{children}</Box>
    </Card>
  );
}
CardAuth.defaultProps = {
  typeColor: "primary",
  title: "title",
  variant: "outlined",
};
