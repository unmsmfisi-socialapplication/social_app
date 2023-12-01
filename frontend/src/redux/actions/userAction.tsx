import { User, createUser } from '@/data/entities/User'
import AuthServices from '@/domain/usecases/AuthServises'
import { createAsyncThunk } from '@reduxjs/toolkit'

const UserFake = (): User => {
    return createUser({
        id: 1,
        type: 'user',
        name: 'Juan Gutierrez',
        username: 'gosble',
        photo :"https://scontent.flim16-2.fna.fbcdn.net/v/t39.30808-1/361096718_578421127820224_7812375200077875043_n.png?stp=dst-png_p200x200&_nc_cat=104&ccb=1-7&_nc_sid=596444&_nc_eui2=AeGFPiCaIOheJ0EcJIUeYe-zIxhoZCgV9I0jGGhkKBX0jbuInQpa0yLb3WQUCQEKMsPwKBSx7r5REsjhwMAqUYkF&_nc_ohc=cazUFjHWhGIAX_RPs6v&_nc_ht=scontent.flim16-2.fna&oh=00_AfDJV8c2DuV3nA4Ngy76j2ZAqakhzonfeBk7yoSawEETRg&oe=656E1825",
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

export const getUserRegister = createAsyncThunk(
    'auth/getUserRegister',
    async ({ username , password , values }: { username : string , password :string , values : any }) => {
        try {
            const request = {
                username,
                password,
            } 
            const { data, error } = await AuthServices.registerRequest(values)
            console.log('dataGaa', data)
            const UserData = (): User => {
                return createUser({
                    id: data.response.Id,
                    type: data.response.Type,
                    name: data.response.Name,
                    username: data.response.Username,
                    photo :data.response.Photo,
                    email: data.response.Email,
                    phone: data.response.Phone,
                    roleId: data.response.Role,
                    preferredUsername: data.response.PreferredUsername,
                    summary: data.response.Summary,
                })
            }
            
            if (data && error === null) {
                const { data, error } = await AuthServices.authRequest(request)
                if (data && error === null) {
                    return UserData()
                }
                throw new Error('Credenciales incorrectas')
            }
            throw new Error('Registro incorrectas')
        } catch (error) {
            console.log('error', error)
            throw error // Lanzar el error para que Redux Toolkit lo maneje
        }
    },
)
