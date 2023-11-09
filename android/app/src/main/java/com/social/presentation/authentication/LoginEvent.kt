package com.social.presentation.authentication

sealed class LoginEvent {
    data class EnterUser(val value: String) : LoginEvent()

    data class EnterPassword(val value: String) : LoginEvent()
    object SearchUser : LoginEvent()
}
