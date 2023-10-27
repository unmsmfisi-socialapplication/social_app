package com.social.presentation.profile

import android.os.Bundle
import android.text.Editable
import android.text.TextWatcher
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
        userVerification()
    }

    private fun userVerification() {
        binding.inputUserName.addTextChangedListener(
            object : TextWatcher {
                override fun beforeTextChanged(
                    s: CharSequence?,
                    start: Int,
                    count: Int,
                    after: Int
                ) {}

                override fun onTextChanged(
                    s: CharSequence?,
                    start: Int,
                    before: Int,
                    count: Int
                ) {
                    if (userUnique(s.toString())) {
                        binding.errorUsername.visibility = View.GONE
                    } else {
                        binding.errorUsername.visibility = View.VISIBLE
                    }
                }
                override fun afterTextChanged(s: Editable?) {}
            }
        )
    }

    private fun userUnique(username: String): Boolean {
        //consulta a backend
        if (username == "c" || username == "cr" || username == "crh") {
            return false
        }
        return true
    }
}