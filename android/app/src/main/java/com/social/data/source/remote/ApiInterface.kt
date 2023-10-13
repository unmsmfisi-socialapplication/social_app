package com.social.data.source.remote

import com.social.data.source.remote.dto.CreatePostDto
import com.social.data.source.remote.dto.LoginDto
import com.social.domain.model.CreatePostBody
import com.social.domain.model.LoginBody
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
}
