package com.social.data.source.remote.dto
import com.social.domain.model.Login

data class LoginDto(
    val response: String,
    val status: String,
)

fun LoginDto.aLogin(): Login {
    return Login(response, status)
}
