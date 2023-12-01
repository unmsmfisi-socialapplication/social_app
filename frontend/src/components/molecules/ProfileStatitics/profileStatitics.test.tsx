import React from 'react'
import { render, screen } from '@testing-library/react'
import '@testing-library/jest-dom'
import WProfileStatitics from '.'

describe('WProfileStatitics', () => {
    it('should display the number of posts, photos, followers, and following passed as props', () => {
        const posts = 10
        const photos = 5
        const followers = 100
        const following = 50
        render(<WProfileStatitics posts={posts} photos={photos} followers={followers} following={following} />)
        const postElement = screen.getByText(posts.toString())
        expect(postElement).toBeInTheDocument()
        const photoElement = screen.getByText(photos.toString())
        expect(photoElement).toBeInTheDocument()
        const followerElement = screen.getByText(followers.toString())
        expect(followerElement).toBeInTheDocument()
        const followingElement = screen.getByText(following.toString())
        expect(followingElement).toBeInTheDocument()
    })

    it('should render the correct number of posts, photos, followers, and following even if decimal values are passed as props', () => {
        const posts = 10.5
        const photos = 5.2
        const followers = 100.8
        const following = 50.3
        render(<WProfileStatitics posts={posts} photos={photos} followers={followers} following={following} />)
        const postElement = screen.getByText(posts.toString())
        expect(postElement).toBeInTheDocument()
        const photoElement = screen.getByText(photos.toString())
        expect(photoElement).toBeInTheDocument()
        const followerElement = screen.getByText(followers.toString())
        expect(followerElement).toBeInTheDocument()
        const followingElement = screen.getByText(following.toString())
        expect(followingElement).toBeInTheDocument()
    })
})
