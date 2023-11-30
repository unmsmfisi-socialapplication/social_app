'use client'
import React from 'react'
import { useAppDispatch } from '@/redux/hooks'
import IntranetHoc from './intranet'
import { WPostSocial } from '@/components'
import { Button } from '@mui/material'

export default function HomePage() {
    return (
        <IntranetHoc sideBar rightBar>
            <h1>Inicio</h1>
            <Button variant="contained" onClick={() => window.open('socialappHTML.html', '_blank')}>Abrir Dashboard</Button>
            <WPostSocial />
            <WPostSocial />
            <WPostSocial />
        </IntranetHoc>
    )
}
