import React, { useState } from 'react'
import WButton from '@/components/atoms/Button/button'
import './index.scss'

interface WTopicSelectorProps {
    topicClass: string
    topic1: string
    topic2: string
    topic3: string
    topic4: string
}

// Componente principal WTopicSelector
const WTopicSelector: React.FC<WTopicSelectorProps> = ({ topicClass, topic1, topic2, topic3, topic4 }) => {
    // Estado para gestionar el tema seleccionado
    const [selectedTopic, setSelectedTopic] = useState<string | null>(null)

    // Función para manejar clics en los botones de temas
    const handleButtonClick = (topic: string) => {
        // Alternar la selección del tema
        setSelectedTopic((prevSelected) => (prevSelected === topic ? null : topic))
    }

    return (
        <div className="WTopicSelector">
            {/* Mostrar el nombre de la clase del tema */}
            <div className="WTopicSelector__class">{topicClass}</div>

            {/* Contenedor para los botones de temas */}
            <div className="WTopicSelector__topics">
                {/* Botón para el primer tema */}
                <div className="WTopicSelector__button">
                    <WButton
                        text={topic1}
                        size="large"
                        typeColor={selectedTopic === topic1 ? 'quaternary' : 'white'}
                        onClick={() => handleButtonClick(topic1)}
                    />
                </div>

                {/* Botón para el segundo tema */}
                <div className="WTopicSelector__button">
                    <WButton
                        text={topic2}
                        size="large"
                        typeColor={selectedTopic === topic2 ? 'quaternary' : 'white'}
                        onClick={() => handleButtonClick(topic2)}
                    />
                </div>

                {/* Botón para el tercer tema */}
                <div className="WTopicSelector__button">
                    <WButton
                        text={topic3}
                        size="large"
                        typeColor={selectedTopic === topic3 ? 'quaternary' : 'white'}
                        onClick={() => handleButtonClick(topic3)}
                    />
                </div>

                {/* Botón para el cuarto tema */}
                <div className="WTopicSelector__button">
                    <WButton
                        text={topic4}
                        size="large"
                        typeColor={selectedTopic === topic4 ? 'quaternary' : 'white'}
                        onClick={() => handleButtonClick(topic4)}
                    />
                </div>
            </div>
        </div>
    )
}

// Propiedades por defecto del componente
WTopicSelector.defaultProps = {
    topicClass: 'Entretenimiento',
    topic1: 'Series y TV',
    topic2: 'Celebridades',
    topic3: 'Comedia',
    topic4: 'Conciertos',
}

export default WTopicSelector
