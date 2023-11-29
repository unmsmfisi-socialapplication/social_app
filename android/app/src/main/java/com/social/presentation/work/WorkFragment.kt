package com.social.presentation.work

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.social.R
import com.social.databinding.FragmentWorkBinding

class WorkFragment : Fragment(R.layout.fragment_work) {
    private lateinit var binding: FragmentWorkBinding

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentWorkBinding.bind(view)
    }
}
