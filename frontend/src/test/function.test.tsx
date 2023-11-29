//const { distanceLevenshtein, mostSimilarPhrase } = require('./index');
import {
    validatePassword,
    distanceLevenshtein,
    mostSimilarPhrase,
    findMatchingWords,
    generateUniqueUsernames,
    filterContentByTag,
    countCharacters,
} from '../utilities/Functions'

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

//Test: generateUniqueUsernames
describe('generateUniqueUsernames', () => {
    it('Should generate three unique usernames based on the base username and existing usernames', () => {
        const existingUsernames = ['user1', 'user2', 'user3']
        const baseUsername = 'newUser'
        const numAlternatives = 3

        const uniqueUsernames = generateUniqueUsernames(existingUsernames, baseUsername, numAlternatives)

        // Check if the generated usernames are unique and not in the existing usernames
        expect(uniqueUsernames.length).toEqual(numAlternatives)
        uniqueUsernames.forEach((username) => {
            expect(existingUsernames.includes(username)).toBeFalsy()
        })
    })

    it('Should handle the case where numAlternatives is zero', () => {
        const existingUsernames = ['user1', 'user2', 'user3']
        const baseUsername = 'user'
        const numAlternatives = 0

        const uniqueUsernames = generateUniqueUsernames(existingUsernames, baseUsername, numAlternatives)

        // Check if the result is an empty array
        expect(uniqueUsernames).toEqual([])
    })
})
//Test: Function filterContentByTag
describe('Filter content by tag', () => {
    const contentArray = [
        { tags: ['#javascript', '#tutorial'], content: 'Contenido del #tutorial de #JavaScript' },
        { tags: ['#typescript', '#tutorial'], content: 'Contenido del #tutorial de #TypeScript' },
        { tags: ['#javascript', '#guide'], content: '#Guide of #JavaScript' },
        { tags: ['#typescript', '#guide'], content: '#Guide of #TypeScript' },
        { tags: ['#python', '#tutorial'], content: 'Contenido del #tutorial de #Python' },
    ]

    it('should filter content by tag and return an array of matching content', () => {
        const tagToFilter = '#javascript'
        const filteredContent = filterContentByTag(tagToFilter, contentArray)
        const expectedContent = ['Contenido del #tutorial de #JavaScript', '#Guide of #JavaScript']
        expect(filteredContent).toEqual(expectedContent)
    })

    it('should return an empty array if no content matches the tag', () => {
        const tagToFilter = '#ruby'
        const filteredContent = filterContentByTag(tagToFilter, contentArray)
        expect(filteredContent).toEqual([])
    })
})
describe('countCharacters', () => {
    it('should return the correct character count for valid input', () => {
        expect(countCharacters('Hello', 10)).toBe(5)
        expect(countCharacters('This is a test', 15)).toBe(14)
    })

    it('should return the maxLength if input exceeds maxLength', () => {
        expect(countCharacters('TooLongInput', 5)).toBe(5)
        expect(countCharacters('123456789', 5)).toBe(5)
    })

    it('should handle empty input correctly', () => {
        expect(countCharacters('', 10)).toBe(0)
    })
})
