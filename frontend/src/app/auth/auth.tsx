'use client'
import { useEffect, useState } from 'react'
import { Box, Grid } from '@mui/material'
import RootEnrollment from '../Enrollment'
import './auth.scss'
export default function EnrollmentHoc({ children }: { children: React.ReactNode }) {
    const [isClient, setIsClient] = useState(false)

    useEffect(() => {
        setIsClient(true)
    }, [])

    return (
        <RootEnrollment>
            {isClient && (
                <Grid className="auth__section" container spacing={2}>
                    <Grid className="auth__section--side" item xs={6}>
                        <span>STUDENT NETWORK</span>
                    </Grid>
                    <Grid className="auth__section--right" item xs={6}>
                        <Box>{children}</Box>
                    </Grid>
                </Grid>
            )}
        </RootEnrollment>
    )
}
