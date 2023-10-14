package com.social.domain

import com.social.data.source.remote.dto.CreatePostDto
import com.social.data.source.remote.dto.LoginDto
import com.social.domain.model.CreatePostBody
import com.social.domain.model.LoginBody

interface SocialAppRepository {
    suspend fun validateUser(loginBody: LoginBody): LoginDto

    suspend fun createPost(createPostBody: CreatePostBody): CreatePostDto
}
