package com.social.presentation.profile

import android.app.Activity
import android.content.Intent
import android.graphics.BitmapFactory
import android.net.Uri
import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.view.View
import android.widget.AdapterView
import android.widget.ArrayAdapter
import android.widget.Toast
import androidx.activity.result.contract.ActivityResultContracts
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
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

        binding.iconLeft.setOnClickListener {
            findNavController().navigate(R.id.action_editProfileFragment_to_userProfileFragment)
        }

        binding.buttonAddLink.setOnClickListener {
            binding.socialLinks.visibility =
                if (binding.socialLinks.visibility == View.VISIBLE) {
                    View.GONE
                } else {
                    View.VISIBLE
                }
        }

        userVerification()
        setupAutoCompleteTextView()
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

    private fun setupAutoCompleteTextView() {
        val itemsSocial = listOf("Facebook", "Twitter", "Instagram", "Snapchat", "LinkedIn")
        val adapter = ArrayAdapter(requireContext(), R.layout.items, itemsSocial)
        binding.dropDownSocialNetwork.setAdapter(adapter)
        binding.dropDownSocialNetwork.onItemClickListener =
            AdapterView.OnItemClickListener { _, _, position, _ ->
                val selectedItem = adapter.getItem(position)
                Toast.makeText(requireContext(), "$selectedItem", Toast.LENGTH_SHORT)
                    .show()
            }
    }
}
