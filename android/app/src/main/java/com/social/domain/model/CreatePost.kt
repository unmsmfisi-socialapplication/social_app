package com.social.domain.model

data class CreatePost(
    val response: CreatePostResponse,
    val status: String,
)

class InvalidCreatePostException(message: String) : Exception(message)
