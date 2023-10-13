"use client";
import EnrollmentHoc from "@/app/auth/auth";
import { WInput, WButton, WLink, WCardAuth } from "@/components";
import { Box } from "@mui/material";

export default function LoginPage() {
  return (
    <EnrollmentHoc>
      <WCardAuth title="Bienvenido de nuevo" size="large">
        <span>Correo</span>
        <WInput placeholder="Ingrese su correo" fullWidth size="small" />
        <span>Contraseña</span>
        <WInput
          type="password"
          placeholder="Ingrese su contraseña"
          fullWidth
          size="small"
        />
        <WLink text="Has olvidado tu contraseña?" underline="none" />
        <Box>
          <span style={{ marginRight : "10px"}} >¿No tienes una cuenta? </span>
          <WLink
            text="Registrarse"
            underline="none"
            displayType="inline-flex"
          />
        </Box>
        <WButton text="Iniciar Sesión" size="large" />
      </WCardAuth>
    </EnrollmentHoc>
  );
}
