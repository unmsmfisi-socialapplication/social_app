package com.social.presentation.friends

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentAddFriendsBinding

class AddFriendsFragment : Fragment(R.layout.fragment_add_friends) {
    private lateinit var binding: FragmentAddFriendsBinding

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentAddFriendsBinding.bind(view)
    }
}
