'use client'
import React, { useEffect, useState } from 'react'
import { useFormik } from 'formik'
import { Box } from '@mui/material'
import { useRouter } from 'next/navigation'
import WButtonMotion from '../../../components/atoms/ButtonAnimated/buttonAnimated'

import './404.css' // Estilo opcional

const Error404 = () => {
    return (
        <div className="error-container">
            <div className="logo-container">
                {/* Logo izquierda superior */}
                <img src="/images/FrameStudentNET.png" alt="Logo" style={{ width: '50px', marginRight: '10px' }} />
                <span className="brand-text">STUDENT NETWORK</span>
            </div>

            {/* CONTENT error */}
            <div className="error-content">
                <h1>Oops...</h1>
                <h2>Parece que te perdiste...</h2>
                <p>
                    Lo sentimos, no pudimos encontrar la pagina que estabas buscando. Ésta página no existe o fue movida
                    a otra.
                    <br></br>
                    <br></br>
                    No te preocupes, siempre puedes volver al inicio.
                </p>

                {/* button home */}
                <WButtonMotion id="use-client-button" type="submit" text="Ir al inicio" />
            </div>

            {/* Image NOT FOUND */}
            <div className="error-image">
                <img src="/images/notFound.png" alt="Error Image" />
            </div>
        </div>
    )
}

export default Error404
