<script lang="ts">
    import { dayOneStr }  from "$lib/input"

    const strArr = dayOneStr.split('\n')

    let count = 0
    const numbers = new Set(['1','2','3','4','5','6','7','8','9','0'])

    // iterate throught the word, if it contains spelled number
    // add the number string to the middle of the string 
    // (case: oneight = 18)
    function expandText (text: string): string {
        let expandedText = text;
        expandedText = expandedText.replaceAll("one", "o1ne")
        .replaceAll("two", "t2wo")
        .replaceAll("three", "t3hree")
        .replaceAll("four", "f4our")
        .replaceAll("five", "f5ive")
        .replaceAll("six", "s6ix")
        .replaceAll("seven", "s7even")
        .replaceAll("eight", "e8ight")
        .replaceAll("nine", "n9ine");

        return expandedText;
    }

    for (let word of strArr) {
        word = expandText(word)
        let numberToAdd: string = ''

        // look from front
        for (let i = 0; i < word.length; i++) {
            // look for number
            if (numbers.has(word[i])) {
                numberToAdd += word[i]
                break
            }
        }

        // look from back
        for (let i = word.length - 1; i >= 0; i--) {
            // look for number
            if (numbers.has(word[i])) {
                numberToAdd += word[i]
                break
            }
        }
        count += Number(numberToAdd)
    }
    console.log(count)
</script>

<div class="w-full h-[90vh] flex items-center justify-center">
    <p class="h3"> Answer: {count}</p>
</div>
