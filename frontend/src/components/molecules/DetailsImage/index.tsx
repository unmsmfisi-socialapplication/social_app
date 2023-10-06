import React from "react";
import WCircleIcon from "@/components/atoms/CircleIcon/circleIcon";

interface WDetailsImageProps {
  accountName: string;
  op: string;
  icon: React.ComponentType;
}

const WDetailsImage: React.FC<WDetailsImageProps> = ({
  accountName = "Jane Doe",
  op = "janedoe123",
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
  const pStyleOP = {
    color: "rgba(0, 0, 0, 0.40)",
  };

  return (
    <div style={divStyle}>
      <WCircleIcon iconSize={30} typeColor={"secondary"} icon={IconComponent} />
      <div>
        <p style={divStyleText}>{accountName}</p>
        <p style={{ ...divStyleText, ...pStyleOP }}>Posteado por {op}</p>
      </div>
    </div>
  );
};

export default WDetailsImage;
