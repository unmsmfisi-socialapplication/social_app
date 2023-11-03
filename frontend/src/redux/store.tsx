'use client'
import { createStore, combineReducers, compose, applyMiddleware } from 'redux'
import thunk from 'redux-thunk'
import { configureStore } from '@reduxjs/toolkit'
import userReducer from './ducks/user'

export const store = configureStore({
    reducer: {
        auth: userReducer,
    },
    middleware: [thunk],
    devTools: true,
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
