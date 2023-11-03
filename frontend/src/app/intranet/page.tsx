"use client"
import React from 'react'
import { logout } from '@/redux/ducks/user'
import { useAppDispatch } from '@/redux/hooks'
export default function HomePage() {
    const dispatch = useAppDispatch()
    const handleLogout = () => {
        localStorage.clear()
        window.location.href = '/'
    }
    return (<div>
        <h1>Home Page intranet</h1>
        <button onClick={() => handleLogout()}>cerrar session</button>
    </div>)
}
