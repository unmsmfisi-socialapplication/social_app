import React from 'react'

const WLogo = () => {
    const containerStyle = {
        display: 'flex',
        alignItems: 'center',
    }

    const imageStyle = {
        marginRight: '15px',
        width: '50px',
        height: '50px',
    }

    return (
        <div style={containerStyle}>
            <img src="/images/FrameStudentNET.png" alt="FrameStudentNET" style={imageStyle} />
            <h1>Student Network</h1>
        </div>
    )
}

export default WLogo
