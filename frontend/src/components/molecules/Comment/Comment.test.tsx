import React from 'react'
import { render } from '@testing-library/react'
import '@testing-library/jest-dom'
import WComment from './comment'

describe('Comment', () => {
    it('renders the component molecule Comment', () => {
        const { getByText } = render(
            <WComment
                user={{
                    usuario: 'millys123',
                    name: 'Milly',
                }}
                sendUser="janedoel123"
                comment="¡Felicidades!"
                countHeart={12}
                countComment={5}
                countShareds={6}
                countStatistics={126}
                countStars={23}
            />,
        )

        expect(getByText('Milly')).toBeInTheDocument()
    })
})
