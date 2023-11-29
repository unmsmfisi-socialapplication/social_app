'use client'
import React from 'react'
import { Box, Grid, Button } from '@mui/material'
import WCircleImage from '../../../components/atoms/CircleImage/circleImage'
import { WPostSocial } from '@/components'
import IntranetHoc from '../intranet'

import './index.scss'

export default function ProfilePage() {
    return (
        <IntranetHoc sideBar>
            <img
                style={{
                    maxHeight: '250px',
                    height: 'auto',
                    width: '100%',
                    objectFit: 'cover',
                }}
                src="https://images.unsplash.com/photo-1519751138087-5bf79df62d5b?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
                alt=""
            />
            <Grid container spacing={2} paddingInline={4}>
                <Grid item xs={2.5} sm={2.5} md={2.5}>
                    <Box
                        sx={{ width: 150 }}
                        style={{
                            position: 'relative',
                            top: '-90px',
                        }}
                    >
                        <WCircleImage
                            avatarDefaultURL="https://previews.123rf.com/images/grase/grase1410/grase141000110/32759482-lindo-gato-somal%C3%AD-miente-dentro-del-rect%C3%A1ngulo-de-pl%C3%A1stico.jpg"
                            alt="MCMXCVIII's avatar"
                            size={150}
                        />
                        <h3 className="name_profile">MICHITO 21</h3>
                        <div className="userName">@tumimicho</div>
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
