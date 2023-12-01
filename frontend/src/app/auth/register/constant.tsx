import * as Yup from 'yup'
import { emailRegex, nameRegex, passwordRegex, phoneRegex, usernameRegex } from '@/utilities/Constant'

export const YUP_SCHEMA = {
    name: Yup.string()
        .matches(nameRegex, {
            message: 'Formato no valido',
        })
        .required('El nombre completo es requerido'),
    email: Yup.string()
        .matches(emailRegex, {
            message: 'Formato no valido',
        })
        .required('El correo electrónico es requerido'),

    username: Yup.string()
        .matches(usernameRegex, {
            message: 'Formato no valido',
        })
        .required('El uusario es requerido'),
    phone: Yup.string()
        .matches(phoneRegex, {
            message: 'Formato no valido',
        })
        .required('El uusario es requerido'),
    password: Yup.string()
        .matches(passwordRegex, {
            message: 'Formato no valido',
        })
        .required('La contraseña es requerida'),

    password_confirm: Yup.string()
        .matches(passwordRegex, {
            message: 'Formato no valido',
        })
        .required('La contraseña es requerida'),
}

export const REGISTER_VALUES = {
    PHONE : 'phone',
    NAME: 'name',
    EMAIL: 'email',
    USERNAME: 'username',
    PASSWORD: 'password',
    PASSWORD_CONFIRM: 'password_confirm',
}

export const INITIAL_FORMIK_VALUES = {
    phone: '',
    name: '',
    email: '',
    username: '',
    password: '',
    password_confirm: '',
}
