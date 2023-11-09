package com.social.interactionTest.chatsTest

import android.graphics.Bitmap
import com.social.domain.model.ChatMessageUser
import com.social.presentation.interaction.chats.ChatAdapter
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Test
import org.mockito.Mockito.mock

class ChatAdapterTest {
    private lateinit var chatAdapter: ChatAdapter
    private lateinit var chatMessageList: MutableList<ChatMessageUser>

    @Before
    fun setUp() {
        chatMessageList = mutableListOf<ChatMessageUser>()

        val sendProfileImage: Bitmap = mock(Bitmap::class.java)
        val receiverProfileImage: Bitmap = mock(Bitmap::class.java)
        val senderId = "sender123"

        chatAdapter = ChatAdapter(chatMessageList, receiverProfileImage, sendProfileImage, senderId)
    }

    @Test
    fun testChatAdapterItemCount() {
        assertEquals(chatMessageList.size, chatAdapter.itemCount)
    }

    @Test
    fun testChatAdapterSenderMessageType() {
        // Agregar un mensaje enviado por el remitente al adaptador
        chatMessageList.add(
            ChatMessageUser(senderId = "sender123", receiverId = "receiver456", message = "Hello", dateTime = "2023-10-06 10:00 AM"),
        )

        // Verificar que el tipo de vista para este mensaje sea VIEW_TYPE_SENT
        assertEquals(ChatAdapter.VIEW_TYPE_SENT, chatAdapter.getItemViewType(0))
    }

    @Test
    fun testChatAdapterReceiverMessageType() {
        // Agregar un mensaje recibido por el receptor al adaptador
        chatMessageList.add(
            ChatMessageUser(senderId = "receiver456", receiverId = "sender123", message = "Hi", dateTime = "2023-10-06 10:05 AM"),
        )

        // Verificar que el tipo de vista para este mensaje sea VIEW_TYPE_RECEIVER
        assertEquals(ChatAdapter.VIEW_TYPE_RECEIVER, chatAdapter.getItemViewType(0))
    }

    @Test
    fun testChatAdapterMessageContent() {
        val message = "This is a test message"
        chatMessageList.add(
            ChatMessageUser(senderId = "sender123", receiverId = "receiver456", message = message, dateTime = "2023-10-06 10:10 AM"),
        )

        assertEquals(message, chatAdapter.chatMessage[0].message)
    }
}
