'use client'
import { Provider } from 'react-redux'
import { PersistGate } from 'redux-persist/integration/react'
import { persistStore } from 'redux-persist'
import { store } from './store'

interface Props {
    children: React.ReactNode
}

const persistor = persistStore(store)
export function Providers({ children }: Props) {
    return (
        <PersistGate loading={null} persistor={persistor}>
            <Provider store={store}>{children}</Provider>
        </PersistGate>
    )
}
