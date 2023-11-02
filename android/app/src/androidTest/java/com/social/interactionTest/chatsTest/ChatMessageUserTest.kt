package com.social.interactionTest.chatsTest

import com.social.domain.model.ChatMessageUser
import org.junit.Assert.assertEquals
import org.junit.Test

class ChatMessageUserTest {
    @Test
    fun testChatMessageUserProperties() {
        val chatMessageUser =
            ChatMessageUser(
                senderId = "sender123",
                receiverId = "receiver456",
                message = "Hello, World!",
                dateTime = "10:00",
            )

        assertEquals("sender123", chatMessageUser.senderId)
        assertEquals("receiver456", chatMessageUser.receiverId)
        assertEquals("Hello, World!", chatMessageUser.message)
        assertEquals("10:00", chatMessageUser.dateTime)
    }
}
