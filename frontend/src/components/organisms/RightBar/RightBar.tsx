'use client'

import { WSearch } from '@/components'
import { Box, Button } from '@mui/material'

interface PropsRightBar {}

export default function RightBar({}: PropsRightBar) {
    const handleLogout = () => {
        console.log('handleLogout called')
        localStorage.clear()
        window.location.href = '/'
    }
    return (
        <Box>
            <WSearch />
            <Button variant="contained" onClick={handleLogout}>
                Salir
            </Button>
        </Box>
    )
}
