package com.social.presentation.authentication

import com.social.domain.use_case.ValidateUser

data class AuthenticationUseCase(
    val validateUser: ValidateUser
)