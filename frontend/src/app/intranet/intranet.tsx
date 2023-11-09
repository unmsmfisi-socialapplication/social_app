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
                <Grid container spacing={{ xs: 2, md: 3 }}>
                    <Grid item xs={3} sm={3} md={3}>
                        {sideBar && <SideBar />}
                    </Grid>
                    <Grid item xs={rightBar ? 6 : 8} sm={rightBar ? 6 : 8} md={rightBar ? 6 : 8}>
                        <Box className="home_section">{children}</Box>
                    </Grid>
                    {rightBar && (
                        <Grid item xs={3} sm={3} md={3}>
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
