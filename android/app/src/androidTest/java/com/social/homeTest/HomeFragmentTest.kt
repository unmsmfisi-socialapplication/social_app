package com.social.homeTest

import android.widget.ImageView
import android.widget.TextView
import androidx.fragment.app.testing.FragmentScenario
import androidx.fragment.app.testing.launchFragmentInContainer
import androidx.test.ext.junit.runners.AndroidJUnit4
import com.social.R
import com.social.domain.model.Post
import com.social.presentation.home.HomeFragment
import org.junit.Assert.assertEquals
import org.junit.Before
import org.junit.Test
import org.junit.runner.RunWith
import org.mockito.Mockito.mock

@RunWith(AndroidJUnit4::class)
class HomeFragmentTest {
    private lateinit var fragmentScenario: FragmentScenario<HomeFragment>

    @Before
    fun setUp() {
        // Inicialización del FragmentScenario
        fragmentScenario = launchFragmentInContainer(themeResId = R.style.Theme_Social)
    }

    @Test
    fun testHandleIconLikeClick() {
        val post =
            Post(
                content = "This is a test content",
                hour = "12:00 PM",
                image = "",
                names = "John Doe",
            )
        val initialLikeCount = post.likeCount

        // Llamar al método handleIconLikeClick con el Post de prueba y verificar los cambios
        fragmentScenario.onFragment { fragment ->
            fragment.handleIconLikeClick(post, mockIconLike, mockCountLikes)
        }

        assertEquals(initialLikeCount + 1, post.likeCount)
    }

    companion object {
        // Mocks para los parámetros de handleIconLikeClick
        private val mockIconLike: ImageView = mock()
        private val mockCountLikes: TextView = mock()
    }
}
