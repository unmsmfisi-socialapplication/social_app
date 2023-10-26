'use client';
import * as React from 'react';
import LinkMaterial from '@mui/material/Link';
import Link from "next/link"


interface BasicsLinkProps {
  underline?: "none" | "always" | "none";
  text?: string;
  dataTestid?: string;
  href?: string;
  displayType?: "flex" | "inline-flex";
}

const WBasicsLink: React.FC<BasicsLinkProps> = ({ underline,  text,dataTestid, href,displayType }) => {

  const linkStyle = {  
    right: 100,  
    gap: 2,
    justifyContent: 'space-end',
  };

  return (
  <Link data-testid={dataTestid} href={href || "/"} >
      <LinkMaterial
      sx={linkStyle}
      underline={underline}
       >
      {text}
      </LinkMaterial>
  </Link>
  );
};



export default WBasicsLink;

WBasicsLink.defaultProps = {
    underline:"always",
    text: "link",
    href: "#",
    displayType: "flex",
};