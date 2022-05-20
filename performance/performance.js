const start = new Date().getTime()

for (let i = 0; i < 10000000000; i++) {
	if (i % 1000000000 === 0) {
		console.log(i)
	}
}

const end = new Date().getTime()

const seconds = (end - start) / 1000
console.log(`${seconds}s`)
