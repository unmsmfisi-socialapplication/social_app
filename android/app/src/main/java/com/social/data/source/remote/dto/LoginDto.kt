package com.social.data.source.remote.dto

import com.social.domain.model.LoginUserData

data class LoginDto<T>(
    val response: LoginUserData,
    val status: String,
)

class InvalidUserException(message: String) : Exception(message)
