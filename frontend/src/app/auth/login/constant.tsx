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
        .required('Campo incorrecto o faltante')
        .min(6, 'La contraseña debe tener al menos 6 caracteres') // Minimum length
        .max(56, 'La contraseña no debe tener más de 56 caracteres') // Maximum length
        .matches(
            /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]+$/,
            'La contraseña debe contener al menos una letra minúscula, una letra mayúscula, un número y un carácter especial'
        ),
};

export const ERROR_MESSAGE =
    '¡Oops!... Tu usuario o contraseña es incorrecta. Si no recuerdas tu contraseña, puedes restablecerla.'

export const LOGIN_VALUES = {
    USERNAME: 'username',
    PASSWORD: 'password',
};

export const INITIAL_FORMIK_VALUES = {
    username: '',
    password: '',
};
