'use client'
import EnrollmentHoc from "@/app/auth/auth";
import AccountCircleIcon from '@mui/icons-material/AccountCircle';
import { Button } from "@mui/material";
import {  WInput } from "@/components";

export default function LoginPage() {

  return (
    <EnrollmentHoc>
      <Button variant="contained">Hello World</Button>
      <WInput
        typeColor="primary"
        icon={<AccountCircleIcon />} // Icono de usuario
        placeholder="Nombre de usuario"
        fullWidth
      />

      <WInput
        typeColor="primary"
        icon={<AccountCircleIcon />} 
        placeholder="Correo electrónico"
        fullWidth
      />
    </EnrollmentHoc>
  );
}