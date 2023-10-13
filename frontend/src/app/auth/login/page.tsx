"use client";
import * as Yup from "yup";
import React, { useState } from "react";
import { useFormik } from "formik";
import { Box } from "@mui/material";
import EnrollmentHoc from "@/app/auth/auth";
import { WInput, WButton, WLink, WCardAuth } from "@/components";
import { INITIAL_FORMIK_VALUES, LOGIN_VALUES, YUP_SCHEMA } from "./constant";
import { validateUsername, validatePassword } from "@/utilities/Validation";
import AuthRepository from "@/domain/repositories/AuthRepository";

export default function LoginPage() {
  const [auth, setAuth] = useState<any>(null);
  const authRequestLogin = async (request: any) => {
    const { data, error } = await AuthRepository.authRequest(request);
    if (data && error === null) {
      setAuth({ ...data });
      console.log("data", data);
    } else {
      console.log("error", error);
    }
  };
  const formik = useFormik({
    initialValues: { ...INITIAL_FORMIK_VALUES },
    validationSchema: Yup.object({
      ...YUP_SCHEMA,
    }),
    onSubmit: (values) => {
      // TODO: Add login logic
      console.log(values);
      authRequestLogin(values);
    },
  });
  return (
    <EnrollmentHoc>
      <form onSubmit={formik.handleSubmit}>
        <WCardAuth title="Bienvenido" size="large">
          <span>Nombre de usuario</span>
          <WInput
            name={LOGIN_VALUES.USERNAME}
            value={formik.values.username}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            placeholder="Ingrese su usuario"
            error={
              formik.touched.username && !validateUsername(formik.values.username)
            }
            errorMessage={formik.errors.username}
            fullWidth
            size="small"
          />
          <span>Contraseña</span>
          <WInput
            type="password"
            name={LOGIN_VALUES.PASSWORD}
            value={formik.values.password}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            placeholder="Ingrese su contraseña"
            error={
              formik.touched.password && !validatePassword(formik.values.password)
            }
            errorMessage={formik.errors.password}
            size="small"
          />
          <WLink text="Has olvidado tu contraseña?" underline="none" />
          <Box>
            <span style={{ marginRight: "10px" }}>¿No tienes una cuenta? </span>
            <WLink
              text="Registrarse"
              underline="none"
              displayType="inline-flex"
              href='/auth/register'
            />
          </Box>
          <WButton type="submit" text="Iniciar Sesión" size="large" />
          {auth && <span>{auth?.response}</span>}
        </WCardAuth>
      </form>
    </EnrollmentHoc>
  );
}
