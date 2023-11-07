package com.social.presentation.profile

import android.os.Bundle
import android.view.View
import android.widget.TextView
import androidx.core.content.ContextCompat
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import com.social.R
import com.social.databinding.FragmentUserProfileBinding

class UserProfileFragment : Fragment(R.layout.fragment_user_profile) {
    private lateinit var binding: FragmentUserProfileBinding

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentUserProfileBinding.bind(view)
        setupClickListeners()
        updateTextViewsWithKFormat(
            binding.textNumberPost,
            binding.textNumberPhotos,
            binding.textNumberFollowers,
            binding.textNumberFollowings
        )
        action()
    }

    private fun action() {
        binding.buttonEditProfile.setOnClickListener {
            findNavController().navigate(R.id.action_userProfileFragment_to_editProfileFragment)
        }
    }

    private fun setupClickListeners() {
        val textPost = binding.textPost
        val textDestacados = binding.textDestacados
        val textActividad = binding.textActividad

        textPost.setOnClickListener {
            updateTextStyle(textPost)
            resetTextStyle(textDestacados, textActividad)
        }

        textDestacados.setOnClickListener {
            updateTextStyle(textDestacados)
            resetTextStyle(textPost, textActividad)
        }

        textActividad.setOnClickListener {
            updateTextStyle(textActividad)
            resetTextStyle(textPost, textDestacados)
        }
    }

    private fun updateTextStyle(textView: TextView) {
        textView.setTextColor(ContextCompat.getColor(requireContext(), R.color.black))
        textView.paint.isUnderlineText = true
    }

    private fun resetTextStyle(vararg textViews: TextView) {
        for (textView in textViews) {
            textView.setTextColor(ContextCompat.getColor(requireContext(), R.color.color03))
            textView.paint.isUnderlineText = false
        }
    }

    private fun updateTextViewsWithKFormat(vararg textViews: TextView) {
        for (textView in textViews) {
            val textValue = textView.text.toString()
            val numberToConvert = textValue.toIntOrNull() ?: 0
            val convertedNumber = convertNumberToK(numberToConvert)
            textView.text = convertedNumber
        }
    }

    private fun convertNumberToK(number: Int): String {
        return when {
            number >= 1000 && number < 1000000 -> "${number / 1000}k"
            number >= 1000000 -> "${number / 1000000}M"
            else -> number.toString()
        }
    }
}
