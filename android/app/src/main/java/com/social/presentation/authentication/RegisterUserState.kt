package com.social.presentation.authentication

import com.social.domain.model.RegisterData

data class RegisterUserState(
    val isLoading: Boolean = false,
    val dataRegister: List<RegisterData> = emptyList(),
    val error: String = "",
)
