package com.social.presentation.interaction.chats

import android.graphics.Bitmap
import android.view.LayoutInflater
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView
import com.social.databinding.ContainerReceivedMessageBinding
import com.social.databinding.ContainerSendMessageBinding
import com.social.domain.model.ChatMessageUser

class ChatAdapter(
    val chatMessage: List<ChatMessageUser>,
    private val receiverProfileImage: Bitmap,
    private val sendProfileImage: Bitmap,
    private val senderId: String,
) : RecyclerView.Adapter<RecyclerView.ViewHolder>() {
    companion object {
        const val VIEW_TYPE_SENT = 1
        const val VIEW_TYPE_RECEIVER = 2
    }

    override fun onCreateViewHolder(
        parent: ViewGroup,
        viewType: Int,
    ): RecyclerView.ViewHolder {
        val inflater = LayoutInflater.from(parent.context)
        return when (viewType) {
            VIEW_TYPE_SENT -> {
                val binding = ContainerSendMessageBinding.inflate(inflater, parent, false)
                SentMessageViewHolder(binding)
            }
            VIEW_TYPE_RECEIVER -> {
                val binding = ContainerReceivedMessageBinding.inflate(inflater, parent, false)
                ReceivedMessageViewHolder(binding)
            }
            else -> throw IllegalArgumentException("Invalid view type")
        }
    }

    override fun onBindViewHolder(
        holder: RecyclerView.ViewHolder,
        position: Int,
    ) {
        when (holder) {
            is SentMessageViewHolder -> holder.setData(chatMessage[position], sendProfileImage)
            is ReceivedMessageViewHolder -> holder.setData(chatMessage[position], receiverProfileImage)
        }
    }

    override fun getItemCount(): Int {
        return chatMessage.size
    }

    override fun getItemViewType(position: Int): Int {
        return if (chatMessage[position].senderId == senderId) {
            VIEW_TYPE_SENT
        } else {
            VIEW_TYPE_RECEIVER
        }
    }

    inner class SentMessageViewHolder(private val binding: ContainerSendMessageBinding) :
        RecyclerView.ViewHolder(binding.root) {
        fun setData(
            chatMessageUser: ChatMessageUser,
            sendProfileImage: Bitmap,
        ) {
            binding.sendMessage.text = chatMessageUser.message
            binding.sendHourMessage.text = chatMessageUser.dateTime
            binding.messageUserProfile.setImageBitmap(sendProfileImage)
        }
    }

    inner class ReceivedMessageViewHolder(private val binding: ContainerReceivedMessageBinding) :
        RecyclerView.ViewHolder(binding.root) {
        fun setData(
            chatMessageUser: ChatMessageUser,
            receiverProfileImage: Bitmap,
        ) {
            binding.receivedMessage.text = chatMessageUser.message
            binding.receivedHourMessage.text = chatMessageUser.dateTime
            binding.messageContactProfile.setImageBitmap(receiverProfileImage)
        }
    }
}
