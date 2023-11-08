import * as Yup from 'yup'
import { usernameRegex, passwordRegex } from '@/utilities/Constant'

export const YUP_SCHEMA = {
    username: Yup.string()
        .matches(usernameRegex, {
            message: 'Campo incorrecto o faltante',
        })
        .required('Campo incorrecto o faltante'),
    password: Yup.string()
        .matches(passwordRegex, {
            message: 'Campo incorrecto o faltante',
        })
        .required('Campo incorrecto o faltante'),
}
export const LOGIN_VALUES = {
    USERNAME: 'username',
    PASSWORD: 'password',
}

export const INITIAL_FORMIK_VALUES = {
    username: '',
    password: '',
}
