"use client";
import * as Yup from "yup";
import { useFormik } from "formik";
import { Box } from "@mui/material";
import EnrollmentHoc from "@/app/auth/auth";
import { WInput, WButton, WLink, WCardAuth } from "@/components";
import { INITIAL_FORMIK_VALUES, LOGIN_VALUES, YUP_SCHEMA } from "./constant";
import { validateEmail, validatePassword } from "@/utilities/Validation";

export default function LoginPage() {
  const formik = useFormik({
    initialValues: { ...INITIAL_FORMIK_VALUES },
    validationSchema: Yup.object({
      ...YUP_SCHEMA,
    }),
    onSubmit: (values) => {
      // TODO: Add login logic
      console.log(values);
    },
  });

  return (
    <EnrollmentHoc>
      <form onSubmit={formik.handleSubmit}>
        <WCardAuth title="Bienvenido de nuevo" size="large">
          <span>Correo Electrónico</span>
          <WInput
            name={LOGIN_VALUES.EMAIL}
            value={formik.values.username}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            placeholder="Ingrese su correo"
            error={
              formik.touched.username && !validateEmail(formik.values.username)
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
              formik.touched.password &&
              !validatePassword(formik.values.password)
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
            />
          </Box>
          <WButton type="submit" text="Iniciar Sesión" size="large" />
        </WCardAuth>
      </form>
    </EnrollmentHoc>
  );
}
