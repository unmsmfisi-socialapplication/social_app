'use client'
import React from 'react'
import { useAppDispatch } from '@/redux/hooks'
import IntranetHoc from './intranet'
export default function HomePage() {
    const handleLogout = () => {
        console.log('handleLogout called')
        localStorage.clear()
    }
    return (
        <IntranetHoc sideBar rightBar>
            <h1>Home Page</h1>
            <button
                onClick={() => {
                    handleLogout()
                    window.location.href = '/'
                }}
            >
                Logout
            </button>
        </IntranetHoc>
    )
}
