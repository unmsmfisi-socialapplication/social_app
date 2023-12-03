'use client'
import EnrollmentHoc from '@/app/auth/auth'
import * as Yup from 'yup'
import { WInput, WButton, WCardAuth, WLink } from '@/components'
import { Box } from '@mui/material'
import VisibilityOffOutlinedIcon from '@mui/icons-material/VisibilityOffOutlined'
import { useEffect, useState } from 'react'
import { INITIAL_FORMIK_VALUES, REGISTER_VALUES, YUP_SCHEMA } from './constant'
import { useFormik } from 'formik'
import { validateEmail, validateName, validatePassword, validatePhone, validateUsername } from '@/utilities/Validation'
import AuthServices from '@/domain/usecases/AuthServises'
import { useAppDispatch, useAppSelector } from '@/redux/hooks'
import { useRouter } from 'next/navigation'
import { apiSattus } from '@/utilities/Constant'
import { getUserRegister } from '@/redux/actions/userAction'

export default function RegisterPage() {
    const [register, setRegister] = useState<any>(null)
    const router = useRouter()
    const dispatch = useAppDispatch()
    const useSelector = useAppSelector((state) => state.auth)

    const handleStatusAuth = (status: string) => {
        console.log('status', status)
        switch (status) {
            case apiSattus.SUCCES:
                router.push('/intranet')
                break
            case apiSattus.FAILED:
                setRegister({ response: 'No pudo registrarse' })
                break
            default:
                break
        }
    }
    const formik = useFormik({
        initialValues: { ...INITIAL_FORMIK_VALUES },
        validationSchema: Yup.object({
            ...YUP_SCHEMA,
        }),
        onSubmit: (values) => {
            const { username, password } = values

            // TODO: Add login logic
            dispatch(getUserRegister({ username, password, values })).then(() => {
                handleStatusAuth(useSelector.status)
            })
        },
    })

    useEffect(() => {
        handleStatusAuth(useSelector.status)
    }, [useSelector.status])

    return (
        <EnrollmentHoc>
            <form onSubmit={formik.handleSubmit}>
                <WCardAuth title="Registro" variant="outlined">
                    <span>Nombre Completo</span>
                    <WInput
                        placeholder="Ingrese su Nombre"
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
                        placeholder="Ingrese su Correo"
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
                        placeholder="Ingrese su Usuario"
                        size="small"
                        fullWidth
                        type="text"
                        name={REGISTER_VALUES.USERNAME}
                        value={formik.values.username}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.username && !validateUsername(formik.values.username)}
                        errorMessage={formik.errors.username}
                    />
                    <span>Telefono</span>
                    <WInput
                        placeholder="Telefono"
                        size="small"
                        fullWidth
                        type="text"
                        name={REGISTER_VALUES.PHONE}
                        value={formik.values.phone}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        error={formik.touched.phone && !validatePhone(formik.values.phone)}
                        errorMessage={formik.errors.phone}
                    />
                    <span>Contraseña</span>
                    <WInput
                        icon={<VisibilityOffOutlinedIcon />}
                        placeholder="Ingrese su Contraseña"
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
