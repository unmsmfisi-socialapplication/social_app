'use client'
import React from 'react'
import { useAppDispatch } from '@/redux/hooks'
import IntranetHoc from './intranet'
import CommentThink from '@/components/molecules/CommentThink'
export default function HomePage() {
    const handleLogout = () => {
        console.log('handleLogout called')
        localStorage.clear()
    }
    return (
        <IntranetHoc sideBar rightBar>
            <h1>Inicio</h1>
            <CommentThink />
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
