import * as Yup from 'yup'
import { emailRegex, nameRegex, passwordRegex } from '@/utilities/Constant'

export const YUP_SCHEMA = {
    name: Yup.string()
        .matches(nameRegex, {
            message: 'El nombre completo no tiene un formato válido',
        })
        .required('El nombre completo es requerido'),
    email: Yup.string()
        .matches(emailRegex, {
            message: 'El correo electrónico no tiene un formato válido',
        })
        .required('El correo electrónico es requerido'),

    username: Yup.string()
        .matches(emailRegex, {
            message: 'El correo electrónico no tiene un formato válido',
        })
        .required('El correo electrónico es requerido'),
    password: Yup.string()
        .matches(passwordRegex, {
            message:
                'La contraseña debe contener al menos 8 caracteres, incluyendo al menos una letra minúscula, una letra mayúscula, un número y un carácter especial (@$!%*?&)',
        })
        .required('La contraseña es requerida'),

    password_confirm: Yup.string()
        .matches(passwordRegex, {
            message:
                'La contraseña debe contener al menos 8 caracteres, incluyendo al menos una letra minúscula, una letra mayúscula, un número y un carácter especial (@$!%*?&)',
        })
        .required('La contraseña es requerida'),
}

export const REGISTER_VALUES = {
    NAME: 'name',
    EMAIL: 'email',
    USERNAME: 'username',
    PASSWORD: 'password',
    PASSWORD_CONFIRM: 'password_confirm',
}

export const INITIAL_FORMIK_VALUES = {
    name: '',
    email: '',
    username: '',
    password: '',
    password_confirm: '',
}
