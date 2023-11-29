'use client'
import React from 'react'
import IntranetHoc from '../intranet'
import { WComment, WPostSocial } from '@/components/index'
import ArrowBackIcon from '@mui/icons-material/ArrowBack'
import { Box, Grid } from '@mui/material'
import { useHistory } from '@/utilities/Functions'
import Link from 'next/link'
export default function PostPage() {
    return (
        <IntranetHoc sideBar>
            <div style={{ display: 'flex', alignItems: 'center' }}>
                <Link href="/intranet" style={{ textDecoration: 'none' }}>
                    <ArrowBackIcon />
                </Link>
                <h2>Post</h2>
            </div>
            <div style={{ width: '90%' }}>
                <WPostSocial />
                <div style={{ width: '97%', paddingLeft: '30px' }}>
                    <WComment />
                    <WComment />
                    <WComment />
                </div>
            </div>
        </IntranetHoc>
    )
}
