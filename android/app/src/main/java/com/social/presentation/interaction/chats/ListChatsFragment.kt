package com.social.presentation.interaction.chats

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import com.social.R
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
            barButton()
        }

        binding.btnListChats.setOnClickListener {
            listChatsButton()
        }

        binding.btnMessages.setOnClickListener {
            messageButton()
        }

        binding.btnMessagesConfig.setOnClickListener {
            listChatsConfigButton()
        }
    }

    private fun barButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun listChatsButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun messageButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun listChatsConfigButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    // backend
}
