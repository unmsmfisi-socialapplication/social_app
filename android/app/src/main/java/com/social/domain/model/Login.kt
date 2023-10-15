package com.social.domain.model

data class Login(
    val response: String,
    val status: String,
)

class InvalidUserException(message: String) : Exception(message)
