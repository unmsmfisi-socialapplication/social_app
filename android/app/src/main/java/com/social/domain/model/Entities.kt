package com.social.domain.model

data class Post(
    val names: String,
    val hour: String,
    val content: String,
    val image: String,
    var isLiked: Boolean = false,
    var likeCount: Int = 0,
)

data class User(
    val username: String,
)
