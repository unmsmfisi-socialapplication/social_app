package com.social.presentation.home

import android.Manifest
import android.app.Activity
import android.content.Intent
import android.content.pm.PackageManager
import android.graphics.Bitmap
import android.location.Geocoder
import android.location.Location
import android.os.Bundle
import android.provider.MediaStore.Images.Media.getBitmap
import android.view.View
import androidx.activity.result.contract.ActivityResultContracts
import androidx.core.app.ActivityCompat
import androidx.fragment.app.Fragment
import com.google.android.gms.location.FusedLocationProviderClient
import com.google.android.gms.location.LocationServices
import com.google.android.material.imageview.ShapeableImageView
import com.social.R
import com.social.databinding.FragmentNewPostBinding
import com.social.utils.FragmentUtils
import com.social.utils.PermissionUtils
import java.util.Locale

class NewPostFragment : Fragment(R.layout.fragment_new_post) {
    private lateinit var binding: FragmentNewPostBinding
    lateinit var imageViews: List<ShapeableImageView>
    val selectedImages: MutableList<Bitmap> = mutableListOf()
    private lateinit var fusedLocationClient: FusedLocationProviderClient
    private var isLocationVisible = false

    companion object {
        const val MAX_IMAGES = 6
    }

    private val galleryLauncher =
        registerForActivityResult(
            ActivityResultContracts.StartActivityForResult(),
        ) { result ->
            if (result.resultCode == Activity.RESULT_OK) {
                val data = result.data
                if (data?.clipData != null) {
                    val count =
                        minOf(
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
        fusedLocationClient = LocationServices.getFusedLocationProviderClient(requireActivity())

        imageViews =
            listOf(
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
            toggleLocationVisibility()
        }
        binding.iconGif.setOnClickListener {
            openGalleryForGifs()
        }
        binding.iconPadlock.setOnClickListener {
            FragmentUtils.replaceFragment(
                requireActivity().supportFragmentManager,
                HomeFragment(),
            )
        }
    }

    private fun toggleLocationVisibility() {
        if (isLocationVisible) {
            clearLocationData()
        } else {
            if (PermissionUtils.requestLocationPermission(this)) {
                getLastLocation()
            }
        }
    }

    private fun clearLocationData() {
        binding.textLocation.text = ""
        binding.linearLayoutLocation.visibility = View.INVISIBLE
        isLocationVisible = false
    }

    @Suppress("DEPRECATION")
    private fun getLastLocation() {
        if (ActivityCompat.checkSelfPermission(
                requireContext(),
                Manifest.permission.ACCESS_FINE_LOCATION,
            ) != PackageManager.PERMISSION_GRANTED && ActivityCompat.checkSelfPermission(
                requireContext(),
                Manifest.permission.ACCESS_COARSE_LOCATION,
            ) != PackageManager.PERMISSION_GRANTED
        ) {
            return
        }
        fusedLocationClient.lastLocation
            .addOnSuccessListener { location: Location? ->
                location?.let {
                    val geocoder =
                        Geocoder(
                            requireContext(),
                            Locale.getDefault(),
                        )
                    val addresses =
                        geocoder.getFromLocation(
                            it.latitude,
                            it.longitude,
                            1,
                        )
                    if (addresses != null) {
                        if (addresses.isNotEmpty()) {
                            val city = addresses[0]?.locality
                            val country = addresses[0]?.countryName
                            val locationString = "$city - $country"
                            binding.textLocation.text = locationString
                            binding.linearLayoutLocation.visibility = View.VISIBLE
                            isLocationVisible = true
                        }
                    }
                }
            }
    }

    @Suppress("DEPRECATION")
    private fun getBitmapFromUri(imageUri: android.net.Uri): Bitmap {
        val contentResolver = requireContext().contentResolver
        return getBitmap(contentResolver, imageUri)
    }

    fun removeImage(index: Int) {
        if (index < selectedImages.size) {
            selectedImages.removeAt(index)
            updateImageUI()
        }
    }

    fun updateImageUI() {
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
        galleryLauncher.launch(Intent.createChooser(intent, ""))
    }

    private fun openGalleryForVideos() {
    }

    private fun openGalleryForGifs() {
    }
}
