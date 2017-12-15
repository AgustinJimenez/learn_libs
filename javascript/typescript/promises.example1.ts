function longProcess(){
	return new Promise((resolve, reject) => {
		setTimeout(() => {
				let result = 0;
				for( let i = 0; i <= 2000000000; i++ )
					result = result + 1;
				resolve( result );
			}, 0);
		} 
	)
}

async function main() {
	let result = await longProcess()
						.then((res) => console.log(res))
						.catch((error) => console.log(error));
	console.log("I am done");
}

main();

for(let i = 0; i <= 10; i++ ){
	console.log("i is ${i}");
}














