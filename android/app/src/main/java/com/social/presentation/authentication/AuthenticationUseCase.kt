package com.social.presentation.authentication

import com.social.domain.usecase.ValidateUser

data class AuthenticationUseCase(
    val validateUser: ValidateUser,
)
