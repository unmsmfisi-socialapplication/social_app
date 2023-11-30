'use client'
import * as Yup from 'yup'
import React, { useEffect, useState } from 'react'
import { useFormik } from 'formik'
import { Box } from '@mui/material'
import EnrollmentHoc from '@/app/auth/auth'
import { WInput, WButton, WLink, WCardAuth } from '@/components'
import { INITIAL_FORMIK_VALUES, LOGIN_VALUES, YUP_SCHEMA } from './constant'
import { useAppSelector, useAppDispatch } from '@/redux/hooks'
import { validateUsername, validatePassword } from '@/utilities/Validation'
import { useRouter } from 'next/navigation'
import { getUser } from '@/redux/actions/userAction'
import { apiSattus } from '@/utilities/Constant'
export default function LoginPage() {
    const router = useRouter()

    const [auth, setAuth] = useState<any>(null)
    const dispatch = useAppDispatch()
    const useSelector = useAppSelector((state) => state.auth)

    const handleStatusAuth = (status: string) => {
        switch (status) {
            case apiSattus.SUCCES:
                router.push('/intranet')
                break
            case apiSattus.FAILED:
                setAuth({ response: 'Credenciales incorrectas' })
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
            // TODO: Add login logic
            const { username, password } = values
            dispatch(getUser({ username, password })).then(() => {
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
                <WCardAuth title="Bienvenido" size="large">
                    <span>Nombre de usuario</span>
                    <WInput
                        name={LOGIN_VALUES.USERNAME}
                        value={formik.values.username}
                        onChange={formik.handleChange}
                        onBlur={formik.handleBlur}
                        placeholder="Ingrese su usuario"
                        error={formik.touched.username && !validateUsername(formik.values.username)}
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
                        error={formik.touched.password && !validatePassword(formik.values.password)}
                        errorMessage={formik.errors.password}
                        size="small"
                    />
                    <WLink text="Has olvidado tu contraseña?" underline="none" />
                    <Box>
                        <span style={{ marginRight: '10px' }}>¿No tienes una cuenta? </span>
                        <WLink text="Registrarse" underline="none" displayType="inline-flex" href="/auth/register" />
                    </Box>
                    <WButton
                        onClick={() => handleStatusAuth(useSelector.status)}
                        loading={useSelector.loading}
                        type="submit"
                        text="Iniciar Sesión"
                        size="large"
                    />
                    {auth && <span>{auth?.response}</span>}
                </WCardAuth>
            </form>
        </EnrollmentHoc>
    )
}
