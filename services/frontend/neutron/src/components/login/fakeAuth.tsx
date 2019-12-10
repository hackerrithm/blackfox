export async function login({ username, password }: any) {
	return new Promise((resolve, reject) => {
		setTimeout(() => {
			if (username === "kem" && password === "123") {
				resolve();
			} else {
				reject();
			}
		}, 1000);
	});
}
