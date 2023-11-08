package com.social.presentation.authentication

import com.social.domain.usecase.RegisterNewUser
import com.social.domain.usecase.ValidateUser

data class AuthenticationUseCase(
    val validateUser: ValidateUser,
    val registerNewUser: RegisterNewUser,
)
