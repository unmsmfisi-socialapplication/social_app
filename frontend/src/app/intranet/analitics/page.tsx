'use client'
import React from 'react'
import { PowerBIEmbed } from 'powerbi-client-react'
import IntranetHoc from '../intranet'
import { Box } from '@mui/material'
import './index.scss'
const UrlConfig =
    'https://app.powerbi.com/view?r=eyJrIjoiMDI4OWExZGUtYjZlYS00YzFmLThkYjMtYzcwMTFjNDQ4ZGQ4IiwidCI6ImRmODY3OWNkLWE4MGUtNDVkOC05OWFjLWM4M2VkN2ZmOTVhMCJ9'

export default function index() {
    return (
        <IntranetHoc sideBar>
            <h1>Estadisticas</h1>
            <Box className="container__bi">
                <PowerBIEmbed
                    cssClassName="container__bi--embed"
                    embedConfig={{
                        type: 'dashboard',
                        embedUrl: UrlConfig,
                    }}
                />
            </Box>
        </IntranetHoc>
    )
}
