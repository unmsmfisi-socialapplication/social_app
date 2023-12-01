'use client'
import { useState } from 'react'
import { Button } from '@mui/material'
import {
    WButton,
    WInput,
    WCircleIcon,
    WModalPhoto,
    WComment,
    WButtonMotion,
    WCardFollow,
    WSearch,
    WTopicFollow,
    WUserCHATCTA,
    WPublicationConfirm,
    WReportPublication,
} from '@/components'
import AccountCircleIcon from '@mui/icons-material/AccountCircle'
import CheckIcon from '@mui/icons-material/Check'
import AllInclusive from '@mui/icons-material/AllInclusive'
import WTag from '@/components/atoms/Tag/tag'
import SearchIcon from '@mui/icons-material/Search'
import NotificationsNoneIcon from '@mui/icons-material/NotificationsNone'
import MailOutlineIcon from '@mui/icons-material/MailOutline'
import CottageOutlinedIcon from '@mui/icons-material/CottageOutlined'
import FormatListBulletedIcon from '@mui/icons-material/FormatListBulleted'
import BookmarkBorderIcon from '@mui/icons-material/BookmarkBorder'
import PeopleOutlineIcon from '@mui/icons-material/PeopleOutline'
import PersonOutlineIcon from '@mui/icons-material/PersonOutline'
import AvatarInput from '@/components/molecules/AvatarInput'
import WDetailsImage from '@/components/molecules/DetailsImage/index'
import CommentThink from '@/components/molecules/CommentThink'
import WPostTypes from '@/components/molecules/PostTypes/index'
import AddPhotoAlternateIcon from '@mui/icons-material/AddPhotoAlternate'
import SubscriptionsIcon from '@mui/icons-material/Subscriptions'
import GifBoxIcon from '@mui/icons-material/GifBox'
import LocationOnIcon from '@mui/icons-material/LocationOn'
import RootLayout from '../layout'
import WSelectedText from '@/components/atoms/SelectText/selectText'
import ExpandCircleDownIcon from '@mui/icons-material/ExpandCircleDown'

export default function TestPage() {
    const [count, setCount] = useState(0)
    const [password, setPassword] = useState('')
    const socialNetworkOptions = [
        { value: 'Twitter' },
        { value: 'Facebook' },
        { value: 'Instagram' },
        { value: 'Tik Tok' },
    ]
    const [modalConfirm, setModalConfirm] = useState<boolean>(false)
    const [reportModal, setReportModal] = useState<boolean>(false)

    const handleCount = () => {
        setCount(count + 1)
        alert(count)
    }
    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setPassword(e.target.value)
    }
    return (
        <RootLayout>
            <Button variant="contained">Hello World</Button>
            <WButtonMotion />
            <div
                style={{
                    paddingLeft: '15px',
                    width: '1000px',
                    height: '100px',
                    backgroundColor: 'red',
                }}
            >
                <WButton typeColor="secondary" text="DDdasdasdasdasdasdasdasdasd" />
                <WButton typeColor="terciary" text="button" size="large" />
            </div>
            <WButton text="test" size="large" typeColor="terciary" />
            <WComment />
            <h1>Test Page</h1>
            <button onClick={handleCount}>presioname</button>
            <div>
                <WInput placeholder="Nombre" error={true} errorMessage="error" />

                <WInput
                    typeColor="primary"
                    icon={<AccountCircleIcon />} // Icono de usuario
                    placeholder="Nombre de usuario"
                    fullWidth
                />

                <WInput typeColor="secondary" icon={<AccountCircleIcon />} placeholder="Correo electrónico" fullWidth />
            </div>

            <WInput
                value={password}
                name="password"
                placeholder="Contraseña"
                type="password"
                onChange={handleChange}
                fullWidth
            />
            <div
                //Estilos a usar para la caja
                style={{
                    width: '500px',
                    height: '150px',
                    display: 'flex',
                    flexDirection: 'column',
                    justifyContent: 'center',
                    alignItems: 'center',
                    margin: '15px auto',
                    gap: '15px',
                }}
            >
                <WInput
                    typeColor="primary"
                    icon={<AccountCircleIcon />}
                    placeholder="Correo electrónico"
                    size="small"
                    fullWidth
                    type="text"
                />
                <WInput
                    typeColor="primary"
                    icon={undefined}
                    placeholder="Contraseña"
                    size="small"
                    fullWidth
                    type="password"
                />
            </div>

            <div>
                <WSearch />
                {/* Otras partes de tu aplicación */}
            </div>

            <WCircleIcon iconSize={30} icon={CheckIcon} />
            <WCircleIcon iconSize={50} icon={AllInclusive} typeColor="secondary" />
            <WModalPhoto warning />
            <div style={{ display: 'flex', flexDirection: 'column', gap: '10px', margin: '10px' }}>
                <WTag text="Home" icon={CottageOutlinedIcon} isActive />
                <WTag text="Explorer" icon={SearchIcon} />
                <WTag text="Notifications" icon={NotificationsNoneIcon} />
                <WTag text="Messages" icon={MailOutlineIcon} />
                <WTag text="Lists" icon={FormatListBulletedIcon} />
                <WTag text="Bookmarks" icon={BookmarkBorderIcon} />
                <WTag text="Communities" icon={PeopleOutlineIcon} />
                <WTag text="Profile" icon={PersonOutlineIcon} />
                <Button variant="contained">Post</Button>
            </div>
            <AvatarInput />
            <div>
                <WDetailsImage />
            </div>
            <CommentThink
                avatarDefaultURL="https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
                publicTag={'Público'}
                placeholder={'Escribe lo que estás pensando'}
            />
            <div>
                <WCardFollow
                    avatar="https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"
                    name="XokasXD"
                    userhandle="XokasXD"
                />
            </div>
            <WUserCHATCTA />
            <div>
                <WTopicFollow name="Tecnología" topicHandle="Todo sobre tecnología" />
            </div>
            <div style={{ display: 'flex', flexDirection: 'column', rowGap: '8px' }}>
                <WPostTypes iconComponent={<AddPhotoAlternateIcon />} typeName="FOTO" />
                <WPostTypes iconComponent={<SubscriptionsIcon />} typeName="VIDEO" />
                <WPostTypes iconComponent={<GifBoxIcon />} typeName="GIF" />
                <WPostTypes iconComponent={<LocationOnIcon />} typeName="UBICACIÓN" />
            </div>
            <div
                style={{
                    width: '500px',
                    height: '150px',
                    display: 'flex',
                    flexDirection: 'column',
                    justifyContent: 'center',
                    alignItems: 'center',
                    margin: '15px auto',
                    gap: '15px',
                }}
            >
                <WSelectedText
                    label="Red social"
                    options={socialNetworkOptions}
                    iconComponent={ExpandCircleDownIcon}
                ></WSelectedText>
            </div>
            <WButton text="Abrir publicación" onClick={() => setModalConfirm(true)} />
            <WPublicationConfirm
                open={modalConfirm}
                onClose={() => setModalConfirm(false)}
                onConfirm={() => console.log('Publicación subida')}
            />
            <WButton text="Reportar este usuario" onClick={() => setReportModal(true)} />
            <WReportPublication
                open={reportModal}
                onClose={() => setReportModal(false)}
                onConfirm={(reason) => console.log(`Se reporto al usuario por ${reason}`)}
            />
        </RootLayout>
    )
}

const TestIcon: React.FC = () => {
    return (
        <div>
            <img
                style={{ height: '30px', width: '30px', borderRadius: '100%' }}
                src="https://images.unsplash.com/photo-1542309667-2a115d1f54c6?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1936&h=1936&q=80"
            />
        </div>
    )
}
