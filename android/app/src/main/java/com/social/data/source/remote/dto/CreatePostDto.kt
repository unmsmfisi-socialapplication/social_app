package com.social.data.source.remote.dto

import com.social.domain.model.CreatePost

data class CreatePostDto(
    val response: String,
    val status: String
)

fun CreatePostDto.aCreatePost(): CreatePost {
    return CreatePost(response, status)
}