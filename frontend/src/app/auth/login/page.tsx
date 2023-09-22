'use client'
import EnrollmentHoc from "@/app/auth/auth";
import CardAuth from "@/components/organisms/CardAuth";

import { WInput } from "@/components";
import {WButton} from "@/components";
import BasicsLink from "@/components/atoms/link/link";

export default function LoginPage() {

  return (
    <EnrollmentHoc>
        <CardAuth title="Bienvenido de nuevo" size="large">
            <br />
            <span>Correo</span>
            <WInput typeColor="primary" placeholder="Ingrese su correo" fullWidth size="small"/>
            <br />
            <br />
            <span>Contraseña</span>
            <WInput type="password" typeColor="primary" placeholder="Ingrese su contraseña" fullWidth size="small" />
            <BasicsLink text="Has olvidado tu contraseña?" underline="none" />
            <br /> <br />
            <div >
            <span >¿No tienes una cuenta? </span>
            <BasicsLink  text="Registrarse" underline="none" displayType="inline-flex" />
            </div>
            <WButton text="Iniciar Sesión" size="large" />
        </CardAuth>
    </EnrollmentHoc>
  );
}