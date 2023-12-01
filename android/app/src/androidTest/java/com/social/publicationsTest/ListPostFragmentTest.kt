package com.social.publicationsTest

import android.widget.ImageView
import com.social.data.repository.PostRepository
import com.social.domain.model.Post
import com.social.presentation.publications.ListPostFragment
import com.social.presentation.publications.ListPostViewModel
import junit.framework.TestCase.assertEquals
import junit.framework.TestCase.assertNotNull
import kotlinx.coroutines.test.runBlockingTest
import org.junit.Test
import org.mockito.Mockito.mock
import org.mockito.Mockito.`when`

class ListPostFragmentTest {
    @Test
    fun testObtainData() =
        runBlockingTest {
            // Arrange
            val viewModel = ListPostViewModel()
            val mockedRepository = mock(PostRepository::class.java)
            viewModel.repository = mockedRepository

            // Simula la respuesta del repositorio
            val fakeData = listOf(Post("Test", "12:00", "Fake content", ""))
            `when`(mockedRepository.obtainPost()).thenReturn(fakeData.toMutableList())

            // Act
            viewModel.obtainData()

            // Assert
            assertEquals(fakeData, viewModel.data.value)
        }

    @Test
    fun testAdapterInitialization() {
        // Arrange
        val fragment = ListPostFragment()

        // Act
        val adapter = fragment.adapter

        // Assert
        assertNotNull(adapter)
    }

    @Test
    fun testHandleIconLikeClick() {
        // Arrange
        val fragment = ListPostFragment()
        val iconLike = mock(ImageView::class.java)

        // Act
        fragment.handleIconLikeClick(iconLike)
    }
}
