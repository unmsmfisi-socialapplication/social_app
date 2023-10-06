package com.social.utilsTest

import com.social.utils.CodeGenerator
import org.junit.Assert.assertNotEquals
import org.junit.Assert.assertNotNull
import org.junit.Assert.assertTrue
import org.junit.Test

class CodeGeneratorTest {
    @Test
    fun testNoDuplicateCodes() {
        val codes = mutableSetOf<String>()
        for (i in 1..1000) {
            codes.add(CodeGenerator.generateCode())
        }
        assertNotEquals(1000, codes.size)
    }

    @Test
    fun testGeneratedCodesDifferent() {
        val code1 = CodeGenerator.generateCode()
        val code2 = CodeGenerator.generateCode()
        assertNotEquals(code1, code2)
    }

    @Test
    fun testGeneratedCodeFormat() {
        val code = CodeGenerator.generateCode()
        assertTrue(code.matches(Regex("\\d{4}")))
    }

    @Test
    fun testGeneratedCodeNotNull() {
        val code = CodeGenerator.generateCode()
        assertNotNull(code)
        assertTrue(code.isNotEmpty())
    }
}
