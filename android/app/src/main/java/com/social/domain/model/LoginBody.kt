package com.social.domain.model

import android.os.Parcelable
import kotlinx.parcelize.Parcelize

@Parcelize
data class LoginBody(
    val username: String,
    val password: String,
) : Parcelable
