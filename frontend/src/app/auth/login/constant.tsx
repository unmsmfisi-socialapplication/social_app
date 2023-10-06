import * as Yup from "yup";
import { emailRegex, passwordRegex } from "@/utilities/Constant";

export const YUP_SCHEMA = {
  username: Yup.string()
    .matches(emailRegex, {
      message: "El correo electrónico no tiene un formato válido",
    })
    .required("El correo electrónico es requerido"),
  password: Yup.string()
    .matches(passwordRegex, {
      message:
        "La contraseña debe contener al menos 8 caracteres, incluyendo al menos una letra minúscula, una letra mayúscula, un número y un carácter especial (@$!%*?&)",
    })
    .required("La contraseña es requerida"),
};
export const LOGIN_VALUES = {
  EMAIL: "username",
  PASSWORD: "password",
};

export const INITIAL_FORMIK_VALUES = {
  username: "",
  password: "",
};