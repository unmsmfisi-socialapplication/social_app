import React from 'react'
import WCircleIcon from '../../../components/atoms/CircleIcon/circleIcon'
interface WDetailsImageProps {
    accountName: string
    name: string
    icon: React.ComponentType
}

const DefaultIconPropValue: React.FC = () => {
    // Default Icon prop value
    return (
        <div>
            <img
                style={{ height: '30px', width: '30px', borderRadius: '100%' }}
                src="https://images.unsplash.com/photo-1606425270259-998c37268501?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=1758&q=80"
            />
        </div>
    )
}

const WDetailsImage: React.FC<WDetailsImageProps> = ({ accountName, name, icon }) => {
    const divStyle = {
        display: 'flex',
        alignItems: 'flex-start',
        justifyContent: 'flex-start',
        gap: '11px',
    }
    const divStyleText = {
        margin: '0',
        fontFamily: 'Poppins',
        fontWeight: '600',
        fontSize: '14px',
        lineHeight: '19px',
        color: '#000',
    }
    const pStyleName = {
        color: 'rgba(0, 0, 0, 0.40)',
    }

    return (
        <div style={divStyle}>
            <WCircleIcon iconSize={30} typeColor={'secondary'} icon={icon} />
            <div>
                <p style={divStyleText}>{accountName}</p>
                <p style={{ ...divStyleText, ...pStyleName }}>Posteado por {name}</p>
            </div>
        </div>
    )
}

export default WDetailsImage

WDetailsImage.defaultProps = {
    accountName: 'Jane Doe',
    name: 'janedoe123',
    icon: DefaultIconPropValue,
}
