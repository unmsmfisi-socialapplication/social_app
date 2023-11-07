import { IUser } from '@/data/entities/User'
import { createSlice } from '@reduxjs/toolkit'
import { getUser } from '../actions/userAction'
import build from 'next/dist/build'
import { apiSattus } from '@/utilities/Constant'

// TODO:Add datafake

interface UserState {
    user: any
    loading: boolean
    status: string
    error: string | null
}

const initialUser = {} as IUser
const initialState: UserState = {
    user: initialUser,
    loading: false,
    status: apiSattus.IDLE,
    error: null,
}

export const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {
        logout: (state) => {
            state.user = initialUser
            state.loading = false
            state.status = apiSattus.IDLE
            //TODO: clear localstorage
        },
    },
    extraReducers: (builder) => {
        builder.addCase(getUser.fulfilled, (state, action) => {
            state.loading = false
            state.status = apiSattus.SUCCES
            if (action.payload === null) {
                state.error = 'credenciales incorrectas'
            } else {
                state.user = action.payload
                localStorage.setItem('user', JSON.stringify(action.payload))
            }
        })
        builder.addCase(getUser.pending, (state) => {
            state.loading = true
            state.status = apiSattus.LOADING
        })
        builder.addCase(getUser.rejected, (state, action) => {
            state.loading = false
            state.status = apiSattus.FAILED
            state.error = 'credenciales incorrectas'
        })
    },
})

export const selectUser = (state: any) => state.auth.user
export const { logout } = authSlice.actions

export default authSlice.reducer
