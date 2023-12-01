package com.social.data.repository

import com.social.domain.model.ChatUserData

interface ChatRepository {
    suspend fun getChats(): MutableList<ChatUserData>
}
