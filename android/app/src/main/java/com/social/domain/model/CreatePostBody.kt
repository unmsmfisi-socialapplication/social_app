package com.social.domain.model

import android.os.Parcelable
import kotlinx.parcelize.Parcelize

@Parcelize
data class CreatePostBody(
    val userId: Int? = 0,
    val title: String,
    val description: String,
    val hasMultimedia: Boolean,
    val public: Boolean,
    val multimedia: String? = "",
) : Parcelable
