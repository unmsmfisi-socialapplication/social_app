package com.social.presentation.home

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentNewPostBinding

class NewPostFragment : Fragment(R.layout.fragment_new_post) {
    private lateinit var binding: FragmentNewPostBinding

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentNewPostBinding.bind(view)

        action()
    }

    private fun action() {
        binding.iconLocation.setOnClickListener {
        }
    }
}
