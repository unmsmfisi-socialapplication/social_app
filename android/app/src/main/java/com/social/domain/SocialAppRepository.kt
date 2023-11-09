package com.social.domain

import com.social.data.source.remote.dto.CreatePostDto
import com.social.data.source.remote.dto.LoginDto
import com.social.data.source.remote.dto.RegisterDto
import com.social.data.source.remote.dto.RegisterUserErrorDto
import com.social.domain.model.CreatePostBody
import com.social.domain.model.LoginBody
import com.social.domain.model.RegisterBody

interface SocialAppRepository {
    suspend fun validateUser(loginBody: LoginBody): LoginDto

    suspend fun createPost(createPostBody: CreatePostBody): CreatePostDto

    suspend fun registerNewUser(registerBody: RegisterBody): RegisterDto

    suspend fun registerNewUserError(registerBody: RegisterBody): RegisterUserErrorDto
}
