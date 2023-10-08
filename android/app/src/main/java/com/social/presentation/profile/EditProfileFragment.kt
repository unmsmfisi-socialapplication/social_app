package com.social.presentation.profile

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentEditProfileBinding


class EditProfileFragment : Fragment(R.layout.fragment_edit_profile) {
   private lateinit var binding: FragmentEditProfileBinding

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentEditProfileBinding.bind(view)

        action()
    }

    private fun action() {

    }
}