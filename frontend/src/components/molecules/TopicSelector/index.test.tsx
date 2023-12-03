import '@testing-library/jest-dom'
import React from 'react'
import { render, fireEvent } from '@testing-library/react'
import WTopicSelector from './index'
import './index.scss'

describe('WTopicSelector', () => {
  // Prueba para verificar si el componente se renderiza correctamente
  test('renders correctly', () => {
    // Renderizar el componente con temas de prueba
    const { getByText } = render(
      <WTopicSelector
        topicClass="Test Class"
        topic1="Topic 1"
        topic2="Topic 2"
        topic3="Topic 3"
        topic4="Topic 4"
      />,
    )

    // Verificar si cada tema se muestra en el componente
    expect(getByText('Test Class')).toBeInTheDocument()
    expect(getByText('Topic 1')).toBeInTheDocument()
    expect(getByText('Topic 2')).toBeInTheDocument()
    expect(getByText('Topic 3')).toBeInTheDocument()
    expect(getByText('Topic 4')).toBeInTheDocument()
  })

  // Prueba para verificar si los clics en los botones se manejan correctamente
  test('handles button clicks', () => {
    // Renderizar el componente con temas de prueba
    const { getByText } = render(
      <WTopicSelector
        topicClass="Test Class"
        topic1="Topic 1"
        topic2="Topic 2"
        topic3="Topic 3"
        topic4="Topic 4"
      />,
    )

    // Obtener el botón del primer tema
    const buttonTopic1 = getByText('Topic 1')

    // Simular un clic en el botón del primer tema
    fireEvent.click(buttonTopic1)

    // Verificar si el botón del primer tema tiene la clase 'typeButton--quaternary'
    expect(buttonTopic1).toHaveClass('typeButton--quaternary')

    // Simular otro clic en el botón del primer tema
    fireEvent.click(buttonTopic1)

    // Verificar si el botón del primer tema ya no tiene la clase 'typeButton--quaternary'
    expect(buttonTopic1).not.toHaveClass('typeButton--quaternary')

    // Obtener el botón del segundo tema
    const buttonTopic2 = getByText('Topic 2')

    // Simular un clic en el botón del segundo tema
    fireEvent.click(buttonTopic2)

    // Verificar si el botón del segundo tema tiene la clase 'typeButton--quaternary'
    expect(buttonTopic2).toHaveClass('typeButton--quaternary')
  })

  // Prueba para verificar el manejo de clics en el tercer tema
  test('handles button clicks for topic 3', () => {
    const { getByText } = render(
      <WTopicSelector
        topicClass="Test Class"
        topic1="Topic 1"
        topic2="Topic 2"
        topic3="Topic 3"
        topic4="Topic 4"
      />,
    )

    const buttonTopic3 = getByText('Topic 3')

    fireEvent.click(buttonTopic3)

    expect(buttonTopic3).toHaveClass('typeButton--quaternary')
  })

  // Prueba para verificar el manejo de clics en el cuarto tema
  test('handles button clicks for topic 4', () => {
    const { getByText } = render(
      <WTopicSelector
        topicClass="Test Class"
        topic1="Topic 1"
        topic2="Topic 2"
        topic3="Topic 3"
        topic4="Topic 4"
      />,
    )

    const buttonTopic4 = getByText('Topic 4')

    fireEvent.click(buttonTopic4)

    expect(buttonTopic4).toHaveClass('typeButton--quaternary')
  })
})
