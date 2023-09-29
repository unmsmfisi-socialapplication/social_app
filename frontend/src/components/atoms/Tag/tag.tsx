"use client";
import React from 'react'
import { Box, SvgIcon, SvgIconProps } from '@mui/material'
import AllInclusive from '@mui/icons-material/AllInclusive';
import "./index.scss"

interface WTagProps{
  icon: React.ComponentType<SvgIconProps>;
  text?:string;
  isActive?: boolean;
  path?: string;
}


const WTag:React.FC<WTagProps> = ({icon,text,isActive,path}) => {
  return (
    <Box onClick={()=>console.log(`Ir a la ruta ${text}`)} className={isActive? "tagLink tagLink--active": "tagLink"} >
      <SvgIcon component={icon}></SvgIcon>
      {text}
    </Box>
  )
}

export default WTag

WTag.defaultProps={
  icon : AllInclusive,
  isActive: false,
  text: "TagLink",
  path:"/",
}