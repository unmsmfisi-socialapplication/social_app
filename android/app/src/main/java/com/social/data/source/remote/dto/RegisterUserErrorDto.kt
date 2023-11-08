package com.social.data.source.remote.dto

import com.social.domain.model.RegisterDataError

data class RegisterUserErrorDto(
    val response: String,
    val status: String,
)

fun RegisterUserErrorDto.aRegisterError(): RegisterDataError {
    return RegisterDataError(response, status)
}
