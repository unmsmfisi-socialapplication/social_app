'use client'
import { WLogo } from '@/components'
import { Box } from '@mui/material'
import './index.scss'
interface SideBarProps {}

export default function SideBar({}: SideBarProps) {
    return (
        <Box className="sidebar__section">
            <WLogo alt="FrameStudentNET" size={30}/>
            <div>lista de items</div>
        </Box>
    )
}
