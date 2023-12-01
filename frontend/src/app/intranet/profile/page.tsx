'use client'
import React, { useEffect } from 'react'
import { Box, Grid, Button } from '@mui/material'
import WCircleImage from '../../../components/atoms/CircleImage/circleImage'
import { WPostSocial } from '@/components'
import IntranetHoc from '../intranet'
import { useAppSelector } from '@/redux/hooks'

import './index.scss'

export default function ProfilePage() {
    const useSelector = useAppSelector((state) => state.auth)
    useEffect(() => {
        console.log('useSelector', useSelector.user)
    }, [useSelector])
    return (
        <IntranetHoc sideBar>
            <img
                style={{
                    maxHeight: '550px',
                    width: '100%',
                }}
                src="data:image/jpeg;base64,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"
                alt=""
            />
            <Grid container spacing={2} paddingInline={4}>
                <Grid item xs={2.5} sm={2.5} md={2.5}>
                    <Box
                        sx={{ width: '150px' }}
                        style={{
                            position: 'relative',
                            top: '-90px',
                        }}
                    >
                        <WCircleImage avatarDefaultURL={useSelector.user.photo} alt="MCMXCVIII's avatar" size={150} />
                        <h3 className="name_profile">{useSelector.user.name}</h3>
                        <div className="userName">@{useSelector.user.username}</div>
                        <span className="description">I am a developer</span>
                    </Box>
                </Grid>
                <Grid item xs={9.5} sm={9.5} md={9.5} container justifyContent="flex-end">
                    <Box sx={{ width: 150 }}>
                        <Button variant="contained">Editar perfil</Button>
                    </Box>
                </Grid>
            </Grid>
            <Box
                style={{
                    position: 'relative',
                    top: '-50px',
                }}
            >
                <WPostSocial />
            </Box>
        </IntranetHoc>
    )
}
