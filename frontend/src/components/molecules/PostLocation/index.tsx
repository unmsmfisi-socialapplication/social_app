'use client'
import React from 'react'
import './index.scss'
import { Box, Button, Typography } from '@mui/material'
import WButton from './../../atoms/Button/button'

const WPostLocation: React.FC = ({}) => {
    return (
        <div className="post_location_main_container">
            <div className="post_location_container">
                <Typography variant="h6" component="div" className="headerText">
                    publicar UBICACIÓN
                </Typography>
                <Button className="closeButton">X</Button>
            </div>
            <iframe
                width="100%"
                height="400"
                style={{ border: 0 }}
                allowFullScreen
                src="https://www.google.com/maps/embed?pb=!1m18!1m12!1m3!1d387273.5004385689!2d-74.25908989999999!3d40.6976701!2m3!1f0!2f0!3f0!3m2!1i1024!2i768!4f13.1!3m3!1m2!1s0x89c24fa5d33f083b%3A0xc80b8f06e177fe62!2sNew%20York%2C%20NY%2C%20USA!5e0!3m2!1sen!2sin!4v1638529735378!5m2!1sen!2sin"
                title="Google Maps"
            ></iframe>
            <div className="post_location_footer">
                <Box className="modal--report--cancel">
                    <WButton text={`Cancelar`} variant="outlined" typeColor="white" />
                </Box>
                <Box className="modal--report--submit">
                    <WButton text={`Publicar`} />
                </Box>
            </div>
        </div>
    )
}

export default WPostLocation
