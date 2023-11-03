'use client'
import * as Yup from 'yup'
import React, { useState } from 'react'
import { useFormik } from 'formik'
import { Box } from '@mui/material'
import EnrollmentHoc from '@/app/auth/auth'
import { WInput, WButton, WLink, WCardAuth } from '@/components'
import { INITIAL_FORMIK_VALUES, LOGIN_VALUES, YUP_SCHEMA } from './constant'
import { useAppSelector, useAppDispatch } from '@/redux/hooks'
import { validateUsername, validatePassword } from '@/utilities/Validation'
import AuthRepository from '@/domain/repositories/AuthRepository'
import { login } from '@/redux/ducks/user'

export default function LoginPage() {
    const [auth, setAuth] = useState<any>(null)
    const dispatch = useAppDispatch()
    const onLogin = ({ username, password }: { username: string; password: string }) => {
        try {
            dispatch(login(username, password))
        } catch (error) {
            console.log('error', error)
        }
    }
    const formik = useFormik({
        initialValues: { ...INITIAL_FORMIK_VALUES },
        validationSchema: Yup.object({
            ...YUP_SCHEMA,
        }),
        onSubmit: (values) => {
            // TODO: Add login logic
            console.log(values)
            const { username, password } = values
            onLogin({ username, password })
        },
    })
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
                    <WButton type="submit" text="Iniciar Sesión" size="large" />
                    {auth && <span>{auth?.response}</span>}
                </WCardAuth>
            </form>
        </EnrollmentHoc>
    )
}
