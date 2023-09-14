'use client'
import { useState } from "react"
import Layout from "../layout"
import { Button } from "@mui/material";
import { WButton, WInput } from "@/components";
import AccountCircleIcon from '@mui/icons-material/AccountCircle';

export default function TestPage() {
    const [count, setCount] = useState(0);

  const handleCount = () => {
    setCount(count + 1);
    alert(count);
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
        <WButton>DD</WButton>
        <WButton typeColor="secondary">DD</WButton>
        <WButton typeColor="secondary" size="large">
          DD
        </WButton>
      </div>
      <h1>Test Page</h1>
      <button onClick={handleCount}>presioname</button>
      <div>
      <WInput
        typeColor="primary"
        icon={<AccountCircleIcon />} // Icono de usuario
        placeholder="Nombre de usuario"
        size="small"
        variant="filled"
        fullWidth
      />

      <WInput
        typeColor="secondary"
        icon={<AccountCircleIcon />} // Icono de usuario
        placeholder="Correo electrónico"
        size="medium"
        variant="filled"
        fullWidth
      />
    </div>
        
    <div
      //Estilos a usar para la caja 
        style={{
          width: "500px",
          height: "150px",
          backgroundColor: "red", /*Referencia de la ubicación de los objetos de la caja*/
          display: "flex",
          flexDirection: "column",
          justifyContent: "center",
          alignItems: "center",
          margin: "15px auto",
          gap: "15px",
        }}>
        <WInput
          typeColor="primary"
          icon={<AccountCircleIcon />}
          placeholder="Correo electrónico"
          size="small"
          variant="filled"
          fullWidth
          type="text"
        />
        <WInput
          typeColor="primary"
          icon={undefined}
          placeholder="Contraseña"
          size="small"
          variant="filled"
          fullWidth
          type="password" 
        />
    </div>
    </Layout>
  );
}