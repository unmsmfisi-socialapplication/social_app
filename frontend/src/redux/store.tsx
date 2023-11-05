'use client'
import { configureStore, ThunkAction, Action, combineReducers } from '@reduxjs/toolkit'
import userReducer from './ducks/user'
import { persistStore, persistReducer } from 'redux-persist'
import storage from 'redux-persist/lib/storage'

const persistConfig = {
    key: 'root',
    storage,
    whitelist: ['auth'],
}

const rootReducer = combineReducers({
    auth: userReducer,
})
const persistedReducer = persistReducer(persistConfig, rootReducer)

export const store = configureStore({
    reducer: persistedReducer,
    devTools: true,
})
export const persistor = persistStore(store)

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch
export type Thunk = ThunkAction<ReturnType<any>, RootState, unknown, Action<void>>
