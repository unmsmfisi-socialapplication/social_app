'use client';
import * as React from 'react';
import Link from '@mui/material/Link';
import { useRouter } from 'next/router';


interface BasicsLinkProps {
  underline?: "none" | "always" | "none";
  text?: string;
  dataTestid?: string;
  href?: string;
  displayType?: "flex" | "inline-flex";
}

const WBasicsLink: React.FC<BasicsLinkProps> = ({ underline,  text,dataTestid, href,displayType }) => {
  const router = useRouter()

  const linkStyle = {  
    right: 100,  
    gap: 2,
    justifyContent: 'space-end',
  };

  return (
      <Link
      data-testid={dataTestid}
      sx={linkStyle}
      onClick={()=> router.push(href || "/")}
      underline={underline}
       >
      {text}
      </Link>
  );
};



export default WBasicsLink;

WBasicsLink.defaultProps = {
    underline:"always",
    text: "link",
    href: "#",
    displayType:"flex",
};