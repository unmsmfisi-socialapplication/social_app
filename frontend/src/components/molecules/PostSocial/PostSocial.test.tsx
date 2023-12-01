import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import WPostSocial from './../../../components/molecules/PostSocial/index'

describe('WPostSocial', () => {
    it('renders the PostSocial molecule component', () => {
        const { getByTestId } = render(
            <WPostSocial
                published="hace 1 hora"
                postText="El diseño es un mayor trabajo para la comunidad por ello es importante tomar diferentes variantes , y ocupar trabajos de diferentes tipos como lo menciona san diego en su libro de diseño emocional , El diseño es un mayor trabajo para la comunidad por ello es importante tomar diferentes variantes , y ocupar trabajos de diferentes tipos como lo menciona san diego en su libro de diseño emocional"
                user={{
                    name: 'Don Norman',
                    userHandler: 'jndler',
                }}
                postImg="https://previews.123rf.com/images/grase/grase1410/grase141000110/32759482-lindo-gato-somal%C3%AD-miente-dentro-del-rect%C3%A1ngulo-de-pl%C3%A1stico.jpg"
                dataTestid='postSocialId'
            />,
        )

        expect(getByTestId('postSocialId')).toBeInTheDocument();
        expect(screen.getByTestId('postSocialId')).toHaveTextContent('Don Norman');
    })
})
