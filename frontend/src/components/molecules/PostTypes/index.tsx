import React from 'react'
import AddPhotoAlternateIcon from '@mui/icons-material/AddPhotoAlternate'

interface WPostTypesProps {
    iconComponent: React.ReactNode
    typeName: string
}

const WPostTypes: React.FC<WPostTypesProps> = ({ iconComponent, typeName }) => {
    const divStyle = {
        display: 'flex',
        width: '167px',
        height: '38px',
        padding: '8px 12px',
        alignItems: 'center',
        gap: '8px',
        alignSelf: 'stretch',
        borderRadius: '20px',
        background: '#CAE9FF',
        boxShadow: '0px 4px 4px 0px rgba(0, 0, 0, 0.25)',
        cursor: 'pointer',
    }
    const pStyleIcon = {
        width: '38px',
        height: '38px',
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
    }
    const divStyleText = {
        maxWidth: '576px',
        color: '#1B4965',
        fontFamily: 'Poppins',
        fontSize: '12px',
        fontStyle: 'normal',
        fontWeight: '400',
        lineHeight: '22.78px',
        flex: 1,
    }
    const styleType = {
        fontWeight: 'bold',
    }

    return (
        <div style={divStyle} data-testid="w-post-types">
            <div style={pStyleIcon}>{iconComponent}</div>
            <div style={divStyleText}>
                Publicar <span style={styleType}>{typeName}</span>
            </div>
        </div>
    )
}

export default WPostTypes

WPostTypes.defaultProps = {
    iconComponent: <AddPhotoAlternateIcon />,
    typeName: 'FOTO',
}
