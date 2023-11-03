import { User, createUser } from '@/data/entities/User'
import AuthServices from '@/domain/usecases/AuthServises'
import { createAsyncThunk } from '@reduxjs/toolkit'

const UserFake = (): User => {
    return createUser({
        id: 1,
        type: 'user',
        name: 'Juan Gutierrez',
        username: 'gosble',
        email: 'gosble@social.com',
        phone: '941593329',
        roleId: 1,
        preferredUsername: 'gosble',
        summary: 'I am a developer',
    })
}
export const getUser = createAsyncThunk(
    'auth/getUser',
    async ({ username, password }: { username: string; password: string }) => {
        try {
            const request = {
                username,
                password,
            }
            const { data, error } = await AuthServices.authRequest(request)
            if (data && error === null) {
                return UserFake()
            }
            if (request.username === 'myuser123' && request.password === 'Social@123') {
                return UserFake()
            }
            throw new Error('Credenciales incorrectas')
        } catch (error) {
            console.log('error', error)
            throw error // Lanzar el error para que Redux Toolkit lo maneje
        }
    },
)
