package com.social.presentation.interaction.chats

import android.os.Bundle
import android.view.View
import androidx.fragment.app.Fragment
import androidx.navigation.fragment.findNavController
import com.social.R
import com.social.databinding.FragmentChatScreenBinding
import com.social.domain.model.ChatUserData
import com.social.utils.Toast
import com.vanniktech.emoji.EmojiManager
import com.vanniktech.emoji.EmojiPopup
import com.vanniktech.emoji.google.GoogleEmojiProvider

class ChatScreenFragment : Fragment(R.layout.fragment_chat_screen) {
    private lateinit var binding: FragmentChatScreenBinding
    private lateinit var receiverUser: ChatUserData
    private lateinit var emojiPopup: EmojiPopup

    override fun onViewCreated(
        view: View,
        savedInstanceState: Bundle?,
    ) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentChatScreenBinding.bind(view)
        setupViews()
    }

    private fun setupViews() {
        EmojiManager.install(GoogleEmojiProvider())

        emojiPopup = EmojiPopup.Builder.fromRootView(binding.rootView).build(binding.inputSendMessage)

        binding.btnFace.setOnClickListener {
            emojiPopup.toggle()
        }

        binding.btnSendMessage.setOnClickListener {
            val message = binding.inputSendMessage.text.toString()
            sendMessage(message)
        }

        binding.btnCall.setOnClickListener {
            onCallUserButton()
        }

        binding.btnVideocam.setOnClickListener {
            onVideoCamUserButton()
        }

        binding.btnCamera.setOnClickListener {
            onCameraButton()
        }

        binding.btnMicrophone.setOnClickListener {
            onMicrophoneButton()
        }

        binding.btnBack.setOnClickListener {
            findNavController().navigate(R.id.action_chatScreenFragment_to_listChatsFragment)
        }
    }

    private fun onCallUserButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun onVideoCamUserButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun onCameraButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun onMicrophoneButton() {
        Toast.showMessage(requireContext(), "En desarrollo")
    }

    private fun sendMessage(message: String) {
        Toast.showMessage(requireContext(), message)
    }
}
