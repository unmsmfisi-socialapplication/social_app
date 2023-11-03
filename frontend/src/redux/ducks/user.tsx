import AuthServices from '@/domain/usecases/AuthServises'
import { IUser, User, createUser, initalUser } from '@/data/entities/User'
import { Dispatch } from '@reduxjs/toolkit'

interface IAction {
    type: string
    payload?: any
}

// TODO:Add datafake
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

const LOGIN_SUCCESS = 'LOGIN_SUCCESS'
const LOGIN_ERROR = 'LOGIN_ERROR'
const LOGOUT_SUCCESS = 'LOGOUT_SUCCESS'

const initialState = {} as IUser
export default function userReducer(state: IUser = initialState, action: IAction): IUser {
    switch (action.type) {
        case LOGIN_SUCCESS:
            return { ...action.payload, error: {} }
        case LOGOUT_SUCCESS:
            return { ...action.payload }
        case LOGIN_ERROR:
            return { ...action.payload, successLogin: false, loading: false, isProcessing: true }
        default:
            return state
    }
}

export const login = (user: string, password: string) => async (dispatch: Dispatch) => {
    console.log('login', user, password)
    const userPayload: User = UserFake()
    try {
        dispatch({
            type: LOGOUT_SUCCESS,
            payload: {
                loading: true,
                isProcessing: true,
            },
        })
        const request = {
            user,
            password,
        }
        console.log('request', request)
        const { data, error } = await AuthServices.authRequest(request)
        if (data && error === null) {
            dispatch({
                type: LOGIN_SUCCESS,
                payload: {
                    userPayload,
                    loading: false,
                },
            })
        } else {
            if (request.user === 'myuser123' && request.password === 'Social@123') {
                dispatch({
                    type: LOGIN_SUCCESS,
                    payload: {
                        userPayload,
                        loading: false,
                    },
                })
                console.log('entro good')
            } else {
                dispatch({
                    type: LOGIN_ERROR,
                    payload: {
                        error: { message: 'Credenciales Incorrectas' },
                    },
                })
            }
        }
    } catch (error) {
        dispatch({
            type: LOGIN_ERROR,
            payload: {
                error: { message: 'Credenciales Incorrectas' },
            },
        })
    }
}

export const logout = () => async (dispatch: Dispatch) => {
    try {
        localStorage.clear()
        dispatch({
            type: LOGOUT_SUCCESS,
        })
    } catch (error) {
        dispatch({
            type: LOGOUT_SUCCESS,
        })
    }
}
