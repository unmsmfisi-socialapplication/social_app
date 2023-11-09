package com.social.presentation.interaction.chats

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import com.social.R
import com.social.databinding.FragmentChatScreenBinding
import com.social.databinding.FragmentListChatBinding
import com.social.utils.Toast

class ListChatsFragment : Fragment(R.layout.fragment_list_chat) {
    private lateinit var binding: FragmentListChatBinding

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentListChatBinding.bind(view)
        action()
    }

    private fun action() {
        binding.btnBar.setOnClickListener {
            BarButton()
        }

        binding.btnListChats.setOnClickListener{
            ListChatsButton()
        }

        binding.btnMessages.setOnClickListener{
            MessageButton()
        }

        binding.btnMessagesConfig.setOnClickListener {
            ListChatsConfigButton()
        }
    }

    private fun BarButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun ListChatsButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun MessageButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun ListChatsConfigButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    // backend
}
