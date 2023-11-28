//TODO: Create function reutilization file for all the constant values
export function validatePassword(password1: string, password2: string): boolean {
    return password1 === password2
}

export function distanceLevenshtein(str1: string, str2: string): number {
    if (str1.length < str2.length) {
        return distanceLevenshtein(str2, str1)
    }

    if (str2.length === 0) {
        return str1.length
    }

    let previousRow: number[] = []
    for (let i = 0; i <= str2.length; i++) {
        previousRow.push(i)
    }

    for (let i = 0; i < str1.length; i++) {
        let currentRow: number[] = [i + 1]

        for (let j = 0; j < str2.length; j++) {
            const insertions = previousRow[j + 1] + 1
            const deletions = currentRow[j] + 1
            const substitutions = previousRow[j] + (str1[i] !== str2[j] ? 1 : 0)

            currentRow.push(Math.min(insertions, deletions, substitutions))
        }

        previousRow = currentRow
    }

    return previousRow[str2.length]
}

export function mostSimilarPhrase(phrase: string, arrayPhrase: string[]): string | null {
    let mostSimilarPhrase: string | null = null
    let minimalDistance: number = Infinity

    for (const f of arrayPhrase) {
        const distance = distanceLevenshtein(phrase, f)

        if (distance < minimalDistance) {
            minimalDistance = distance
            mostSimilarPhrase = f
        }
    }

    return mostSimilarPhrase
}

export function findMatchingWords(text: string, words: string[]): string[] {
    if (text === '') {
        return ['The entrance is empty']
    }

    const searchText = text.toLowerCase()

    const wordsMatching = words.filter((word) => word.toLowerCase().indexOf(searchText) === 0)

    if (wordsMatching.length > 0) {
        return wordsMatching
    } else {
        return ['No matching words found']
    }
}

export function generateUniqueUsernames(
    existingUsernames: string[],
    baseUsername: string,
    numAlternatives: number,
): string[] {
    const uniqueUsernames: string[] = []
    let suffix = 1

    while (uniqueUsernames.length < numAlternatives) {
        const newUsername = `${baseUsername}${suffix}`

        if (!existingUsernames.includes(newUsername) && !uniqueUsernames.includes(newUsername)) {
            uniqueUsernames.push(newUsername)
        }

        suffix++
    }

    return uniqueUsernames
}
export function filterContentByTag(tag: string, contentArray: { tags: string[]; content: string }[]): string[] {
    return contentArray.filter((item) => item.tags.some((t) => t === tag)).map((item) => item.content)
}
//This function must be called on the text input event (every time the user enters a character in a post)
export function countCharacters(inputText: string, maxLength: number): number {
    if (inputText.length <= maxLength) {
        return inputText.length
    } else {
        return maxLength
    }
}

export const useHistory = (url: string) => {
    window.history.pushState({}, '', url)
}
