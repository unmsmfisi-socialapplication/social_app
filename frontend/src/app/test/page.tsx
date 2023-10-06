"use client";
import { useState } from "react";
import Layout from "../layout";
import { Button } from "@mui/material";
import { WButton, WInput, WCircleIcon } from "@/components";
import AccountCircleIcon from "@mui/icons-material/AccountCircle";
import CheckIcon from "@mui/icons-material/Check";
import AllInclusive from "@mui/icons-material/AllInclusive";
import WTag from "@/components/atoms/Tag/tag";
import SearchIcon from "@mui/icons-material/Search";
import NotificationsNoneIcon from "@mui/icons-material/NotificationsNone";
import MailOutlineIcon from "@mui/icons-material/MailOutline";
import CottageOutlinedIcon from "@mui/icons-material/CottageOutlined";
import AvatarInput from "@/components/molecules/AvatarInput";
import WDetailsImage from "@/components/molecules/DetailsImage/index";
import PetsIcon from "@mui/icons-material/Pets"; //used in WDetailsImage for testing purposes

export default function TestPage() {
  const [count, setCount] = useState(0);
  const [password, setPassword] = useState("");
  const [avatar, setAvatar] = useState<File | undefined>();

  const handleCount = () => {
    setCount(count + 1);
    alert(count);
  };
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setPassword(e.target.value);
  };
  return (
    <Layout>
      <Button variant="contained">Hello World</Button>
      <div
        style={{
          paddingLeft: "15px",
          width: "1000px",
          height: "100px",
          backgroundColor: "red",
        }}
      >
        <WButton typeColor="primary" text="DD" disabled />
        <WButton typeColor="secondary" text="button" size="large" />
        <WButton text="test" size="large" disabled />
      </div>
      <h1>Test Page</h1>
      <button onClick={handleCount}>presioname</button>
      <div>
        <WInput placeholder="Nombre" error={true} errorMessage="error" />

        <WInput
          typeColor="primary"
          icon={<AccountCircleIcon />} // Icono de usuario
          placeholder="Nombre de usuario"
          fullWidth
        />

        <WInput
          typeColor="secondary"
          icon={<AccountCircleIcon />}
          placeholder="Correo electrónico"
          fullWidth
        />
      </div>

      <WInput
        value={password}
        name="password"
        placeholder="Contraseña"
        type="password"
        onChange={handleChange}
        fullWidth
      />

      <div
        //Estilos a usar para la caja
        style={{
          width: "500px",
          height: "150px",
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          margin: "15px auto",
          gap: "15px",
        }}
      >
        <WInput
          typeColor="primary"
          icon={<AccountCircleIcon />}
          placeholder="Correo electrónico"
          size="small"
          fullWidth
          type="text"
        />
        <WInput
          typeColor="primary"
          icon={undefined}
          placeholder="Contraseña"
          size="small"
          fullWidth
          type="password"
        />
      </div>
      <WCircleIcon iconSize={30} icon={CheckIcon} />
      <WCircleIcon iconSize={50} icon={AllInclusive} typeColor="secondary" />
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          gap: "10px",
          margin: "10px",
        }}
      >
        <WTag text="Home" icon={CottageOutlinedIcon} isActive />
        <WTag text="Explorer" icon={SearchIcon} />
        <WTag text="Notifications" icon={NotificationsNoneIcon} />
        <WTag text="Messages" icon={MailOutlineIcon} />
      </div>
      <AvatarInput
        avatarValue={avatar}
        onChangeAvatar={(avatar) => setAvatar(avatar)}
      />

      <div>
        <WDetailsImage
          accountName={"Influencia Animal"}
          op={"Miguel D."}
          icon={PetsIcon}
        />
      </div>
    </Layout>
  );
}
