package com.social.interactionTest.chatsTest

import com.social.domain.model.ChatUserData
import org.junit.Assert.assertEquals
import org.junit.Test

class ChatUserDataTest {
    @Test
    fun testChatUserData() {
        val chatUserData =
            ChatUserData(
                name = "John Doe",
                image = "base64encodedimage",
                message = "Hello",
                hourSend = "12:00 PM",
                countNotification = "3",
                token = "token123",
            )

        assertEquals("John Doe", chatUserData.name)
        assertEquals("base64encodedimage", chatUserData.image)
        assertEquals("Hello", chatUserData.message)
        assertEquals("12:00 PM", chatUserData.hourSend)
        assertEquals("3", chatUserData.countNotification)
        assertEquals("token123", chatUserData.token)
    }
}
