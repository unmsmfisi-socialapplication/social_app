package com.social.domain.model

data class Register(
    val email: String,
    val user_name: String,
)

class InvalidRegisterException(message: String) : Exception(message)
