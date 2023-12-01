package com.social.homeTest

import android.graphics.Bitmap
import android.view.View
import com.google.android.material.imageview.ShapeableImageView
import com.social.presentation.home.NewPostFragment
import org.junit.Assert.assertEquals
import org.junit.Test
import org.mockito.Mockito.mock
import org.mockito.Mockito.verify

class NewPostFragmentTest {
    @Test
    fun testRemoveImage() {
        val newPostFragment = NewPostFragment()
        newPostFragment.selectedImages.clear()

        newPostFragment.imageViews =
            listOf(
                createMockImageView(),
                createMockImageView(),
                createMockImageView(),
                createMockImageView(),
                createMockImageView(),
                createMockImageView(),
            )

        val bitmap1 = createTestBitmap()
        val bitmap2 = createTestBitmap()
        val bitmap3 = createTestBitmap()
        newPostFragment.selectedImages.addAll(
            listOf(
                bitmap1,
                bitmap2,
                bitmap3,
            ),
        )

        newPostFragment.removeImage(1)

        assertEquals(
            2,
            newPostFragment.selectedImages.size,
        )
        assertEquals(
            bitmap1,
            newPostFragment.selectedImages[0],
        )
        assertEquals(
            bitmap3,
            newPostFragment.selectedImages[1],
        )
    }

    private fun createTestBitmap(): Bitmap {
        return Bitmap.createBitmap(
            1,
            1,
            Bitmap.Config.ARGB_8888,
        )
    }

    private fun createMockImageView(): ShapeableImageView {
        return mock(ShapeableImageView::class.java)
    }

    @Test
    fun testUpdateImageUI() {
        val newPostFragment = NewPostFragment()
        newPostFragment.selectedImages.clear()

        val mockImageView1 = createMockImageView()
        val mockImageView2 = createMockImageView()
        val mockImageView3 = createMockImageView()

        newPostFragment.imageViews =
            listOf(
                mockImageView1,
                mockImageView2,
                mockImageView3,
            )

        val bitmap1 = createTestBitmap()
        val bitmap2 = createTestBitmap()
        val bitmap3 = createTestBitmap()
        newPostFragment.selectedImages.addAll(
            listOf(
                bitmap1,
                bitmap2,
                bitmap3,
            ),
        )

        newPostFragment.updateImageUI()

        verify(mockImageView1).visibility = View.VISIBLE
        verify(mockImageView1).setImageBitmap(bitmap1)

        verify(mockImageView2).visibility = View.VISIBLE
        verify(mockImageView2).setImageBitmap(bitmap2)

        verify(mockImageView3).visibility = View.VISIBLE
        verify(mockImageView3).setImageBitmap(bitmap3)
    }
}
