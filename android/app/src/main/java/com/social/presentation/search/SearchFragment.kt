package com.social.presentation.search

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentSearchBinding

class SearchFragment : Fragment(R.layout.fragment_search) {
    private lateinit var binding: FragmentSearchBinding

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentSearchBinding.bind(view)
    }
}
