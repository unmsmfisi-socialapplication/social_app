'use client'
import React from 'react'
import { useEffect, useState } from 'react'
import RootLayout from '../layout'
import { Box, Grid, Button } from '@mui/material'
import { WButton, WTag } from '@/components'
import './index.scss'
import SideBar from '@/components/organisms/SideBar/SideBar'
import RightBar from '@/components/organisms/RightBar/RightBar'
import SearchIcon from '@mui/icons-material/Search'
import NotificationsNoneIcon from '@mui/icons-material/NotificationsNone'
import MailOutlineIcon from '@mui/icons-material/MailOutline'
import CottageOutlinedIcon from '@mui/icons-material/CottageOutlined'
import FormatListBulletedIcon from '@mui/icons-material/FormatListBulleted'
import BookmarkBorderIcon from '@mui/icons-material/BookmarkBorder'
import PeopleOutlineIcon from '@mui/icons-material/PeopleOutline'
import PersonOutlineIcon from '@mui/icons-material/PersonOutline'
interface PropsIntranet {
    children: React.ReactNode
    rightBar?: boolean
    sideBar?: boolean
}

export default function IntranetHoc({ children, sideBar, rightBar }: PropsIntranet) {
    const [isClient, setIsClient] = useState(false)

    useEffect(() => {
        setIsClient(true)
    }, [])
    return (
        <RootLayout>
            {isClient && (
                <Grid container spacing={{ xs: 2, md: 3 }}>
                    <Grid item xs={3} sm={3} md={3}>
                        {sideBar && <SideBar />}
                        <Box className="listOption">
                            <WTag text="Home" icon={CottageOutlinedIcon} isActive />
                            <WTag text="Explorer" icon={SearchIcon} />
                            <WTag text="Notifications" icon={NotificationsNoneIcon} />
                            <WTag text="Messages" icon={MailOutlineIcon} />
                            <WTag text="Lists" icon={FormatListBulletedIcon} />
                            <WTag text="Bookmarks" icon={BookmarkBorderIcon} />
                            <WTag text="Communities" icon={PeopleOutlineIcon} />
                            <WTag text="Profile" icon={PersonOutlineIcon} />
                        </Box>
                        <WButton variant="contained" text="Post" size="large" />
                    </Grid>
                    <Grid item xs={rightBar ? 6 : 8} sm={rightBar ? 6 : 8} md={rightBar ? 6 : 8}>
                        <Box className="home_section">{children}</Box>
                    </Grid>
                    {rightBar && (
                        <Grid item xs={3} sm={3} md={3}>
                            <RightBar />
                        </Grid>
                    )}
                </Grid>
            )}
        </RootLayout>
    )
}

IntranetHoc.defaultProps = {
    rihtBar: false,
    sideBar: false,
}
