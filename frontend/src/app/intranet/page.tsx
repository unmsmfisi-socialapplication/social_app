'use client'
import React from 'react'
import { useAppDispatch } from '@/redux/hooks'
import IntranetHoc from './intranet'
import { WPostSocial } from '@/components'

export default function HomePage() {
    return (
        <IntranetHoc sideBar rightBar>
            <h1>Inicio</h1>
            <WPostSocial />
            <WPostSocial />
            <WPostSocial />
        </IntranetHoc>
    )
}
