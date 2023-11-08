package com.social.presentation.profile

import android.app.Activity
import android.content.Intent
import android.graphics.BitmapFactory
import android.net.Uri
import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.view.View
import androidx.activity.result.contract.ActivityResultContracts
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentEditProfileBinding

class EditProfileFragment : Fragment(R.layout.fragment_edit_profile) {
    private lateinit var binding: FragmentEditProfileBinding
    private val imagePicker =
        registerForActivityResult(
            ActivityResultContracts.StartActivityForResult(),
        ) { result ->
            if (result.resultCode == Activity.RESULT_OK && result.data != null) {
                result.data?.data?.let {
                    uploadImage(it)
                }
            }
        }

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentEditProfileBinding.bind(view)
        action()
    }

    private fun action() {
        binding.buttonUploadProfile.setOnClickListener {
            openGallery()
        }
        userVerification()
    }

    private fun openGallery() {
        val intent =
            Intent(Intent.ACTION_GET_CONTENT)
        intent.type = "image/*"
        imagePicker.launch(intent)
    }

    private fun uploadImage(uri: Uri) {
        val bitmap =
            BitmapFactory.decodeStream(
                requireContext().contentResolver.openInputStream(uri),
            )
        binding.imagenProfile.setImageBitmap(bitmap)
    }

    private fun userVerification() {
        binding.inputUserName.addTextChangedListener(
            object : TextWatcher {
                override fun beforeTextChanged(
                    s: CharSequence?,
                    start: Int,
                    count: Int,
                    after: Int,
                ) {
                }

                override fun onTextChanged(
                    s: CharSequence?,
                    start: Int,
                    count: Int,
                    after: Int,
                ) {
                    if (userUnique(s.toString())) {
                        binding.errorUsername.visibility = View.GONE
                    } else {
                        binding.errorUsername.visibility = View.VISIBLE
                    }
                }

                override fun afterTextChanged(s: Editable?) {}
            },
        )
    }

    fun userUnique(username: String): Boolean {
        if (username == "c" || username == "cr" || username == "crh") {
            return false
        }
        return true
    }
}
