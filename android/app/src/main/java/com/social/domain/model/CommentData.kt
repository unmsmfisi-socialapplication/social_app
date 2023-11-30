package com.social.domain.model

data class CommentData(
    val username: String,
    val time: String,
    val reactionText: String,
    val replyText: String,
    val reactionCount: Int,
)
