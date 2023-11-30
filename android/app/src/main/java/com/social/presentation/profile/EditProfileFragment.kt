package com.social.presentation.profile

import android.app.Activity
import android.content.Intent
import android.graphics.BitmapFactory
import android.net.Uri
import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
import android.view.LayoutInflater
import android.view.View
import android.widget.ArrayAdapter
import android.widget.Toast
import androidx.activity.result.contract.ActivityResultContracts
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentEditProfileBinding
import com.social.databinding.ItemSocialLinkLayoutBinding
import com.social.utils.FragmentUtils

class EditProfileFragment : Fragment(R.layout.fragment_edit_profile) {
    private lateinit var binding: FragmentEditProfileBinding
    val socialLinkList: MutableList<View> = mutableListOf()
    private val selectedNetworks: MutableSet<String> = mutableSetOf()
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
        addSocialLinkItem("")
    }

    private fun action() {
        binding.buttonUploadProfile.setOnClickListener {
            openGallery()
        }

        binding.iconLeft.setOnClickListener {
            cleanFormEditProfile()
            FragmentUtils.replaceFragment(
                requireActivity().supportFragmentManager,
                UserProfileFragment(),
            )
        }

        binding.buttonAddLink.setOnClickListener {
            toggleLinearLayoutVisibility()
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

    private fun cleanFormEditProfile() {
        binding.linearLayoutContainer.removeAllViews()
        socialLinkList.clear()
        binding.inputName.text?.clear()
        binding.inputUserName.text?.clear()
        binding.inputBiography.text?.clear()
    }

    private fun toggleLinearLayoutVisibility() {
        val visibility =
            if (binding.linearLayoutContainer.visibility == View.VISIBLE) View.GONE else View.VISIBLE
        binding.linearLayoutContainer.visibility = visibility
    }

    private fun setupAddButton(socialLinkItemBinding: ItemSocialLinkLayoutBinding) {
        val addButton = socialLinkItemBinding.buttonAddSocial
        addButton.setOnClickListener {
            val autoCompleteTextView = socialLinkItemBinding.dropDownSocialNetwork
            val inputSocialLink = socialLinkItemBinding.inputSocialLink

            val selectedNetwork = autoCompleteTextView.text.toString().trim()
            val socialLink = inputSocialLink.text.toString().trim()

            if (selectedNetwork.isNotEmpty() && socialLink.isNotEmpty()) {
                val currentIndex = socialLinkList.indexOf(socialLinkItemBinding.root)
                if (currentIndex == socialLinkList.lastIndex && shouldAllowAddMoreItems()) {
                    addSocialLinkItem(selectedNetwork)
                    showMessage("Red social añadida")
                } else {
                    showMessage("Usted ya ingresó todas sus redes sociales")
                }
            }
        }
    }

    private fun shouldAllowAddMoreItems(): Boolean {
        val allSocialNetworks = arrayOf("Facebook", "Instagram", "TikTok", "Twitter", "LinkedIn")
        val remainingNetworks = allSocialNetworks.filterNot { selectedNetworks.contains(it) }
        return remainingNetworks.size > 1
    }

    private fun showMessage(message: String) {
        Toast.makeText(requireContext(), message, Toast.LENGTH_SHORT).show()
    }

    fun addSocialLinkItem(selectedNetwork: String) {
        val inflater = LayoutInflater.from(requireContext())
        val socialLinkItemBinding = ItemSocialLinkLayoutBinding.inflate(inflater, null, false)
        socialLinkList.add(socialLinkItemBinding.root)

        binding.linearLayoutContainer.addView(socialLinkItemBinding.root)
        setupAddButton(socialLinkItemBinding)

        val autoCompleteTextView = socialLinkItemBinding.dropDownSocialNetwork
        val adapter =
            ArrayAdapter(
                requireContext(),
                android.R.layout.simple_dropdown_item_1line,
                getFilteredSocialNetworks(selectedNetwork),
            )
        autoCompleteTextView.setAdapter(adapter)
    }

    private fun getFilteredSocialNetworks(selectedNetwork: String): Array<String> {
        selectedNetworks.add(selectedNetwork)

        val allSocialNetworks = arrayOf("Facebook", "Instagram", "TikTok", "Twitter", "LinkedIn")
        val filteredNetworks = allSocialNetworks.filterNot { selectedNetworks.contains(it) }
        return filteredNetworks.toTypedArray()
    }
}
