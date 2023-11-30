'use client'
import { WLogo, WButton, WTag, WButtonPost } from '@/components'
import { Box, Button } from '@mui/material'
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
        <div className="sidebar__section">
            <WLogo alt="FrameStudentNET" size={30} />
            <Box className="listOptions">
                <WTag path="/intranet" text="Home" icon={CottageOutlinedIcon} />
                <WTag path="#" text="Explorer" icon={SearchIcon} />
                <WTag path="#" text="Notifications" icon={NotificationsNoneIcon} />
                <WTag path="#" text="Messages" icon={MailOutlineIcon} />
                <WTag path="#" text="Lists" icon={FormatListBulletedIcon} />
                <WTag path="/intranet/analitics" text="Estadisticas" icon={BookmarkBorderIcon} />
                <WTag path="#" text="Communities" icon={PeopleOutlineIcon} />
                <WTag path="/intranet/profile" text="Perfil" icon={PersonOutlineIcon} />
                <WButtonPost text="Post" />
            </Box>
        </div>
    )
}
