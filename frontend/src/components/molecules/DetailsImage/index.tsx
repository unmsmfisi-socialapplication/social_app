import React from "react";
import WCircleIcon from "@/components/atoms/CircleIcon/circleIcon";

interface WDetailsImageProps {
  accountName: string;
  name: string;
  icon: React.ComponentType;
}

const WDetailsImage: React.FC<WDetailsImageProps> = ({
  accountName,
  name,
  icon: IconComponent,
}) => {
  const divStyle = {
    display: "flex",
    alignItems: "flex-start",
    justifyContent: "flex-start",
    gap: "11px",
  };
  const divStyleText = {
    margin: "0",
    fontFamily: "Poppins",
    fontWeight: "600",
    fontSize: "14px",
    lineHeight: "19px",
    color: "#000",
  };
  const pStyleName = {
    color: "rgba(0, 0, 0, 0.40)",
  };

  return (
    <div style={divStyle}>
      <WCircleIcon iconSize={30} typeColor={"secondary"} icon={IconComponent} />
      <div>
        <p style={divStyleText}>{accountName}</p>
        <p style={{ ...divStyleText, ...pStyleName }}>Posteado por {name}</p>
      </div>
    </div>
  );
};

export default WDetailsImage;

WDetailsImage.defaultProps = {
  accountName: "Jane Doe",
  name: "janedoe123",
  icon: React.Component,
};