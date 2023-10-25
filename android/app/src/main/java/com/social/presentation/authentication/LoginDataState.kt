package com.social.presentation.authentication

import com.social.data.source.remote.dto.LoginDto
import com.social.domain.model.LoginUserData

data class LoginDataState(
    val isLoading: Boolean = false,
    val dataLogin: String?="",//List<LoginDto> = emptyList()
    val error : String?= ""
)