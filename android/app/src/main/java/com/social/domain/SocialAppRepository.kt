package com.social.domain

import com.social.data.source.local.UserEntity
import com.social.data.source.remote.dto.CreatePostDto
import com.social.data.source.remote.dto.LoginDto
import com.social.data.source.remote.dto.RegisterUserErrorDto
import com.social.domain.model.CreatePostBody
import com.social.domain.model.LoginBody
import com.social.domain.model.RegisterBody
import com.social.domain.model.RegisterData
import java.util.Objects

interface SocialAppRepository {
    suspend fun validateUser(loginBody: LoginBody): LoginDto<Objects>

    suspend fun createPost(createPostBody: CreatePostBody): CreatePostDto

    suspend fun registerNewUser(registerBody: RegisterBody): RegisterData

    suspend fun registerNewUserError(registerBody: RegisterBody): RegisterUserErrorDto

    suspend fun insertUser(userEntity: UserEntity)

    suspend fun getUser(): UserEntity

    suspend fun updateUser(
        id: Int,
        username: String,
        photo: String,
    )

    suspend fun deleteUser()
}
