'use client'
import EnrollmentHoc from '@/app/auth/auth'
import * as Yup from 'yup'
import { WInput, WButton, WCardAuth, WLink } from '@/components'
import { Box } from '@mui/material'
import VisibilityOffOutlinedIcon from '@mui/icons-material/VisibilityOffOutlined'
import { useState } from 'react'
import { INITIAL_FORMIK_VALUES, REGISTER_VALUES, YUP_SCHEMA } from './constant'
import { useFormik } from 'formik'
import { validateEmail, validateName, validatePassword } from '@/utilities/Validation'
import AuthServices from '@/domain/usecases/AuthServises'

export default function RegisterPage() {
    const [register, setRegister] = useState<any>(null)

    const registerRequestLogin = async (resquest: any) => {
        const { data, error } = await AuthServices.registerRequest(resquest)
        if (data && error == null) {
            setRegister({ ...register })
        } else {
            console.log(error)
        }
    }

    const formik = useFormik({
        initialValues: { ...INITIAL_FORMIK_VALUES },
        validationSchema: Yup.object({
            ...YUP_SCHEMA,
        }),
        onSubmit: (values) => {
            // TODO: Add login logic
            registerRequestLogin(values)
        },
    })

    return (
        <EnrollmentHoc>
            <form onSubmit={formik.handleSubmit}>
                <WCardAuth title="Registro" variant="outlined">
                    <span>Nombre Completo</span>
                    <WInput
                        placeholder="Nombre Completo"
                        size="small"
                        fullWidth
                        type="text"
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        name={REGISTER_VALUES.NAME}
                        value={formik.values.name}
                        error={formik.touched.name && !validateName(formik.values.name)}
                        errorMessage={formik.errors.name}
                    />
                    <span>Correo</span>
                    <WInput
                        placeholder="Correo"
                        size="small"
                        fullWidth
                        type="text"
                        name={REGISTER_VALUES.EMAIL}
                        value={formik.values.email}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.email && !validateEmail(formik.values.email)}
                        errorMessage={formik.errors.email}
                    />
                    <span>Nombre de Usuario</span>
                    <WInput
                        placeholder="Nombre de Usuario"
                        size="small"
                        fullWidth
                        type="text"
                        name={REGISTER_VALUES.USERNAME}
                        value={formik.values.username}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.username && !validateName(formik.values.username)}
                        errorMessage={formik.errors.username}
                    />
                    <span>Contraseña</span>
                    <WInput
                        icon={<VisibilityOffOutlinedIcon />}
                        placeholder="Contraseña"
                        size="small"
                        fullWidth
                        type="password"
                        name={REGISTER_VALUES.PASSWORD}
                        value={formik.values.password}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.password && !validatePassword(formik.values.password)}
                        errorMessage={formik.errors.password}
                    />
                    <span>Confirmar Contraseña</span>
                    <WInput
                        size="small"
                        type="password"
                        fullWidth
                        name={REGISTER_VALUES.PASSWORD_CONFIRM}
                        value={formik.values.password_confirm}
                        onChange={formik.handleChange}
                        icon={<VisibilityOffOutlinedIcon />}
                        placeholder="Confirmar Contraseña"
                        onBlur={formik.handleBlur}
                        error={formik.touched.password_confirm && !validatePassword(formik.values.password_confirm)}
                        errorMessage={formik.errors.password_confirm}
                    />
                    <Box>
                        <span style={{ marginRight: '10px' }}>¿Ya tienes una cuenta? </span>
                        <WLink text="¡Entra aquí!" underline="none" displayType="inline-flex" href="/auth/login" />
                    </Box>
                    <WButton type="submit" typeColor="primary" text="Registrarse" size="large" />
                    {register && <span>{register?.reponse}</span>}
                </WCardAuth>
            </form>
        </EnrollmentHoc>
    )
}
