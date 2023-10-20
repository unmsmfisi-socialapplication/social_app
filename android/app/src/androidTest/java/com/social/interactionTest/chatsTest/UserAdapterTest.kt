package com.social.interactionTest.chatsTest

import androidx.test.ext.junit.runners.AndroidJUnit4
import com.social.domain.model.ChatUserData
import com.social.presentation.interaction.chats.UserAdapter
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith

@RunWith(AndroidJUnit4::class)
class UserAdapterTest {
    private lateinit var usersAdapter: UserAdapter

    @Before
    fun setUp() {
        val users =
            listOf(
                ChatUserData(
                    name = "John Doe",
                    image = "base64encodedimage",
                    message = "Hello",
                    hourSend = "12:00 PM",
                    countNotification = "3",
                ),
                ChatUserData(
                    name = "Jeanpiere Palacios",
                    image = "base64encodedimage2",
                    message = "Esto es una prueba",
                    hourSend = "12:06 PM",
                    countNotification = "1",
                ),
            )
        usersAdapter = UserAdapter(users)
    }

    @Test
    fun testUserAdapterItemCount() {
        assertEquals(2, usersAdapter.itemCount)
    }

    @Test
    fun testUserAdapterFirstUserName() {
        val firstUser = usersAdapter.users[0]
        assertEquals("John Doe", firstUser.name)
    }

    @Test
    fun testUserAdapterFirstUserMessage() {
        val firstUser = usersAdapter.users[0]
        assertEquals("Hello", firstUser.message)
    }

    @Test
    fun testUserAdapterFirstUserHourSend() {
        val firstUser = usersAdapter.users[0]
        assertEquals("12:00 PM", firstUser.hourSend)
    }

    @Test
    fun testUserAdapterFirstUserCountNotification() {
        val firstUser = usersAdapter.users[0]
        assertEquals("3", firstUser.countNotification)
    }

    @Test
    fun testUserAdapterSecondUserName() {
        val secondUser = usersAdapter.users[1]
        assertEquals("Jeanpiere Palacios", secondUser.name)
    }

    @Test
    fun testUserAdapterSecondUserMessage() {
        val secondUser = usersAdapter.users[1]
        assertEquals("Esto es una prueba", secondUser.message)
    }

    @Test
    fun testUserAdapterSecondUserHourSend() {
        val secondUser = usersAdapter.users[1]
        assertEquals("12:06 PM", secondUser.hourSend)
    }

    @Test
    fun testUserAdapterSecondUserCountNotification() {
        val secondUser = usersAdapter.users[1]
        assertEquals("1", secondUser.countNotification)
    }
}
