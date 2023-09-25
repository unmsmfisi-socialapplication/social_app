import * as React from 'react';
import Link from '@mui/material/Link';


interface BasicsLinkProps {
  underline?: "none" | "always" | "none";
  text?: string;
  href?: string;
  displayType?: "flex" | "inline-flex";
}

const WBasicsLink: React.FC<BasicsLinkProps> = ({ underline, text, href,displayType }) => {
  const linkStyle = {  
    right: 100,  
    gap: 2,
    justifyContent: 'space-end',
  };

  return (
    <Link
      sx={linkStyle}
      href={href}
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
    displayType: "flex",
};