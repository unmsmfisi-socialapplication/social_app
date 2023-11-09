'use client'

import { WSearch } from '@/components'
import { Box } from '@mui/material'

interface PropsRightBar {}

export default function RightBar({}: PropsRightBar) {
    return (
        <Box>
            <WSearch />
        </Box>
    )
}
