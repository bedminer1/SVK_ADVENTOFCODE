<script lang="ts">
    import { dayTwoStr } from "$lib/input"

    const strArr = dayTwoStr.split('\n')
    let count = 0
    
    for (let index in strArr) {
        const game = strArr[index]
        const setArr = game.split(':')[1].split(';')
        const numbers = new Set(['1','2','3','4','5','6','7','8','9','0'])

        let greenMax = 0
        let blueMax = 0
        let redMax = 0

        // check for over
        for (let set of setArr) {
            let cubeArr = set.split(',')

            for (let cube of cubeArr) {
                // extract out numbers, might be more than one digit
                let cubeNumber = ''
                for (let c of cube) {
                    if (numbers.has(c)) 
                        cubeNumber += c
                }

                if (cube.includes('blue')) {
                    blueMax = Math.max(Number(cubeNumber), blueMax)
                }

                if (cube.includes('green')) {
                    greenMax = Math.max(Number(cubeNumber), greenMax)
                }

                if (cube.includes('red')) {
                    redMax = Math.max(Number(cubeNumber), redMax)
                }

            }
        }

        count += redMax * greenMax * blueMax
    }
</script>

<div class="w-full h-[90vh] flex items-center justify-center">
    <p class="h3"> Answer: {count}</p>
</div>