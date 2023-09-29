import * as Yup from "yup";
import { emailRegex, passwordRegex } from "@/utilities/Constant";

export const YUP_SCHEMA = {
  username: Yup.string().matches(emailRegex).required("El correo electrónico es requerido"),
  password: Yup.string().matches(passwordRegex).required(
    "La contraseña debe contener al menos 8 caracteres, incluyendo al menos una letra minúscula, una letra mayúscula, un número y un carácter especial (@$!%*?&)"
  ),
};
export const LOGIN_VALUES = {
  EMAIL: "username",
  PASSWORD: "password",
};

export const INITIAL_FORMIK_VALUES = {
  username: "",
  password: "",
};