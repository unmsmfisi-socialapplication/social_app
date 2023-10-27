package com.social.presentation.authentication

data class LoginDataState(
    val isLoading: Boolean = false,
    val dataLogin: String? = "",
    val error: String? = "",
)
