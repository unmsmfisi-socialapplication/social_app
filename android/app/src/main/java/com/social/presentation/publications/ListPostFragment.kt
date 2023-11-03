package com.social.presentation.publications

import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import androidx.fragment.app.Fragment
import androidx.lifecycle.ViewModelProvider
import com.social.R
import com.social.databinding.FragmentListPostBinding
import com.social.databinding.ItemPostBinding
import com.social.domain.model.Post
import com.social.utils.BaseAdapter

class ListPostFragment : Fragment(R.layout.fragment_list_post) {
    private lateinit var binding: FragmentListPostBinding
    private lateinit var globalView: View

    private val adapter: BaseAdapter<Post> = object : BaseAdapter<Post>(emptyList()){
        override fun getViewHolder(parent: ViewGroup) : BaseViewHolder<Post>{
            val view = LayoutInflater.from(parent.context)
                .inflate(R.layout.item_post, parent, false)
            return object : BaseViewHolder<Post>(view){
                private val binding: ItemPostBinding = ItemPostBinding.bind(view)
                override fun bind(entity: Post) = with(binding){
                    textNames.text = entity.names
                    textHour.text = entity.hour
                    textContentPost.text = entity.content
                }
            }
        }
    }

    private val viewModel: ListPostViewModel by lazy {
        ViewModelProvider(this)[ListPostViewModel::class.java]
    }

    override fun onViewCreated(view: View, savedInstanceState: Bundle?) {
        super.onViewCreated(view, savedInstanceState)
        binding = FragmentListPostBinding.bind(view)
        globalView = view

        setupAdapter()
        observeViewModel()
    }

    private fun setupAdapter(){
        binding.recyclerPost.adapter = adapter
    }

    private fun observeViewModel() {
        viewModel.data.observe(viewLifecycleOwner) { posts ->
            adapter.updateList(posts)
        }
        viewModel.obtainData()
    }
}