
package com.social.presentation.home
import android.os.Bundle
import android.view.LayoutInflater
import android.view.View
import android.view.ViewGroup
import android.widget.ImageView
import android.widget.TextView
import androidx.core.content.ContextCompat
import androidx.fragment.app.Fragment
import com.social.R

class HomeFragment : Fragment() {
    override fun onCreateView(
        inflater: LayoutInflater,
        container: ViewGroup?,
        savedInstanceState: Bundle?,
    ): View? {
        val view = inflater.inflate(R.layout.fragment_home, container, false)

        val iconLike = view.findViewById<ImageView>(R.id.icon_like)
        val txtCountLikes = view.findViewById<TextView>(R.id.txt_countLikes)
        val iconComments = view.findViewById<ImageView>(R.id.icon_comment)
        val txtCountComments = view.findViewById<TextView>(R.id.txt_countComments)
        val iconShare = view.findViewById<ImageView>(R.id.icon_share)
        val txtCountShare = view.findViewById<TextView>(R.id.txt_countShare)
        val iconFavorite = view.findViewById<ImageView>(R.id.icon_favorite)
        val txtCountFavorite = view.findViewById<TextView>(R.id.txt_countFavorite)

        setupLikes(iconLike, txtCountLikes)
        setupShares(iconShare, txtCountShare)
        setupFavorites(iconFavorite, txtCountFavorite)
        setupComments(iconComments, txtCountComments)

        return view
    }

    private fun setupLikes(iconLike: ImageView, txtCountLikes: TextView) {
        var isLiked = false
        var likeCount = 0

        iconLike.setOnClickListener {
            if (!isLiked) {
                likeCount++
                iconLike.setImageResource(R.drawable.likebutton)
                iconLike.setColorFilter(ContextCompat.getColor(requireContext(), R.color.color01))
            } else {
                likeCount--
                iconLike.setImageResource(R.drawable.likebutton)
                iconLike.clearColorFilter()
            }

            txtCountLikes.text = likeCount.toString()

            isLiked = !isLiked
        }
    }

    private fun setupShares(iconShare: ImageView, txtCountShare: TextView) {
        var isShare = false
        var ShareCount = 0

        iconShare.setOnClickListener {
            if (!isShare) {
                ShareCount++
                iconShare.setImageResource(R.drawable.fowardbutton)
                iconShare.setColorFilter(ContextCompat.getColor(requireContext(), R.color.color01))
            } else {
                ShareCount--
                iconShare.setImageResource(R.drawable.fowardbutton)
                iconShare.clearColorFilter()
            }

            txtCountShare.text = ShareCount.toString()

            isShare = !isShare
        }
    }

    private fun setupFavorites(iconFavorite: ImageView, txtCountFavorite: TextView) {
        var isFavorite = false
        var FavoriteCount = 0

        iconFavorite.setOnClickListener {
            if (!isFavorite) {
                FavoriteCount++
                iconFavorite.setImageResource(R.drawable.favoritebutton)
                iconFavorite.setColorFilter(ContextCompat.getColor(requireContext(), R.color.color01))
            } else {
                FavoriteCount--
                iconFavorite.setImageResource(R.drawable.favoritebutton)
                iconFavorite.clearColorFilter()
            }

            txtCountFavorite.text = FavoriteCount.toString()

            isFavorite = !isFavorite
        }
    }

    private fun setupComments(iconComments: ImageView, txtCountComments: TextView) {
        var isComments = false
        var CommentsCount = 0

        iconComments.setOnClickListener {
            if (!isComments) {
                CommentsCount++
                iconComments.setImageResource(R.drawable.commentbutton)
                iconComments.setColorFilter(ContextCompat.getColor(requireContext(), R.color.color01))
            } else {
                CommentsCount--
                iconComments.setImageResource(R.drawable.commentbutton)
                iconComments.clearColorFilter()
            }

            txtCountComments.text = CommentsCount.toString()

            isComments = !isComments
        }
    }
}
