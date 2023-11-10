import React from 'react'

interface WLogoProps {
    alt?: string
    size?: number
}

const containerStyle = {
    display: 'flex',
    alignItems: 'center',
}

const WLogo: React.FC<WLogoProps> = ({ alt, size }) => {
    return (
        <div style={containerStyle}>
            <img
                src="/images/FrameStudentNET.png"
                alt={alt}
                style={{ marginRight: '15px', width: size, height: size }}
            />
            <h3>Student Network</h3>
        </div>
    )
}

export default WLogo

WLogo.defaultProps = {
    alt: 'FrameStudentNET',
    size: 30,
}
