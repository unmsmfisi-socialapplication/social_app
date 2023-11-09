package com.social.utilsTest

import android.view.View
import android.view.ViewGroup
import androidx.recyclerview.widget.RecyclerView.ViewHolder
import com.social.utils.BaseAdapter
import org.junit.Assert.assertEquals
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertTrue
import org.junit.Before
import org.junit.Test
import org.mockito.Mockito.mock

class BaseAdapterTest {
    private lateinit var adapter: BaseAdapter<String>

    private class ConcreteViewHolder(view: View) : BaseAdapter.BaseViewHolder<String>(view) {
        override fun bind(entity: String) {
        }
    }

    @Before
    fun setUp() {
        adapter =
            object : BaseAdapter<String>(listOf("Item 1", "Item 2", "Item 3")) {
                override fun getViewHolder(parent: ViewGroup): BaseViewHolder<String> {
                    return ConcreteViewHolder(mock(View::class.java))
                }
            }
    }

    @Test
    fun getItemCountTest() {
        val count = adapter.itemCount
        assertEquals(3, count)
    }

    @Test
    fun updateListTest() {
        val newData = listOf("New Item 1", "New Item 2")
        adapter.updateList(newData)
        assertEquals(newData, adapter.data)
    }

    @Test
    fun onCreateViewHolderTest() {
        val viewHolder: ViewHolder =
            adapter.onCreateViewHolder(mock(ViewGroup::class.java), 0)
        assertNotNull(viewHolder)
        assertTrue(viewHolder is ConcreteViewHolder)
    }

    @Test
    fun itemViewTypeTest() {
        val viewType = adapter.getItemViewType(0)
        assertEquals(0, viewType)
    }
}
