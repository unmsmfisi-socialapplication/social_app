import React from "react";
import { WCircleIcon } from '@/components';
import FavoriteIcon from '@mui/icons-material/Favorite';
import ModeCommentIcon from '@mui/icons-material/ModeComment';
import CachedIcon from '@mui/icons-material/Cached';
import BarChartIcon from '@mui/icons-material/BarChart';
import StarIcon from '@mui/icons-material/Star';
import ShareIcon from '@mui/icons-material/Share';
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
interface User {
    usuario?: string;
    name?: string; 
}

interface WCommentProps {
  user?: User;
  sendUser?: string;
  typeColor?: "primary" | "secondary" ;
  comment?: string;
  fullWidth?: boolean;
  countHeart?: number;
  countComment?: number;
  countShareds?: number;
  countStatistics?: number;
  countStars?: number;
}

const WComment: React.FC<WCommentProps> = ({
  user = {
    usuario: "millys123",
    name: "Milly"
  },
  sendUser = "janedoel123",
  typeColor = "primary",
  comment = "Lorem ipsum dolor sit amet consectetur. Congue et nunc sed nascetur malesuada",
  fullWidth = false,
  countHeart = 12,
  countComment = 5,
  countShareds = 6,
  countStatistics = 126,
  countStars = 23
}) => {
  const divStyle = {
    display: "grid",
    alignItems: "center",
    gap: "10px",
    width: "100%",
    minWidth: "500px"
  };

  const divIcon = {
    display: "inherit", 
    alignItems: "center", 
    gap: "5px"
  }

  const divSectionIcon = {
    display: "flex", 
    justifyContent: "space-between", 
    width: "100%"
  }

  const divUser = {
    display: "flex", 
    flex: "start", 
    gap: "10px", 
    alignItems: "center"
  }
  return (
    <div style={divStyle}>
      <div style={divUser}>
        <WCircleIcon 
            typeColor="comment"
            iconSize={50}
            icon={AccountCircleIcon}
        />
        <div style={{display: "inherit", flexDirection: "column"}}>
            <span style={{color: "#878787"}}><strong style={{color: "#000"}}>{user.name}</strong> @{user.usuario}</span>
            <span style={{color: "#878787"}}>Respondiendo a <a href="#" style={{textDecoration: "none", color: "#007AFF"}}>@{sendUser}</a></span>
        </div>
      </div>
      <div>
        {comment}
      </div>
      <div style={divSectionIcon}>
        <div style={divIcon}>
            <WCircleIcon
                typeColor="comment"
                iconSize={18}
                icon={FavoriteIcon}
            />
            <span>{countHeart}</span>
        </div>
        <div style={divIcon}>
            <WCircleIcon
                typeColor="comment"
                iconSize={18}
                icon={ModeCommentIcon}
            />
            <span>{countComment}</span>
        </div>
        <div style={divIcon}>
            <WCircleIcon
                typeColor="comment"
                iconSize={18}
                icon={CachedIcon}
            />
            <span>{countShareds}</span>
        </div>
        <div style={divIcon}>
            <WCircleIcon
                typeColor="comment"
                iconSize={18}
                icon={BarChartIcon}
            />
            <span>{countStatistics}</span>
        </div>
        <div style={divIcon}>
            <WCircleIcon
                typeColor="comment"
                iconSize={18}
                icon={StarIcon}
            />
            <span>{countStars}</span>
        </div>
        <div style={divIcon}>
            <WCircleIcon
                typeColor="comment"
                iconSize={18}
                icon={ShareIcon}
            />
        </div>
      </div>
    </div>
  );
};

export default WComment;