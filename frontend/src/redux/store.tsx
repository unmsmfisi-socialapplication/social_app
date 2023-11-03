'use client'
import { configureStore, ThunkAction, Action } from '@reduxjs/toolkit'
import userReducer from './ducks/user'

export const store = configureStore({
    reducer: {
        auth: userReducer,
    },
    devTools: true,
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export type Thunk = ThunkAction<ReturnType<any>, RootState, unknown, Action<void>>
