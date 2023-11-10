'use client'

import { Box } from '@mui/material'
//import { WCardFollow } from '@/components'
import CardTopic from '@/components/organisms/CardTopic/CardTopic'
interface PropsRightBar {}

export default function RightBar({}: PropsRightBar) {
    return (
        <Box>
            <h1>Who to follow</h1>
            <span>Search</span>
            <div>
                <CardTopic />
            </div>
              
        </Box>
    )
}
