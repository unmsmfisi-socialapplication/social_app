package com.social.presentation.friends

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentFriendsBinding

class FriendsFragment : Fragment(R.layout.fragment_friends) {
    private lateinit var binding: FragmentFriendsBinding

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentFriendsBinding.bind(view)
    }
}
