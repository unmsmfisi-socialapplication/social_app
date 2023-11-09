package com.social.data.source.remote

import com.social.data.source.remote.dto.CreatePostDto
import com.social.data.source.remote.dto.LoginDto
import com.social.data.source.remote.dto.RegisterDto
import com.social.data.source.remote.dto.RegisterUserErrorDto
import com.social.domain.model.CreatePostBody
import com.social.domain.model.LoginBody
import com.social.domain.model.RegisterBody
import retrofit2.http.Body
import retrofit2.http.POST

interface ApiInterface {
    // Login
    @POST("/login")
    suspend fun validateUser(
        @Body loginBody: LoginBody,
    ): LoginDto

    // Create Post
    @POST("/post/")
    suspend fun createPost(
        @Body createPostBody: CreatePostBody,
    ): CreatePostDto

    @POST("/register")
    suspend fun registerUser(
        @Body registerBody: RegisterBody,
    ): RegisterDto

    @POST("/register")
    suspend fun registerUserError(
        @Body registerBody: RegisterBody,
    ): RegisterUserErrorDto
}
