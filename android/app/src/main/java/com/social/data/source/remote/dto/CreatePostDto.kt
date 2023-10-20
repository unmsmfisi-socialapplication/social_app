package com.social.data.source.remote.dto

import com.social.domain.model.CreatePost
import com.social.domain.model.CreatePostResponse

data class CreatePostDto(
    val response: CreatePostResponse,
    val status: String,
)

fun CreatePostDto.aCreatePost(): CreatePost {
    return CreatePost(response, status)
}
