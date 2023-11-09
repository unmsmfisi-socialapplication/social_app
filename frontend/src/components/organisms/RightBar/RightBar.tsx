import React, { useState } from 'react'
import { Box, Button } from '@mui/material'
import { WCardFollow } from '@/components'

interface PropsRightBar {}

export default function RightBar({}: PropsRightBar) {
    const totalCards = 4 // Total data (cards) cardsData.length
    const initialDisplayCount = 3 // Maximum quantity displayed

    const [consumedCount, setConsumedCount] = useState(initialDisplayCount)

    const handleConsumeWCardFollow = () => {
        if (consumedCount < totalCards) {
            setConsumedCount(consumedCount + 1)
        }
    }

    // Test data, the amount of data shown here must be the same as that equaled in totalCards
    const cardsData = [
        {
            avatar: 'https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1',
            name: 'XokasXD',
            userhandle: 'XokasXD',
        },
        {
            avatar: 'https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1',
            name: 'XokasXD',
            userhandle: 'XokasXD',
        },
        {
            avatar: 'https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1',
            name: 'XokasXD',
            userhandle: 'XokasXD',
        },
        {
            avatar: 'https://images.pexels.com/photos/220453/pexels-photo-220453.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1',
            name: 'XokasXD',
            userhandle: 'XokasXD',
        },
        // Add more cards here
    ]

    const cardsToRender = cardsData.slice(0, consumedCount).map((card, index) => (
        <div key={index}>
            <WCardFollow {...card} />
        </div>
    ))

    return (
        <Box>
            <h1>RightBar texto de prueba</h1>
            <span>Search</span>
            {cardsToRender}
            {totalCards >= 4 && consumedCount < totalCards && (
                <Button variant="outlined" onClick={handleConsumeWCardFollow}>
                    Ver más
                </Button>
            )}
               
        </Box>
    )
}
