import * as React from 'react';

import Link from '@mui/material/Link';
import "./link.scss"


interface BasicsLinkProps {
    underline?:"none" | "always" | "none";
    text?: string;
    href?: "#";
    displayType?: "flex" | "inline-flex";
}


const BasicsLink: React.FC<BasicsLinkProps> = ({underline, text,href,displayType}) =>{    
    const linkStyle={
        display : displayType,
        gap:2,
        flexWrap: 'wrap',
        justifyContent:'flex-end'
    }
    return(   
        <Link 
        style={linkStyle}   
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