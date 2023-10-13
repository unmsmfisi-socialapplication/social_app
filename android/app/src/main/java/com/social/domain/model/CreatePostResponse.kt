package com.social.domain.model

data class CreatePostResponse(
    val Id: Int,
    val UserId: Int,
    val Title: String,
    val Description: String,
    val HasMultimedia: Boolean,
    val Public: Boolean,
    val Multimedia: String? = "",
    val InsertionDate: String,
    val UpdateDate: String,
)
