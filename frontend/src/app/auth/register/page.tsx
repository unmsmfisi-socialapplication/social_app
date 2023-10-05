"use client";
import EnrollmentHoc from "@/app/auth/auth";

import { WInput, WButton , WCardAuth, WLink } from "@/components";

import { Box } from "@mui/material";

import VisibilityOffOutlinedIcon from "@mui/icons-material/VisibilityOffOutlined";

export default function RegisterPage() {
  return (
    <EnrollmentHoc>
      <WCardAuth title="Registro" variant="outlined">
          <span>Nombre Completo</span>
          <WInput
            placeholder="Nombre Completo"
            size="small"
            fullWidth
            type="text"
          />
          <span>Correo</span>
          <WInput
            placeholder="Correo"
            size="small"
            fullWidth
            type="text"
          />
          <span>Contraseña</span>
          <WInput
            icon={<VisibilityOffOutlinedIcon />}
            placeholder="Contraseña"
            size="small"
            fullWidth
            type="password"
          />
          <span>Confirmar Contraseña</span>
          <WInput
            icon={<VisibilityOffOutlinedIcon />}
            placeholder="Confirmar Contraseña"
            size="small"
            fullWidth
            type="password"
          />
          <Box>
            <span style={{ marginRight: "10px" }}>¿Ya tienes una cuenta? </span>
            <WLink
              text="¡Entra aquí!"
              underline="none"
              displayType="inline-flex"
              href='/auth/login'
            />
          </Box>
        <WButton typeColor="primary" text="Registrarse" size="large" />
      </WCardAuth>
    </EnrollmentHoc>
  );
}
