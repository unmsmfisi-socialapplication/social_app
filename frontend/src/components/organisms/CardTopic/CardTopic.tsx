import React, { useState } from 'react'
import { Box, Button } from '@mui/material'
import { WCardFollow } from '@/components'

interface PropsRightBar {}

const CardTopic: React.FC<PropsRightBar> = () => {
    const totalCards = 4 //total number of cards
    const initialDisplayCount = 3 //maximum card display

    const [consumedCount, setConsumedCount] = useState(initialDisplayCount)

    const handleConsumeWCardFollow = () => {
        if (consumedCount < totalCards) {
            setConsumedCount(consumedCount + 1)
        }
    }

    //Test card generation
    //comment if you want to create personalized cards
    //defaultCardCount = number of cards we want to generate
    const defaultCardCount = 4
    const cardsData = Array.from({ length: defaultCardCount }, () => ({
        ...WCardFollow.defaultProps,
    }))

    //uncomment if you want to create custom cards, comment if you want to create default cards
    /* 
    const cardsData = [
        // Test data
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
    ]*/

    const cardsToRender = cardsData.slice(0, consumedCount).map((card, index) => (
        <div key={index}>
            <WCardFollow {...card} />
        </div>
    ))

    return (
        <Box>
            {cardsToRender}
            {totalCards >= 4 && consumedCount < totalCards && (
                <Button variant="outlined" onClick={handleConsumeWCardFollow}>
                    Ver más
                </Button>
            )}
        </Box>
    )
}

export default CardTopic
