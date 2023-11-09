'use client'
import { WLogo, WButton, WTag } from '@/components'
import { Box } from '@mui/material'
import './index.scss'
import SearchIcon from '@mui/icons-material/Search'
import NotificationsNoneIcon from '@mui/icons-material/NotificationsNone'
import MailOutlineIcon from '@mui/icons-material/MailOutline'
import CottageOutlinedIcon from '@mui/icons-material/CottageOutlined'
import FormatListBulletedIcon from '@mui/icons-material/FormatListBulleted'
import BookmarkBorderIcon from '@mui/icons-material/BookmarkBorder'
import PeopleOutlineIcon from '@mui/icons-material/PeopleOutline'
import PersonOutlineIcon from '@mui/icons-material/PersonOutline'
interface SideBarProps {}

export default function SideBar({}: SideBarProps) {
    return (
        <Box className="sidebar__section">
            <WLogo />
            <div> </div>
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
            <WButton variant="contained" text="Post" />
        </Box>
    )
}
