package com.social.domain.model

import android.os.Parcelable
import kotlinx.parcelize.Parcelize

@Parcelize
data class RegisterBody(
    val Phone: String,
    val Email: String,
    val User_name: String,
    val Password: String,
) : Parcelable
