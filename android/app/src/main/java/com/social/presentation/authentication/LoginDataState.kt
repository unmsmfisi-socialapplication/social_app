package com.social.presentation.authentication

import com.social.domain.model.LoginUserData

data class LoginDataState(
    val isLoading: Boolean = false,
    val dataLogin: String?="",//List<LoginUserData> = emptyList(),
    val error : String?= ""
)