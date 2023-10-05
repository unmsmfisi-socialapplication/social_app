package com.social.presentation.authentication

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import com.social.R
import com.social.databinding.FragmentRegisterBinding

class RegisterFragment : Fragment(R.layout.fragment_register) {
   private lateinit var binding: FragmentRegisterBinding
   private lateinit var globalView: View

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentRegisterBinding.bind(view)
        globalView = view

        action()
    }

    private fun action(){
        binding.textLogIn.setOnClickListener {
            findNavController().navigate(R.id.action_registerFragment_to_loginFragment)
        }
    }
}