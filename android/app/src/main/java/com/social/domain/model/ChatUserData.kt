package com.social.domain.model

import java.io.Serializable

data class ChatUserData(
    var name: String = "",
    var image: String = "",
    var message: String = "",
    var hourSend: String = "",
    var countNotification: String = "",
    var token: String = "",
) : Serializable
