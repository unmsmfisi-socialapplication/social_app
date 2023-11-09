import { validatePassword } from '../utilities/Functions'
//const { distanceLevenshtein, mostSimilarPhrase } = require('./index');
import { distanceLevenshtein } from '../utilities/Functions'
import { mostSimilarPhrase } from '../utilities/Functions'
import { findMatchingWords } from '../utilities/Functions'

//Test: Validate passwords
describe('Validate passwords', () => {
    it('It should return true if the passwords match', () => {
        const password1 = 'password123'
        const password2 = 'password123'
        const result = validatePassword(password1, password2)
        expect(result).toBe(true)
    })

    it('It should return an error message if the passwords do not match', () => {
        const password1 = 'password123'
        const password2 = 'anotherpassword'
        const result = validatePassword(password1, password2)
        expect(result).toBe(false)
    })
})

//Test: Text Similarity Functions
describe('Text Similarity Functions', () => {
    describe('distanceLevenshtein', () => {
        it('Should return the correct Levenshtein distance for similar strings', () => {
            const distance = distanceLevenshtein('kitten', 'sitting')
            expect(distance).toBe(3)
        })

        it('Should return 0 for identical strings', () => {
            const distance = distanceLevenshtein('apple', 'apple')
            expect(distance).toBe(0)
        })

        it('Should return the correct Levenshtein distance for different strings', () => {
            const distance = distanceLevenshtein('hello', 'world')
            expect(distance).toBe(4)
        })
    })

    describe('mostSimilarPhrase', () => {
        const phrases = ['apple', 'banana', 'cherry', 'date', 'fig']

        it('Should return the most similar phrase when a similar one exists', () => {
            const similarPhrase = mostSimilarPhrase('apples', phrases)
            expect(similarPhrase).toBe('apple')
        })

        it('Should return the exact phrase when its in the list', () => {
            const exactPhrase = mostSimilarPhrase('bananas', phrases)
            expect(exactPhrase).toBe('banana')
        })

        it('Should return the most similar phrase when no exact match is found', () => {
            const similarPhrase = mostSimilarPhrase('dates', phrases)
            expect(similarPhrase).toBe('date')
        })
    })
})

//Test: Function findMatchingWords
describe('findMatchingWords', () => {
    it('Should return an empty array if the input text is empty', () => {
        const text = ''
        const words = ['apple', 'banana', 'cherry']
        const result = findMatchingWords(text, words)
        expect(result).toEqual(['The entrance is empty'])
    })

    it('It should return a list of words that match the input text', () => {
        const text = 'app'
        const words = ['apple', 'banana', 'cherry']
        const result = findMatchingWords(text, words)
        expect(result).toEqual(['apple'])
    })

    it('Should return "No matching words found" if there are no matches', () => {
        const text = 'grape'
        const words = ['apple', 'banana', 'cherry']
        const result = findMatchingWords(text, words)
        expect(result).toEqual(['No matching words found'])
    })

    it('Should handle case-sensitive cases correctly', () => {
        const text = 'BaNaNa'
        const words = ['apple', 'banana', 'cherry']
        const result = findMatchingWords(text, words)
        expect(result).toEqual(['banana'])
    })
})
