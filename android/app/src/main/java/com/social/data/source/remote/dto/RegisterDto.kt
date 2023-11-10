package com.social.data.source.remote.dto

import com.social.domain.model.RegisterData

data class RegisterDto(
    val email: String,
    val user_name: String,
)

fun RegisterDto.aRegister(): RegisterData {
    return RegisterData(email, user_name)
}
