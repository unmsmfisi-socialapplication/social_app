package com.social.data.repository

import com.social.domain.model.ChatUserData

class ChatRepositoryImp : ChatRepository {
    override suspend fun getChats(): MutableList<ChatUserData> {
        val listChats = mutableListOf<ChatUserData>()

        val response =
            listOf(
                ChatUserData(
                    "Usuario 1",
                    "",
                    "Mensaje de ejemplo 1",
                    "10:00 am",
                    "3",
                ),
                ChatUserData(
                    "Usuario 2",
                    "",
                    "Mensaje de ejemplo 2",
                    "11:30 am",
                    "2",
                ),
                ChatUserData(
                    "Usuario 3",
                    "",
                    "Mensaje de ejemplo 3",
                    "11:50 am",
                    "4",
                ),
            )

        response.forEach { map ->
            val chatUserData =
                ChatUserData(
                    map.name,
                    map.image,
                    map.message,
                    map.hourSend,
                    map.countNotification,
                )
            listChats.add(chatUserData)
        }
        return listChats
    }
}
