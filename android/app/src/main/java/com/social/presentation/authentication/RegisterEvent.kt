package com.social.presentation.authentication

sealed class RegisterEvent {
    data class EnterEmail(val value: String) : RegisterEvent()

    data class EnterUsername(val value: String) : RegisterEvent()

    data class EnterPassword(val value: String) : RegisterEvent()

    object RegisterUser : RegisterEvent()
}
