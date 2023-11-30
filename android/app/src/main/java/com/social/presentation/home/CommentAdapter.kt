package com.social.presentation.home

import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView
import com.social.databinding.ItemCommentBinding
import com.social.domain.model.CommentData

class CommentAdapter(private val comments: List<CommentData>) :
    RecyclerView.Adapter<CommentAdapter.CommentViewHolder>() {

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): CommentViewHolder {
        val inflater = LayoutInflater.from(parent.context)
        val binding = ItemCommentBinding.inflate(inflater, parent, false)
        return CommentViewHolder(binding.root)
    }

    override fun onBindViewHolder(holder: CommentViewHolder, position: Int) {
        val comment = comments[position]
        holder.bind(comment)
    }

    override fun getItemCount(): Int = comments.size

    inner class CommentViewHolder(itemView: View) : RecyclerView.ViewHolder(itemView) {
        private val binding = ItemCommentBinding.bind(itemView)

        fun bind(comment: CommentData) {
            binding.txtUsernamePost.text = comment.username
            binding.txtTimeComment.text = comment.time
            binding.txtReactionComment.text = comment.reactionText
            binding.txtReplyComment.text = comment.replyText
            binding.txtCountReactionComment.text = comment.reactionCount.toString()
        }
    }
}
