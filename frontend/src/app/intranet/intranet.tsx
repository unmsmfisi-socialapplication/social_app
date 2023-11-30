'use client'
import React from 'react'
import { useEffect, useState } from 'react'
import RootLayout from '../layout'
import { Box, Grid, Button } from '@mui/material'
import './index.scss'
import SideBar from '@/components/organisms/SideBar/SideBar'
import RightBar from '@/components/organisms/RightBar/RightBar'

interface PropsIntranet {
    children: React.ReactNode
    rightBar?: boolean
    sideBar?: boolean
}

export default function IntranetHoc({ children, sideBar, rightBar }: PropsIntranet) {
    const [isClient, setIsClient] = useState(false)

    useEffect(() => {
        setIsClient(true)
    }, [])
    return (
        <RootLayout>
            {isClient && (
                <Grid container spacing={{ xs: 2, md: 2 }}>
                    <Grid item xs={2.5} sm={2.5} md={2.5}>
                        {sideBar && <SideBar />}
                    </Grid>
                    <Grid item xs={rightBar ? 7 : 9.5} sm={rightBar ? 7 : 9.5} md={rightBar ? 7 : 9.5}>
                        <Box className="home_section">{children}</Box>
                    </Grid>
                    {rightBar && (
                        <Grid item xs={2.5} sm={2.5} md={2.5}>
                            <RightBar />
                        </Grid>
                    )}
                </Grid>
            )}
        </RootLayout>
    )
}

IntranetHoc.defaultProps = {
    rihtBar: false,
    sideBar: false,
}
