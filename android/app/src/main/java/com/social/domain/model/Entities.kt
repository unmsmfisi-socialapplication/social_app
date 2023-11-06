package com.social.domain.model

data class Post(
    val names: String,
    val hour: String,
    val content: String,
    val image: String,
)

data class User(
    val username: String,
)
