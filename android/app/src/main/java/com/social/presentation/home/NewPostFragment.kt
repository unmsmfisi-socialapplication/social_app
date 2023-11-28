package com.social.presentation.home

import android.app.Activity
import android.content.Intent
import android.graphics.Bitmap
import android.os.Bundle
import android.provider.MediaStore.Images.Media.getBitmap
import android.view.View
import androidx.activity.result.contract.ActivityResultContracts
import androidx.fragment.app.Fragment
import com.google.android.material.imageview.ShapeableImageView
import com.social.R
import com.social.databinding.FragmentNewPostBinding

class NewPostFragment : Fragment(R.layout.fragment_new_post) {
    private lateinit var binding: FragmentNewPostBinding
    private lateinit var imageViews: List<ShapeableImageView>
    private val selectedImages: MutableList<Bitmap> = mutableListOf()

    companion object {
        const val MAX_IMAGES = 6
    }

    private val galleryLauncher =
        registerForActivityResult(
            ActivityResultContracts.StartActivityForResult()
        ) { result ->
            if (result.resultCode == Activity.RESULT_OK) {
                val data = result.data
                if (data?.clipData != null) {
                    val count = minOf(
                        data.clipData!!.itemCount,
                        MAX_IMAGES - selectedImages.size,
                    )
                    for (i in 0 until count) {
                        val imageUri = data.clipData!!.getItemAt(i).uri
                        val bitmap = getBitmapFromUri(imageUri)
                        selectedImages.add(bitmap)
                    }
                } else if (data?.data != null && selectedImages.size < MAX_IMAGES) {
                    val bitmap = data.data?.let { getBitmapFromUri(it) }
                    if (bitmap != null) {
                        selectedImages.add(bitmap)
                    }
                }
                updateImageUI()
            }
        }

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentNewPostBinding.bind(view)

        imageViews = listOf(
            binding.imageOne,
            binding.imageTwo,
            binding.imageThree,
            binding.imageFour,
            binding.imageFive,
            binding.imageSix,
        )

        for (i in imageViews.indices) {
            imageViews[i].setOnClickListener {
                removeImage(i)
            }
        }
        action()
    }

    private fun action() {
        binding.iconImage.setOnClickListener {
            openGalleryForImages()
        }

        binding.iconVideo.setOnClickListener {
            openGalleryForVideos()
        }

        binding.iconLocation.setOnClickListener {
        }
    }

    @Suppress("DEPRECATION")
    private fun getBitmapFromUri(imageUri: android.net.Uri): Bitmap {
        val contentResolver = requireContext().contentResolver
        return getBitmap(contentResolver, imageUri)
    }

    private fun removeImage(index: Int) {
        if (index < selectedImages.size) {
            selectedImages.removeAt(index)
            updateImageUI()
        }
    }

    private fun updateImageUI() {
        for (i in imageViews.indices) {
            if (i < selectedImages.size) {
                imageViews[i].visibility = View.VISIBLE
                imageViews[i].setImageBitmap(selectedImages[i])
            } else {
                imageViews[i].visibility = View.GONE
            }
        }
    }

    private fun openGalleryForImages() {
        val intent = Intent(Intent.ACTION_GET_CONTENT)
        intent.type = "image/*"
        intent.putExtra(Intent.EXTRA_ALLOW_MULTIPLE, true)
        galleryLauncher.launch(Intent.createChooser(intent, "Select Images"))
    }

    private fun openGalleryForVideos() {
    }
}
