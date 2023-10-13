import * as Yup from "yup";
import { emailRegex, passwordRegex } from "@/utilities/Constant";

export const YUP_SCHEMA = {
  username: Yup.string()
    .matches(emailRegex, {
      message: "Compo incorrecto o faltante",
    })
    .required("Compo incorrecto o faltante"),
  password: Yup.string()
    .matches(passwordRegex, {
      message:
        "Compo incorrecto o faltante",
    })
    .required("Compo incorrecto o faltante"),
};
export const LOGIN_VALUES = {
  USERNAME: "username",
  PASSWORD: "password",
};

export const INITIAL_FORMIK_VALUES = {
  username: "",
  password: "",
};