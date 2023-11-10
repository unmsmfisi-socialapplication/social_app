'use client'
import React from 'react'
import { useAppDispatch } from '@/redux/hooks'
import IntranetHoc from './intranet'
import AddPost from '@/components/organisms/AddPost'
export default function HomePage() {
    const handleLogout = () => {
        console.log('handleLogout called')
        localStorage.clear()
    }
    return (
        <IntranetHoc sideBar rightBar>
            <h1>Home Page</h1>
            <AddPost
                imageUrl='./images/Ellipse 91.svg'
            />
            <br />
            <br />
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
