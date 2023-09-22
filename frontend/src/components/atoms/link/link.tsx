import * as React from 'react';
import Link from '@mui/material/Link';

interface BasicsLinkProps {
  underline?: "none" | "always" | "none";
  text?: string;
  href?: string;
  displayType?: "flex" | "inline-flex";
}

const BasicsLink: React.FC<BasicsLinkProps> = ({ underline, text, href,displayType }) => {
  const linkStyle = {
    
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



export default BasicsLink;

BasicsLink.defaultProps = {
    underline:"always",
    text: "link",
    href: "#",
    displayType: "flex",
};