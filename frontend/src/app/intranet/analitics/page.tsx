'use client'
import React, { useEffect } from 'react'
import dynamic from 'next/dynamic'
import { EmbedProps } from 'powerbi-client-react'
import IntranetHoc from '../intranet'
import { Box } from '@mui/material'
import './index.scss'
const UrlConfig =
    'https://app.powerbi.com/view?r=eyJrIjoiMDI4OWExZGUtYjZlYS00YzFmLThkYjMtYzcwMTFjNDQ4ZGQ4IiwidCI6ImRmODY3OWNkLWE4MGUtNDVkOC05OWFjLWM4M2VkN2ZmOTVhMCJ9'
const PowerBIEmbed = dynamic<EmbedProps>(() => import('powerbi-client-react').then((m) => m.PowerBIEmbed), {
    ssr: false,
})

export default function index() {
    useEffect(() => {}, [])

    return (
        <IntranetHoc sideBar>
            <h1>Estadisticas</h1>

            <Box className="container__bi">
                {typeof window !== 'undefined' && (
                    <PowerBIEmbed
                        cssClassName="container__bi--embed"
                        embedConfig={{
                            type: 'dashboard',
                            embedUrl: UrlConfig,
                        }}
                    />
                )}
            </Box>
        </IntranetHoc>
    )
}
