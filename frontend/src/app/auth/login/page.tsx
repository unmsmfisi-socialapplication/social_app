"use client";
import EnrollmentHoc from "@/app/auth/auth";
import { WInput, WButton, WLink, WCardAuth } from "@/components";

export default function LoginPage() {
  return (
    <EnrollmentHoc>
      <WCardAuth title="Bienvenido de nuevo" size="large">
        <span>Correo</span>
        <WInput
          placeholder="Ingrese su correo"
          fullWidth
          size="small"
        />
        <span>Contraseña</span>
        <WInput
          type="password"
          placeholder="Ingrese su contraseña"
          fullWidth
          size="small"
        />
        <WLink text="Has olvidado tu contraseña?" underline="none" />
        <span>¿No tienes una cuenta? </span>
        <WLink text="Registrarse" underline="none" displayType="inline-flex" />
        <WButton text="Iniciar Sesión" size="large" />
      </WCardAuth>
    </EnrollmentHoc>
  );
}
